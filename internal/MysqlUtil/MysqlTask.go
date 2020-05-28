package MysqlUtil

import (
	"errors"
	"gindemo/internal/Model/MysqlModel"
	"gindemo/internal/Model/RedisModel"
	"gindemo/internal/Model/ServiceModel"
	"gindemo/internal/RedisUtil"
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/mysql"
	"strings"
	"time"
)

func Submit(body *RedisModel.HistoryInfo) error {
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
		updateMap := map[string]interface{}{
			"useTime":        int64(body.UseTime),
			"title":          strings.Join(*body.Title, ","),
			"coverUrl":       body.CoverUrl,
			"updateCount":    user.VideoHistories.UpdateCount + 1,
			"isDel":          false,
			"submitDateTime": body.SubmitDate,
			"updated_at":     time.Now(),
		}
		result = db.Model(&user.VideoHistories).Updates(updateMap)
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

func Clear(body *ServiceModel.ClearHistoryParameter) error {
	if body == nil {
		return errors.New("body is nil")
	}

	sql := "update videoHistories set isDel=1,updated_at=now() where userId=(select id from historyInfo where openId=?);"
	result := db.Exec(sql, body.OpenId)
	if result.Error != nil {
		return result.Error
	}

	if err := RedisUtil.Clear(body.OpenId); err != nil {
		return err
	}

	return nil
}
