package dal

import (
	"errors"
	"ginDemo/model"
	"github.com/sirupsen/logrus"
)

var tableNameSend = "rp_send_record"

func QueryRecordByBizOutNoAndUserId(bizOutNo, userId string) (*model.RpSendRecord, error) {
	var record model.RpSendRecord
	err := rdb.Table(tableNameSend).Where("user_id = ?", userId).Where("biz_out_no = ?", bizOutNo).Find(&record).Error
	if err != nil {
		logrus.Error("dal.QueryRecordByBizOutNoAndUserId query error %v", err)
		return nil, err
	}
	if record.Id == 0 {
		return nil, nil
	} else {
		return &record, errors.New("record is existed")
	}
}
