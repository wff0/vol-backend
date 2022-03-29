package entity

import "time"

type UserInfo struct {
	ID         uint
	UserName   string    `json:"user_name" gorm:"column:user_name"`
	Password   string    `json:"password" gorm:"column:password"`
	School     string    `json:"school" gorm:"column:school"`
	Class      string    `json:"class" gorm:"column:class"`
	Role       int       `json:"role" gorm:"column:role"` //0 普通用户 1 管理员
	CreateTime time.Time `json:"create_time" gorm:"column:create_time"`
	UpdateTime time.Time `json:"update_time" gorm:"column:update_time"`
}
