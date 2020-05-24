package MysqlUtil

import (
	"fmt"
	"gindemo/internal/Config"
	"gindemo/internal/Model/MysqlModel"
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/mysql"
	"time"
)

func init() {
	test()
}

func test() {
	connArgs := fmt.Sprintf("%s:%s@(%s:%d)/%s?charset=utf8&parseTime=true&loc=Local", Config.GetMysql().User, Config.GetMysql().Password, Config.GetMysql().Host, Config.GetMysql().Port, "History")
	db, err := gorm.Open("mysql", connArgs)
	if err != nil {
		panic("连接数据库失败")
	}

	user := MysqlModel.Users{OpenId: "2a0adb40c5624410baeed12426ced271", AuthorName: "jarvis"}
	p := db.Create(&user)
	videoHistories := MysqlModel.VideoHistories{UserId: user.Id, VideoId: "2a0adb40c5624410baeed12426ced272", UseTIme: 100, Title: "class", CoverUrl: "www.baidu.com", SubmitDate: time.Now()}
	p = db.Create(&videoHistories)
	fmt.Println(p.Error)

	defer db.Close()
}
