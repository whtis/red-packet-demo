package dal

import (
	"ginDemo/model"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"log"
	"time"
)

var rdb *gorm.DB

var tableName = "rp_receive_record"

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

func QueryByUserId(userId string) (*model.RpReceiveRecord, error) {
	var record model.RpReceiveRecord
	err := rdb.Table("rp_receive_record").Where("user_id = ?", userId).First(&record).Error
	if err != nil {
		log.Printf("can not find userId amount, userId: %v, err: %v", userId, err)
		return nil, err
	}
	return &record, nil
}

func InsertRecord(record *model.RpReceiveRecord) (int64, error) {
	err := rdb.Table("rp_receive_record").Create(&record).Error
	if err != nil {
		log.Printf("insert data err: %v\n", err)
		return 0, err
	}
	return record.Id, nil

}
