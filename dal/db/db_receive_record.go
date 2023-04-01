package db

import (
	"errors"
	"ginDemo/model"
	"ginDemo/utils"
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

var tableNameReceive = "rp_receive_record"

func QueryByUserId(userId string) (*model.RpReceiveRecord, error) {
	var record model.RpReceiveRecord
	err := Rdb.Table(tableNameReceive).Where("user_id = ?", userId).First(&record).Error
	if err != nil {
		utils.Errorf("can not find userId amount, userId: %v, err: %v", userId, err)
		return nil, err
	}
	return &record, nil
}

func InsertRecord(record *model.RpReceiveRecord) (int64, error) {
	err := Rdb.Table(tableNameReceive).Create(&record).Error
	if err != nil {
		utils.Errorf("insert data err: %v\n", err)
		return 0, err
	}
	return record.Id, nil

}

func QueryReceiveRecordByBizOutNoAndUserId(c *gin.Context, bizOutNo, userId string) (*model.RpReceiveRecord, error) {
	var record model.RpReceiveRecord
	// find 和first的区别：record not find报错--first；find不报错
	err := Rdb.Table(tableNameReceive).WithContext(c).Where("user_id = ?", userId).Where("biz_out_no = ?", bizOutNo).First(&record).Error
	if err != nil {
		if errors.As(err, &gorm.ErrRecordNotFound) {
			return nil, nil
		}
		utils.Errorf("dal.QuerySendRecordByBizOutNoAndUserId query error %v", err)
		return nil, err
	}
	return &record, nil
}
