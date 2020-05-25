package MysqlModel

import (
	"time"
)

type Users struct {
	Id             int64          `gorm:"column:id"`
	OpenId         string         `gorm:"column:openId"`
	AuthorName     string         `gorm:"column:authorName"`
	VideoHistories VideoHistories `gorm:"ForeignKey:UserId;AssociationForeignKey:Id"`
	CreatedAt      time.Time      `gorm:"column:created_at"`
	UpdatedAt      time.Time      `gorm:"column:updated_at"`
	DeletedAt      *time.Time     `gorm:"column:deleted_at"`
}

func (Users) TableName() string {
	return "users"
}

func BuildUsers(openId string, AuthorName string) *Users {
	return &Users{
		OpenId:     openId,
		AuthorName: AuthorName,
	}
}
