package initialize

import (
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"vol-backend/config"
)

func GormMysql() (*gorm.DB, error) {
	m := config.Mysql{
		Path:     "",
		Port:     "3306",
		Dbname:   "db",
		Username: "wff",
		Password: "WFFxhDX520",
		Config:   "charset=utf8mb4&parseTime=True&loc=Local",
	}
	db, err := gorm.Open(mysql.Open(m.Dsn()), &gorm.Config{})
	if err != nil {
		return nil, err
	}
	return db, nil
}
