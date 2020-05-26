package MysqlModel

import (
	"time"
)

type VideoHistories struct {
	Id             int64      `gorm:"column:id"`
	UserId         int64      `gorm:"column:userId"`
	VideoId        string     `gorm:"column:videoId"`
	UseTime        int64      `gorm:"column:useTime"`
	Title          string     `gorm:"column:title"`
	CoverUrl       string     `gorm:"column:coverUrl"`
	UpdateCount    int        `gorm:"column:updateCount"`
	IsDel          bool       `gorm:"column:isDel"`
	SubmitDateTime time.Time  `gorm:"column:submitDateTime"`
	CreatedAt      time.Time  `gorm:"column:created_at"`
	UpdatedAt      time.Time  `gorm:"column:updated_at"`
	DeletedAt      *time.Time `gorm:"column:deleted_at"`
}

func (VideoHistories) TableName() string {
	return "videoHistories"
}

func BuildVideoHistories(userId int64, videoId string, useTime int, title string, coverUrl string) VideoHistories {
	return VideoHistories{
		UserId:         userId,
		VideoId:        videoId,
		UseTime:        int64(useTime),
		CoverUrl:       coverUrl,
		Title:          title,
		SubmitDateTime: time.Now(),
	}
}
