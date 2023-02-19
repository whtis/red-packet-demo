package db

import (
	"errors"
	"ginDemo/model"
	"github.com/gin-gonic/gin"
	"github.com/sirupsen/logrus"
)

var tableNameSend = "rp_send_record"

func QueryRecordByBizOutNoAndUserId(c *gin.Context, bizOutNo, userId string) (*model.RpSendRecord, error) {
	var record model.RpSendRecord
	err := rdb.Table(tableNameSend).WithContext(c).Where("user_id = ?", userId).Where("biz_out_no = ?", bizOutNo).Find(&record).Error
	if err != nil {
		logrus.Errorf("dal.QueryRecordByBizOutNoAndUserId query error %v", err)
		return nil, err
	}
	if record.Id == 0 {
		return nil, nil
	} else {
		return &record, errors.New("record is existed")
	}
}

func InsertSendRecord(c *gin.Context, record *model.RpSendRecord) (int64, error) {
	err := rdb.Table(tableNameSend).WithContext(c).Create(record).Error
	// err有两种情况 1. 数据库有问题   2. 数据插入重复
	if err != nil {
		logrus.Errorf("dal.InsertSendRecord error %v", err)
		return 0, err
	}
	return record.Id, err
}
