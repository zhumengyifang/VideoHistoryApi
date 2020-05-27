package MysqlUtil

import (
	"errors"
	"fmt"
	"gindemo/internal/Config"
	"gindemo/internal/Model/MysqlModel"
	"gindemo/internal/Model/ServiceModel"
	"gindemo/internal/RedisUtil"
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/mysql"
	"strings"
)

var db *gorm.DB

func init() {
	db = CreateConn()
}

func CreateConn() *gorm.DB {
	connArgs := fmt.Sprintf("%s:%s@(%s:%d)/%s?charset=utf8&parseTime=true&loc=Local", Config.GetMysql().User, Config.GetMysql().Password, Config.GetMysql().Host, Config.GetMysql().Port, "History")
	db, err := gorm.Open("mysql", connArgs)
	if err != nil {
		panic(err.Error())
	}

	db.DB().SetMaxIdleConns(Config.GetMysql().MaxIdleConns)
	db.DB().SetMaxOpenConns(Config.GetMysql().MaxOpenConns)
	db.LogMode(Config.GetMysql().SqlLog)

	return db
}

func Info(body *ServiceModel.InfoHistoryParameter) (*MysqlModel.HistoryInfo, error) {
	if body == nil {
		return nil, errors.New("body is nil")
	}

	historyInfo := MysqlModel.HistoryInfo{}
	result := db.First(&historyInfo)
	if result.Error != nil {
		return nil, result.Error
	}

	result = db.Find(&historyInfo.VideoHistories, "userId=? and videoId=? and isDel=0", historyInfo.Id, body.VideoId)
	if result.Error != nil {
		return nil, result.Error
	}

	return &historyInfo, nil
}

func List(body *ServiceModel.ListHistoryParameter) (*MysqlModel.HistoryInfo, error) {
	if body == nil {
		return nil, errors.New("body is nil")
	}

	historyInfo := MysqlModel.HistoryInfo{}
	result := db.Find(&historyInfo, "openId=?", body.OpenId)
	if result.Error != nil {
		return nil, result.Error
	}

	sql := "select videoId, useTime, submitDateTime, isDel, title, coverUrl, submitDateTime from videoHistories  where userId = ? and isDel=0 order by submitDateTime desc limit ?,?;"
	result = db.Raw(sql, historyInfo.Id, (body.PageCount-1)*body.PageSize, body.PageSize).Scan(&historyInfo.VideosHistories)
	if result.Error != nil {
		return nil, result.Error
	}

	return &historyInfo, nil
}

func Submit(body *ServiceModel.SubmitHistoryParameter) error {
	if body == nil {
		return errors.New("body is nil")
	}

	user := MysqlModel.BuildUsers(body.OpenId, body.AuthorName)
	result := db.FirstOrCreate(&user, "openId=?", user.OpenId)
	if result.Error != nil {
		return result.Error
	}

	user.VideoHistories = MysqlModel.BuildVideoHistories(user.Id, body.VideoId, body.UseTime, strings.Join(*body.Title, ","), body.CoverUrl)
	result = db.Find(&user.VideoHistories, "userId=? and videoId=?", user.Id, body.VideoId)
	if result.Error != nil && result.Error != gorm.ErrRecordNotFound {
		return result.Error
	}

	if result.Error == gorm.ErrRecordNotFound {
		result = db.Create(&user.VideoHistories)
		if result.Error != nil {
			return result.Error
		}
	} else {
		user.VideoHistories.UseTime = int64(body.UseTime)
		user.VideoHistories.Title = strings.Join(*body.Title, ",")
		user.VideoHistories.CoverUrl = body.CoverUrl
		user.VideoHistories.IsDel = false

		result = db.Model(&user.VideoHistories).Update(user.VideoHistories)
		if result.Error != nil {
			return result.Error
		}
	}

	err := RedisUtil.DelCommand(body.OpenId, []string{body.VideoId})
	if err != nil {
		return err
	}

	return nil
}

func Del(body *ServiceModel.DelHistoryParameter) error {
	if body == nil {
		return errors.New("body is nil")
	}

	sql := "update videoHistories set isDel=1,updated_at=now() where userId=(select id from historyInfo where openId=?) and videoId in(?);"
	result := db.Exec(sql, body.OpenId, body.VideoIds)
	if result.Error != nil {
		return result.Error
	}

	err := RedisUtil.DelCommand(body.OpenId, body.VideoIds)
	if err != nil {
		return err
	}

	return nil
}

func Clear(openId string) error {
	if openId == "" {
		return errors.New("body is nil")
	}

	sql := "update videoHistories set isDel=1,updated_at=now() where userId=(select id from historyInfo where openId=?);"
	result := db.Exec(sql, openId)
	if result.Error != nil {
		return result.Error
	}

	if err := RedisUtil.Clear(openId); err != nil {
		return err
	}

	return nil
}
