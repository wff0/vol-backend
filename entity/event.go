package entity

import "time"

type Event struct {
	ID          uint
	Title       string    `json:"title" gorm:"column:title"`
	Location    string    `json:"location" gorm:"column:location"`
	UserID      int       `json:"user_id" gorm:"column:user_id"`
	Description string    `json:"description" gorm:"column:description"`
	Status      string    `json:"status" gorm:"column:status"` //报名中/进行中/已结束
	MaxNum      int       `json:"max_num" gorm:"column:max_num"`
	StartTime   time.Time `json:"start_time" gorm:"column:start_time"`
	EndTime     time.Time `json:"end_time" gorm:"column:end_time"`
	CreateTime  time.Time `json:"create_time" gorm:"column:create_time"`
	UpdateTime  time.Time `json:"update_time" gorm:"column:update_time"`
}
