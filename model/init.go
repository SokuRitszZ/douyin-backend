package model

import (
	"douyin-backend/config"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

var Db *gorm.DB

func Init() (err error) {
	Db, err = gorm.Open(mysql.Open(config.DSN), &gorm.Config{})
	return err
}

func Migrate(model interface{}) error {
	err := Db.AutoMigrate(&model)
	return err
}
