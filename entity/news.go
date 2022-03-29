package entity

import "time"

type News struct {
	ID         uint
	Title      string    `json:"title" gorm:"column:title"`
	Body       string    `json:"body" gorm:"column:body"`
	CreateTime time.Time `json:"create_time" gorm:"column:create_time"`
	UpdateTime time.Time `json:"update_time" gorm:"column:update_time"`
}
