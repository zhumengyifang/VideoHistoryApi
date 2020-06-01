package MysqlUtil

import (
	"errors"
	"fmt"
	"gindemo/internal/Config"
	"gindemo/internal/Model/MysqlModel"
	"gindemo/internal/Model/ServiceModel"
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/mysql"
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
	result := db.Select("id,openId,authorName").Find(&historyInfo, "openId=?", body.OpenId)
	if result.Error != nil {
		return nil, result.Error
	}

	result = db.Select("videoId, useTime, submitDateTime, isDel, title, coverUrl, submitDateTime").Find(&historyInfo.VideoHistories, "userId=? and videoId=? and isDel=0", historyInfo.Id, body.VideoId)
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
	result := db.Select("id,openId,authorName").Find(&historyInfo, "openId=?", body.OpenId)
	if result.Error != nil {
		return nil, result.Error
	}

	sql := "select videoId, useTime, submitDateTime, isDel, title, coverUrl, submitDateTime from videoHistories  where userId = ? and isDel=0 order by userId,isDel,submitDateTime desc limit ?,?;"
	result = db.Raw(sql, historyInfo.Id, (body.PageCount-1)*body.PageSize, body.PageSize).Scan(&historyInfo.VideosHistories)
	if result.Error != nil {
		return nil, result.Error
	}

	return &historyInfo, nil
}
