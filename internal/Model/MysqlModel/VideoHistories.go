package MysqlModel

import (
	"time"
)

type VideoHistories struct {
	Id          int64     `gorm:"column:id"`
	UserId      int64     `gorm:"column:userId"`
	VideoId     string    `gorm:"column:videoId"`
	UseTIme     int64     `gorm:"column:useTime"`
	Title       string    `gorm:"column:title"`
	CoverUrl    string    `gorm:"column:coverUrl"`
	UpdateCount int       `gorm:"column:updateCount"`
	IsDel       bool      `gorm:"column:isDel"`
	SubmitDate  time.Time `gorm:"column:submitDate"`
	CreatedAt   time.Time `gorm:"column:createdAt"`
	UpdatedAt   time.Time `gorm:"column:updatedAt"`
}

func (VideoHistories) TableName() string {
	return "videoHistories"
}
