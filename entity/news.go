package entity

import (
	"time"
	"vol-backend/global"
)

type News struct {
	ID         uint
	Title      string    `json:"title" gorm:"column:title"`
	Body       string    `json:"body" gorm:"column:body"`
	CreateTime time.Time `json:"create_time" gorm:"column:create_time"`
	UpdateTime time.Time `json:"update_time" gorm:"column:update_time"`
}

func CreateNews(news News) error {
	err := global.VB_DB.
		Table(TABLE_NAME_NEWS).
		Omit("create_time", "update_time").
		Create(&news).Error
	if err != nil {
		return err
	}
	return nil
}
