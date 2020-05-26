package MysqlModel

import (
	"time"
)

type HistoryInfo struct {
	Id              int64          `gorm:"column:id"`
	OpenId          string         `gorm:"column:openId"`
	AuthorName      string         `gorm:"column:authorName"`
	VideoHistories  VideoHistories `gorm:"ForeignKey:UserId"`
	VideosHistories []VideoHistories
	CreatedAt       time.Time  `gorm:"column:created_at"`
	UpdatedAt       time.Time  `gorm:"column:updated_at"`
	DeletedAt       *time.Time `gorm:"column:deleted_at"`
}

func (HistoryInfo) TableName() string {
	return "historyInfo"
}

func BuildUsers(openId string, AuthorName string) HistoryInfo {
	return HistoryInfo{
		OpenId:     openId,
		AuthorName: AuthorName,
	}
}
