package handler

import (
	"fmt"
	"ginDemo/consts"
	"ginDemo/dal"
	"ginDemo/model"
	"ginDemo/service/strategy"
	"ginDemo/utils"
	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
	"github.com/sirupsen/logrus"
	"strings"
	"time"
)

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
		logrus.Warnf("[SendRedPacket] check params error, sReq: %v", utils.Json2String(sReq))
		utils.RetErrJson(c, consts.ParamError)
	}
	// 3. 账户、风控校验，略
	if sReq.UserId == sReq.GroupId {
		utils.RetErrJson(c, consts.ParamError)
	}

	// 4. 幂等校验
	record, rErr := dal.QueryRecordByBizOutNoAndUserId(sReq.BizOutNo, sReq.UserId)
	if rErr != nil {
		logrus.Error("[SendRedPacket] query db error %v", err)
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
	sMap := map[string]interface{}{}
	key := fmt.Sprintf("send_rp_id: %s", record.RpId)
	var receiveAmountList []int64
	remain := sReq.Amount
	sum := int64(0)
	for i := int64(0); i < sReq.Number; i++ {
		x := strategy.DoubleAverage(sReq.Number-i, remain)
		receiveAmountList = append(receiveAmountList, x)
		remain -= x
		sum += x
	}
	sMap[key] = receiveAmountList
	// 7. 写入发放记录,可以判断一下重复error

	// 8. 发送延迟消息，期间进行一次对账
}

func checkParams(seq model.SendRpReq) bool {
	return !(seq.UserId == "" || seq.GroupId == "" || seq.Amount <= 0 || (seq.Number*seq.Amount) <= 1 || seq.BizOutNo == "")
}

func QuerySendRecords(c *gin.Context) {

}

func ReceiveRedPacket(c *gin.Context) {

}

func QueryReceiveRecords(c *gin.Context) {

}
