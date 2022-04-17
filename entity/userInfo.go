package entity

import (
	"time"
	"vol-backend/global"
)

type UserInfo struct {
	ID         uint
	Username   string    `json:"username" gorm:"column:username"`
	Password   string    `json:"password" gorm:"column:password"`
	Gender     string    `json:"gender" gorm:"column:gender"`
	School     string    `json:"school" gorm:"column:school"`
	Class      string    `json:"class" gorm:"column:class"`
	CreateTime time.Time `json:"create_time" gorm:"column:create_time"`
	UpdateTime time.Time `json:"update_time" gorm:"column:update_time"`
}

func CreateUserInfo(userInfo UserInfo) error {
	err := global.VB_DB.
		Table(TABLE_NAME_USER_INFO).
		Omit("create_time", "update_time").
		Create(&userInfo).Error
	if err != nil {
		return err
	}
	return nil
}
