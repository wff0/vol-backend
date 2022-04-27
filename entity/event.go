package entity

import (
	"time"
	"vol-backend/global"
)

type Event struct {
	ID          uint
	Title       string    `json:"title" gorm:"column:title"`
	Location    string    `json:"location" gorm:"column:location"`
	UserID      int       `json:"user_id" gorm:"column:user_id"`
	Description string    `json:"description" gorm:"column:description"`
	Status      string    `json:"status" gorm:"column:status"` //报名中/进行中/已结束
	MaxNum      int       `json:"max_num" gorm:"column:max_num"`
	StartTime   string    `json:"start_time" gorm:"column:start_time"`
	EndTime     string    `json:"end_time" gorm:"column:end_time"`
	CreateTime  time.Time `json:"create_time" gorm:"column:create_time"`
	UpdateTime  time.Time `json:"update_time" gorm:"column:update_time"`
}

func CreateEvent(event Event) error {
	event.Status = "报名中"
	err := global.VB_DB.
		Table(TABLE_NAME_EVENT).
		Omit("create_time", "update_time").
		Create(&event).Error
	if err != nil {
		return err
	}
	return nil
}

func UpdateEventAll(event Event, id uint) error {
	err := global.VB_DB.
		Table(TABLE_NAME_EVENT).
		Where("id = ?", id).
		Updates(map[string]interface{}{"title": event.Title,
			"location": event.Location,
			//"user_id":     event.UserID,
			"description": event.Description,
			"status":      event.Status,
			"max_num":     event.MaxNum,
			"start_time":  event.StartTime,
			"end_time":    event.EndTime}).Error
	if err != nil {
		return err
	}
	return nil
}

func UpdateEventStatus(id uint, status string) error {
	err := global.VB_DB.
		Table(TABLE_NAME_EVENT).
		Where("id = ?", id).
		Update("status", status).Error
	if err != nil {
		return err
	}
	return nil
}

func SelectEventById(id uint) (*Event, error) {
	var event *Event
	err := global.VB_DB.Table(TABLE_NAME_EVENT).Where("id = ?", id).Find(&event).Error
	if err != nil {
		return nil, err
	}
	return event, nil
}

func DeleteEventByID(id uint) error {
	var event *Event
	err := global.VB_DB.Table(TABLE_NAME_EVENT).Where("id = ?", id).Delete(&event).Error
	if err != nil {
		return err
	}
	return nil
}

type EventPackage struct {
	ID          uint   `json:"id"`
	Title       string `json:"title"`
	Location    string `json:"location"`
	Username    string `json:"username"`
	Description string `json:"description"`
	Status      string `json:"status"` //报名中/进行中/已结束
	MaxNum      int    `json:"max_num"`
	StartTime   string `json:"start_time"`
	EndTime     string `json:"end_time"`
}

func GetEventList(page int, pageSize int) ([]EventPackage, error) {
	var list []EventPackage
	err := global.VB_DB.
		Table(TABLE_NAME_EVENT).
		Select("event.id", "title", "location",
			"username", "description", "status", "max_num", "start_time", "end_time").
		Joins("left join user_info on event.user_id = user_info.id").
		Offset((page - 1) * pageSize).
		Limit(pageSize).
		Find(&list).
		Error
	if err != nil {
		return nil, err
	}
	return list, nil
}

func GetEventCount() (int64, error) {
	var count int64
	err := global.VB_DB.
		Table(TABLE_NAME_EVENT).
		Count(&count).
		Error
	if err != nil {
		return 0, nil
	}
	return count, nil
}
