package config

import (
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

func NewMysql() *gorm.DB {
	db, err := gorm.Open(mysql.Open(Conf.DB.Mysql), &gorm.Config{})
	if err != nil {
		panic(err)
	}
	return db
}
