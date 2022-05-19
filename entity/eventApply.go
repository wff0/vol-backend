package entity

import (
	"time"
	"vol-backend/global"
)

type EventApply struct {
	ID          uint
	EventID     int       `json:"event_id" gorm:"column:event_id"`
	UserID      int       `json:"user_id" gorm:"column:user_id"`
	ApplyStatus int       `json:"apply_status" gorm:"column:apply_status"` //0 报名中/1 报名通过/2 报名失败
	Remark      string    `json:"remark" gorm:"column:remark"`
	CreateTime  time.Time `json:"create_time" gorm:"column:create_time"`
	UpdateTime  time.Time `json:"update_time" gorm:"column:update_time"`
}

func SelectEventApplyRecordByID(id uint) (*EventApply, error) {
	var tmp *EventApply
	err := global.VB_DB.
		Table(TABLE_NAME_EVENT_APPLY).
		Where("id = ?", id).
		First(&tmp).Error
	if err != nil {
		return nil, err
	}
	return tmp, nil
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

func SearchUserIfApplyEvent(eventID uint, userID uint) (*EventApply, error) {
	record := &EventApply{}
	err := global.VB_DB.
		Table(TABLE_NAME_EVENT_APPLY).
		Where("event_id = ? and user_id = ?", eventID, userID).
		First(record).Error
	if err != nil {
		return nil, err
	}
	return record, nil
}

type EventApplyListPackage struct {
	ID          uint   `json:"id"`
	Gender      string `json:"gender"`
	School      string `json:"school"`
	Classroom   string `json:"classroom"`
	Username    string `json:"username"`
	ApplyStatus int    `json:"apply_status"`
	Remark      string `json:"remark"`
}

func GetAllEventApplyListByEventID(id uint) ([]EventApplyListPackage, error) {
	var list []EventApplyListPackage
	err := global.VB_DB.
		Table(TABLE_NAME_EVENT_APPLY).
		Select("event_apply.id", "username",
			"apply_status", "remark", "gender", "school", "classroom").
		Joins("left join user_info on event_apply.user_id = user_info.id").
		Where("event_id = ?", id).
		Order("event_apply.create_time DESC").
		Find(&list).
		Error
	if err != nil {
		return nil, err
	}
	return list, nil
}

type EventApplyListPackageToUser struct {
	ID          uint   `json:"key"`
	Title       string `json:"title"`
	ApplyStatus int    `json:"apply_status"`
	Remark      string `json:"remark"`
}

func GetAllEventApplyListByUserID(id uint) ([]EventApplyListPackageToUser, error) {
	var list []EventApplyListPackageToUser
	err := global.VB_DB.
		Table(TABLE_NAME_EVENT_APPLY).
		Select("event_apply.id", "title",
			"apply_status", "remark").
		Joins("left join event on event_apply.event_id = event.id").
		Where("event_apply.user_id = ?", id).
		Order("event_apply.create_time DESC").
		Find(&list).
		Error
	if err != nil {
		return nil, err
	}
	return list, nil
}

func UpdateEventApplyRemark(remark string, id uint) error {
	err := global.VB_DB.
		Table(TABLE_NAME_EVENT_APPLY).
		Where("id = ?", id).
		Update("remark", remark).Error
	if err != nil {
		return err
	}
	return nil
}

func DeleteEventApplyByID(id uint) error {
	var eventApply EventApply
	err := global.VB_DB.
		Table(TABLE_NAME_EVENT_APPLY).
		Where("id = ?", id).
		Delete(&eventApply).Error
	if err != nil {
		return err
	}
	return nil
}
