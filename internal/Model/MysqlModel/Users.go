package MysqlModel

import (
	"time"
)

type Users struct {
	Id         int64     `gorm:"column:id"`
	OpenId     string    `gorm:"column:openId"`
	AuthorName string    `gorm:"column:authorName"`
	CreatedAt  time.Time `gorm:"column:createdAt"`
	UpdatedAt  time.Time `gorm:"column:updatedAt"`
}

func (Users) TableName()string{
	return "users"
}