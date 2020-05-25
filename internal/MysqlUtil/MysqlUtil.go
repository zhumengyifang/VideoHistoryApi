package MysqlUtil

import (
	"errors"
	"fmt"
	"gindemo/internal/Config"
	"gindemo/internal/Model/MysqlModel"
	"gindemo/internal/Model/ServiceModel"
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

func Info(body *ServiceModel.InfoHistoryParameter) error {
	if body == nil {
		return errors.New("body is nil")
	}

	var user MysqlModel.Users
	db.First(&user)
	err := db.Model(&user).Association("VideoHistories").Find(&user.VideoHistories)
	if err.Error != nil {
		return err.Error
	}

	//var user1 MysqlModel.Users
	//db.First(&user1)
	//err = db.Model(&user1).Related(&user1.VideoHistories).Find(&user1.VideoHistories)
	//if err.Error != nil {
	//	return err.Error
	//}

	return nil
}

func Submit(body *ServiceModel.SubmitHistoryParameter) error {
	if body == nil {
		return errors.New("body is nil")
	}

	user := MysqlModel.BuildUsers(body.OpenId, body.AuthorName)
	err := db.FirstOrCreate(user, "openId=?", user.OpenId)
	if err.Error != nil {
		return err.Error
	}

	videoHistory := MysqlModel.BuildVideoHistories(user.Id, body.VideoId, body.UseTime, strings.Join(*body.Title, ","), body.CoverUrl)
	err = db.Create(&videoHistory)
	if err.Error != nil {
		return err.Error
	}

	return nil
}
