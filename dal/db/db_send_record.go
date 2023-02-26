package db

import (
	"errors"
	"ginDemo/model"
	"github.com/gin-gonic/gin"
	"github.com/sirupsen/logrus"
)

var tableNameSend = "rp_send_record"

func QuerySendRecordByBizOutNoAndUserId(c *gin.Context, bizOutNo, userId string) (*model.RpSendRecord, error) {
	// todo 这里着重要讲下find和first的区别
	var record model.RpSendRecord
	err := rdb.Table(tableNameSend).WithContext(c).Where("user_id = ?", userId).Where("biz_out_no = ?", bizOutNo).Find(&record).Error
	if err != nil {
		logrus.WithContext(c).Errorf("dal.QuerySendRecordByBizOutNoAndUserId query error %v", err)
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

func QuerySendRecordByRpId(c *gin.Context, rpId string) (*model.RpSendRecord, error) {
	var record model.RpSendRecord
	err := rdb.Table(tableNameSend).WithContext(c).Where("rp_id = ?", rpId).First(&record).Error
	if err != nil {
		logrus.Errorf("dal.QuerySendRecordByBizOutNoAndUserId query error %v", err)
		return nil, err
	}
	return &record, nil
}

func QuerySendRecordByCond(c *gin.Context, req model.QuerySendRecordReq) ([]*model.RpSendRecord, error) {
	records := make([]*model.RpSendRecord, 0)
	// select * from send_record where user_id = xxx and (group_chat_id = ?) order by create_time desc limit size
	tx := rdb.Table(tableNameSend).WithContext(c).Where("user_id = ?", req.UserId)
	if req.GroupId != "" {
		tx.Where("group_chat_id = ?", req.GroupId)
	}
	tx.Where("create_time < ?", req.Cursor) // todo 类型转换 作业

	err := tx.Order("create_time desc").Limit(int(req.Size)).Find(&records).Error
	if err != nil {
		logrus.Errorf("dal.QuerySendRecordByCond query error %v", err)
		return nil, err
	}
	return records, nil
}

func QuerySendRecordByCondPage(c *gin.Context, req model.QuerySendRecordReqByPage) ([]*model.RpSendRecord, error) {
	records := make([]*model.RpSendRecord, 0)
	// select * from send_record where user_id = xxx and (group_chat_id = ?) order by create_time desc limit size
	tx := rdb.Table(tableNameSend).WithContext(c).Where("user_id = ?", req.UserId)
	if req.GroupId != "" {
		tx.Where("group_chat_id = ?", req.GroupId)
	}

	err := tx.Order("create_time desc").Offset(int(req.Page)).Limit(int(req.Size)).Find(&records).Error
	if err != nil {
		logrus.Errorf("dal.QuerySendRecordByCondPage query error %v", err)
		return nil, err
	}
	return records, nil
}

func CountSendRecordByCondPage(c *gin.Context, req model.QuerySendRecordReqByPage) (int64, error) {
	count := int64(0)
	// select * from send_record where user_id = xxx and (group_chat_id = ?) order by create_time desc limit size
	tx := rdb.Table(tableNameSend).WithContext(c).Where("user_id = ?", req.UserId)
	if req.GroupId != "" {
		tx.Where("group_chat_id = ?", req.GroupId)
	}

	err := tx.Order("create_time desc").Offset(int(req.Page)).Limit(int(req.Size)).Count(&count).Error
	if err != nil {
		logrus.Errorf("dal.QuerySendRecordByCondPage query error %v", err)
		return 0, err
	}
	return count, nil
}
