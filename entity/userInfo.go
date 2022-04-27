package entity

import (
	"time"
	"vol-backend/global"
)

type UserInfo struct {
	ID         uint      `json:"id" gorm:"column:id"`
	Username   string    `json:"username" gorm:"column:username"`
	Password   string    `json:"password" gorm:"column:password"`
	Gender     string    `json:"gender" gorm:"column:gender"`
	School     string    `json:"school" gorm:"column:school"`
	Classroom  string    `json:"classroom" gorm:"column:classroom"`
	Role       int       `json:"role" gorm:"column:role"` // 1 管理员 2 志愿者
	CreateTime time.Time `json:"create_time" gorm:"column:create_time"`
	UpdateTime time.Time `json:"update_time" gorm:"column:update_time"`
}

func CreateUserInfo(userInfo UserInfo) error {
	userInfo.Role = 2
	err := global.VB_DB.
		Table(TABLE_NAME_USER_INFO).
		Omit("create_time", "update_time").
		Create(&userInfo).Error
	if err != nil {
		return err
	}
	return nil
}

func UpdateUserInfoByID(info UserInfo, id uint) error {
	err := global.VB_DB.Table(TABLE_NAME_USER_INFO).
		Where("id = ?", id).
		Updates(map[string]interface{}{
			"username":  info.Username,
			"password":  info.Password,
			"school":    info.School,
			"gender":    info.Gender,
			"classroom": info.Classroom,
			"role":      info.Role}).Error
	if err != nil {
		return err
	}
	return nil
}

func DeleteUserByID(id uint) error {
	var user *UserInfo
	err := global.VB_DB.
		Table(TABLE_NAME_USER_INFO).
		Where("id = ?", id).
		Delete(&user).Error
	if err != nil {
		return err
	}
	return nil
}

func FindUserByUsernameAndPassword(userInfo UserInfo) (*UserInfo, error) {
	res := &UserInfo{}
	err := global.VB_DB.
		Table(TABLE_NAME_USER_INFO).
		Where("username = ? and password = ?", userInfo.Username, userInfo.Password).
		First(&res).
		Error
	if err != nil {
		return nil, err
	}
	return res, nil
}

func GetUserInfoList(page int, pageSize int) ([]UserInfo, error) {
	var list []UserInfo
	err := global.VB_DB.
		Table(TABLE_NAME_USER_INFO).
		Select("id", "username", "gender", "school", "classroom", "role").
		Offset((page - 1) * pageSize).
		Limit(pageSize).
		Find(&list).
		Error
	if err != nil {
		return nil, err
	}
	return list, nil
}

func GetUserCount() (int64, error) {
	var count int64
	err := global.VB_DB.
		Table(TABLE_NAME_USER_INFO).
		Count(&count).
		Error
	if err != nil {
		return 0, nil
	}
	return count, nil
}
