package handler

import (
	"errors"
	"ginDemo/consts"
	"ginDemo/dal/db"
	"ginDemo/dal/kv"
	"ginDemo/model"
	"ginDemo/service/strategy"
	"ginDemo/utils"
	"github.com/gin-gonic/gin"
	"github.com/go-sql-driver/mysql"
	"github.com/google/uuid"
	"github.com/sirupsen/logrus"
	"strings"
	"time"
)

// traceID

func SendRedPacket(c *gin.Context) {
	// 1. 参数绑定
	var sReq model.SendRpReq
	err := c.BindJSON(&sReq)
	if err != nil {
		logrus.Error("[SendRedPacket] bind req json error")
		utils.RetErrJson(c, consts.BindError)
	}
	// 2. 参数判断
	ok := checkParams(sReq)
	if !ok {
		logrus.Errorf("[SendRedPacket] check params error, sReq: %v", utils.Json2String(sReq))
		utils.RetErrJson(c, consts.ParamError)
	}
	// 3. 账户、风控校验，略
	// http请求 (ctx-->logid/traceId)
	if sReq.UserId == sReq.GroupId {
		utils.RetErrJson(c, consts.ParamError)
	}

	// 4. 幂等校验
	record, rErr := db.QueryRecordByBizOutNoAndUserId(c, sReq.BizOutNo, sReq.UserId)
	if rErr != nil {
		logrus.Errorf("[SendRedPacket] query db error %v", rErr)
		utils.RetErrJson(c, consts.ServiceBusy)
	}
	if record != nil {
		logrus.Infof("[SendRedPacket] bizOutNo has one record already")
		utils.RetJsonWithData(c, utils.Json2String(record))
	}

	// 初始化一个新的发放记录
	var newRecord model.RpSendRecord
	newRecord.RpId = strings.ReplaceAll(uuid.New().String(), "-", "")
	// 5. 读取过期设置，先设置常量
	newRecord.SendTime = time.Now()
	newRecord.ExpireTime = time.Now().Add(consts.ExpireTime24)
	// 6. 红包预拆包，将结果写入到map中
	var receiveAmountList []int64
	remain := sReq.Amount
	sum := int64(0)
	for i := int64(0); i < sReq.Number; i++ {
		x := strategy.DoubleAverage(sReq.Number-i, remain)
		receiveAmountList = append(receiveAmountList, x)
		remain -= x
		sum += x
	}
	kErr := kv.LPushRp(c, newRecord.RpId, receiveAmountList)
	if kErr != nil {
		logrus.Errorf("[SendRedPacket] insert receive amount into redis error %v", kErr)
		utils.RetErrJson(c, consts.ServiceBusy)
	}
	// 7. 写入发放记录,可以判断一下重复error
	buildSendRecord(newRecord, sReq)
	id, dErr := db.InsertSendRecord(c, &newRecord)
	// err有两种情况 1. 数据插入重复   2. 数据库有问题
	if dErr != nil {
		var mysqlErr *mysql.MySQLError
		if errors.As(dErr, &mysqlErr) && mysqlErr.Number == 1062 {
			//  幂等返回
			oldRecord, oErr := db.QueryRecordByBizOutNoAndUserId(c, sReq.BizOutNo, sReq.UserId)
			if oErr != nil {
				logrus.Errorf("[SendRedPacket] old record query db error %v", oErr)
				utils.RetErrJson(c, consts.ServiceBusy)
			}
			if oldRecord != nil {
				logrus.Infof("[SendRedPacket] bizOutNo has one record already")
				utils.RetJsonWithData(c, utils.Json2String(record))
			}
		} else {
			logrus.Warnf("[SendRedPacket] bizOutNo has one record already")
			utils.RetErrJson(c, consts.InsertError)
		}
	}
	logrus.Infof("[SendRedPacket]: insert rp record success, auto increase id is : %v", id)
	// 8. 发送延迟消息，期间进行一次对账
	// 发一个消息告诉某人，这个红包在xx时刻会过期，如果过期了，请你帮我把红包设置成过期状态，如果这个时候红包没有领完，请你把剩下的钱转给发红包的用户。 todo
	// 简单对账
	// 1. 初始化一个list
	// 2. lpop->list
	// 3. 对账成功， list-> redis
	rList, rlErr := kv.LLenRp(c, newRecord.RpId)
	if rlErr != nil {
		logrus.Warnf("[SendRedPacket] bizOutNo has one record already")
	} else {
		if len(*rList) != len(receiveAmountList) {
			// 1. 回滚数据库、删除发放记录,作业 todo

			// 2. 报错
			utils.RetErrJson(c, consts.ServiceBusy)
		}
	}
	//if amountListInMap, okk := sMap[newRecord.RpId]; okk {
	//	var total int64
	//	for _, val := range amountListInMap {
	//		total += val
	//	}
	//	if total == sReq.Amount {
	//		logrus.Infof("[SendRedPacket] amountListInMap equals user amount")
	//	} else {
	//		// 1. 回滚数据库、删除发放记录,作业 todo
	//
	//		// 2. 报错
	//		utils.RetErrJson(c, consts.ServiceBusy)
	//	}
	//}
	// 9 扣款,请求资金服务

}

func buildSendRecord(record model.RpSendRecord, req model.SendRpReq) {
	record.UserId = req.UserId
	record.GroupChatId = req.GroupId
	record.BizOutNo = req.BizOutNo
	record.Amount = req.Amount
	record.ReceiveAmount = 0
	record.Number = req.Number
	record.Status = consts.RpStatusSend
	record.SendTime = time.Now()
	record.CreateTime = time.Now()
	record.ModifyTime = time.Now()
}

func checkParams(seq model.SendRpReq) bool {
	return !(seq.UserId == "" || seq.GroupId == "" || seq.Amount <= 0 || (seq.Number*seq.Amount) <= 1 || seq.BizOutNo == "")
}

func QuerySendRecords(c *gin.Context) {

}

func ReceiveRedPacket(c *gin.Context) {
	// 1. 参数绑定

	// 2. 参数检查

	// 3. 幂等检查

	// 4. 读取红包记录

	// 5.读取红包个数

	// 6. 生成领取信息

	// 7.

}

func QueryReceiveRecords(c *gin.Context) {

}
