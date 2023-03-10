package db

import (
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"time"
)

var rdb *gorm.DB

func InitDB() {
	dsn := "root:root2023@tcp(127.0.0.1:3306)/tech?charset=utf8mb4&parseTime=True&loc=Local"
	var err error
	rdb, err = gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		panic(err)
	}

	sqlDB, _ := rdb.DB()

	// SetMaxIdleConns 设置空闲连接池中连接的最大数量
	sqlDB.SetMaxIdleConns(10)

	// SetMaxOpenConns 设置打开数据库连接的最大数量。
	sqlDB.SetMaxOpenConns(100)

	// SetConnMaxLifetime 设置了连接可复用的最大时间。
	sqlDB.SetConnMaxLifetime(time.Hour)
}
