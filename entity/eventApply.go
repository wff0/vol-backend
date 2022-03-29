package entity

import "time"

type EventApply struct {
	ID          uint
	EventID     int       `json:"event_id" gorm:"column:event_id"`
	UserID      int       `json:"user_id" gorm:"column:user_id"`
	ApplyStatus string    `json:"apply_status" gorm:"column:"apply_status` //报名中/报名通过/报名失败
	CreateTime  time.Time `json:"create_time" gorm:"column:create_time"`
}
