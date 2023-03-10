package db

import (
	"errors"
	"ginDemo/model"
	"github.com/gin-gonic/gin"
	"github.com/sirupsen/logrus"
	"gorm.io/gorm"
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

func QueryReceiveRecordByBizOutNoAndUserId(c *gin.Context, bizOutNo, userId string) (*model.RpReceiveRecord, error) {
	var record model.RpReceiveRecord
	// find 和first的区别：record not find报错--first；find不报错
	err := rdb.Table(tableNameReceive).WithContext(c).Where("user_id = ?", userId).Where("biz_out_no = ?", bizOutNo).First(&record).Error
	if err != nil {
		if errors.As(err, &gorm.ErrRecordNotFound) {
			return nil, nil
		}
		logrus.Errorf("dal.QuerySendRecordByBizOutNoAndUserId query error %v", err)
		return nil, err
	}
	return &record, nil
}
