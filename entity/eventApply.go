package entity

import (
	"time"
	"vol-backend/global"
)

type EventApply struct {
	ID          uint
	EventID     int       `json:"event_id" gorm:"column:event_id"`
	UserID      int       `json:"user_id" gorm:"column:user_id"`
	ApplyStatus string    `json:"apply_status" gorm:"column:"` //报名中/报名通过/报名失败
	CreateTime  time.Time `json:"create_time" gorm:"column:create_time"`
	UpdateTime  time.Time `json:"update_time" gorm:"column:update_time"`
}

func CreateEventApply(eventApply EventApply) error {
	err := global.VB_DB.
		Table(TABLE_NAME_EVENT_APPLY).
		Omit("create_time", "update_time").
		Create(&eventApply).Error
	if err != nil {
		return err
	}
	return nil
}

func UpdateEventApplyStatus(status string, id uint) error {
	err := global.VB_DB.
		Table(TABLE_NAME_EVENT_APPLY).
		Where("id = ?", id).
		Update("apply_status", status).Error
	if err != nil {
		return err
	}
	return nil
}
