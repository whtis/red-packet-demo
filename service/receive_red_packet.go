package service

import (
	"ginDemo/consts"
	"ginDemo/dal/db"
	"ginDemo/dal/kv"
	"ginDemo/model"
	"ginDemo/utils"
	"github.com/gin-gonic/gin"
	"github.com/redis/go-redis/v9"
	"github.com/sirupsen/logrus"
	"time"
)

func ReceiveRedPacket(c *gin.Context) {
	// 1. 参数绑定
	var rReq model.ReceiveRpReq
	err := c.BindJSON(&rReq)
	if err != nil {
		logrus.Error("[ReceiveRedPacket] bind req json error")
		utils.RetErrJson(c, consts.BindError)
		return
	}
	// 2. 参数检查
	ok := checkReceiveParams(rReq)
	if !ok {
		logrus.Errorf("[ReceiveRedPacket] check params error, rReq: %v", utils.Json2String(rReq))
		utils.RetErrJson(c, consts.ParamError)
		return
	}
	// 3. 幂等检查
	receiveRecord, rErr := db.QueryReceiveRecordByBizOutNoAndUserId(c, rReq.BizOutNo, rReq.UserId)
	if rErr != nil {
		logrus.Errorf("[ReceiveRedPacket] query db error %v", err)
		utils.RetErrJson(c, consts.ServiceBusy)
		return
	}
	if receiveRecord != nil {
		// 请求重入，返回上一次领取的红包记录
		utils.RetJsonWithData(c, utils.Json2String(receiveRecord))
		return
	}
	// 4. 查询发放记录，判断是否可以发放
	sendRecord, sErr := db.QuerySendRecordByRpId(c, rReq.RpId)
	if sErr != nil {
		utils.RetErrJson(c, consts.ServiceBusy)
		return
	}
	// 4. 读取红包记录
	// 校验发送红包记录的状态，只有状态为已发送时才能继续领取红包
	if sendRecord.Status != consts.RpStatusSend {
		logrus.Infof("[ReceiveRedPacket] rp record received or expired, status:%d", sendRecord.Status)
		utils.RetErrJson(c, consts.ServiceBusy)
		return
	}
	// 校验红包是否过期
	if time.Now().After(sendRecord.ExpireTime) {
		logrus.Errorf("[ReceiveRedPacket] rp record has been expired")
		utils.RetErrJson(c, consts.RpExpiredError)
		return
	}
	// 5.领红包
	receiveAmount, aErr := kv.LPop(c, sendRecord.RpId)
	if aErr != nil {
		if aErr == redis.Nil {
			// 红包领完了
			utils.RetErrJson(c, consts.RpReceivedError)
			return
		}
		utils.RetErrJson(c, consts.ServiceBusy)
		return
	}
	// 6. 生成领取信息&更新发放记录数据
	dbReceiveRecord := buildReceiveRecord(rReq, receiveAmount)
	sendRecord.ReceiveAmount = sendRecord.ReceiveAmount + receiveAmount
	// 7. gorm事务
	uErr := db.UpdateSendAndCreateReceiveRecordTx(c, sendRecord, dbReceiveRecord)
	if uErr != nil {
		// todo 把这个单个红包给放回redis队尾
		utils.RetErrJson(c, consts.ServiceBusy)
		return
	}
	utils.RetJsonWithData(c, utils.Json2String(receiveRecord))
	return
}

func buildReceiveRecord(req model.ReceiveRpReq, amount int64) *model.RpReceiveRecord {
	return &model.RpReceiveRecord{
		UserId:      req.UserId,
		GroupChatId: req.GroupId,
		RpId:        req.RpId,
		Amount:      amount,
		BizOutNo:    req.BizOutNo,
		ReceiveTime: time.Now(),
		CreateTime:  time.Now(),
		ModifyTime:  time.Now(),
	}
}

func checkReceiveParams(req model.ReceiveRpReq) bool {
	return !(req.UserId == "" || req.GroupId == "" || req.BizOutNo == "")
}
