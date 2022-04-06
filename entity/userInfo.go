package entity

import (
	"time"
	"vol-backend/global"
)

type UserInfo struct {
	ID         uint
	Username   string    `json:"username" gorm:"column:username"`
	Password   string    `json:"password" gorm:"column:password"`
	School     string    `json:"school" gorm:"column:school"`
	Class      string    `json:"class" gorm:"column:class"`
	Role       int       `json:"role" gorm:"column:role"` //0 普通用户 1 管理员
	CreateTime time.Time `json:"create_time" gorm:"column:create_time"`
	UpdateTime time.Time `json:"update_time" gorm:"column:update_time"`
}

func CreateUserInfo(userInfo UserInfo) error {
	userInfo.Role = 0
	err := global.VB_DB.
		Table(TABLE_NAME_USER_INFO).
		Omit("create_time", "update_time").
		Create(&userInfo).Error
	if err != nil {
		return err
	}
	return nil
}
