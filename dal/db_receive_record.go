package dal

import (
	"ginDemo/model"
	"log"
)

var tableNameReceive = "rp_receive_record"

func QueryByUserId(userId string) (*model.RpReceiveRecord, error) {
	var record model.RpReceiveRecord
	err := rdb.Table(tableNameReceive).Where("user_id = ?", userId).First(&record).Error
	if err != nil {
		log.Printf("can not find userId amount, userId: %v, err: %v", userId, err)
		return nil, err
	}
	return &record, nil
}

func InsertRecord(record *model.RpReceiveRecord) (int64, error) {
	err := rdb.Table(tableNameReceive).Create(&record).Error
	if err != nil {
		log.Printf("insert data err: %v\n", err)
		return 0, err
	}
	return record.Id, nil

}
