package db

import (
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
)

var Rdb *gorm.DB

func InitDB() {
	dsn := "root:jrttroot@tcp(127.0.0.1:3306)/tech?charset=utf8mb4&parseTime=True&loc=Local"
	var err error
	Rdb, err = gorm.Open(mysql.Open(dsn), &gorm.Config{
		Logger: logger.Default.LogMode(logger.Info),
	})
	if err != nil {
		panic(err)
	}
}
