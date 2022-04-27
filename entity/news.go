package entity

import (
	"time"
	"vol-backend/global"
)

type News struct {
	ID         uint
	Title      string    `json:"title" gorm:"column:title"`
	Body       string    `json:"body" gorm:"column:body"`
	UserID     int       `json:"user_id" gorm:"column:user_id"`
	CreateTime time.Time `json:"create_time" gorm:"column:create_time"`
	UpdateTime time.Time `json:"update_time" gorm:"column:update_time"`
}

func CreateNews(news News) error {
	err := global.VB_DB.
		Table(TABLE_NAME_NEWS).
		Omit("create_time", "update_time").
		Create(&news).Error
	if err != nil {
		return err
	}
	return nil
}

func UpdateNewsByID(news News, id uint) error {
	err := global.VB_DB.Table(TABLE_NAME_NEWS).
		Where("id = ?", id).
		Updates(map[string]interface{}{
			"title": news.Title,
			"body":  news.Body}).Error
	if err != nil {
		return err
	}
	return nil
}

func DeleteNewsByID(id uint) error {
	var news *News
	err := global.VB_DB.
		Table(TABLE_NAME_NEWS).
		Where("id = ?", id).
		Delete(&news).Error
	if err != nil {
		return err
	}
	return nil
}

type NewsPackage struct {
	ID       uint   `json:"id"`
	Title    string `json:"title"`
	Body     string `json:"body"`
	Username string `json:"username"`
}

func GetNewsList(page int, pageSize int) ([]NewsPackage, error) {
	var list []NewsPackage
	err := global.VB_DB.
		Table(TABLE_NAME_NEWS).
		Select("news.id", "title", "body", "username").
		Joins("left join user_info on news.user_id = user_info.id").
		Offset((page - 1) * pageSize).
		Limit(pageSize).
		Find(&list).
		Error
	if err != nil {
		return nil, err
	}
	return list, nil
}

func GetNewsCount() (int64, error) {
	var count int64
	err := global.VB_DB.
		Table(TABLE_NAME_NEWS).
		Count(&count).
		Error
	if err != nil {
		return 0, nil
	}
	return count, nil
}
