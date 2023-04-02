package service

import (
	"ginDemo/consts"
	"ginDemo/dal/db"
	"ginDemo/model"
	"ginDemo/utils"
	"github.com/gin-gonic/gin"
	"math"
	"strconv"
)

// 游标模式
func QuerySendRecords(c *gin.Context) {
	// 1. 参数绑定
	var sReq model.QuerySendRecordReq
	if v, b := c.Get("user_id"); b {
		utils.Info("xxxxxx------------:%v", v)
	}

	err := c.BindJSON(&sReq)
	if err != nil {
		utils.Error("[QuerySendRecords] bind req json error")
		utils.RetErrJson(c, consts.BindError)
		return
	}
	// 2. 参数检查
	ok := checkSendRecordParams(sReq)
	if !ok {
		utils.Errorf("[ReceiveRedPacket] check params error, rReq: %v", utils.Json2String(sReq))
		utils.RetErrJson(c, consts.ParamError)
		return
	}
	// 3. 查数据 重点：找到cursor对应的字段
	if sReq.Cursor == "" {
		sReq.Cursor = strconv.FormatInt(math.MaxInt64, 10)
	}
	if sReq.Size == 0 {
		sReq.Size = 10
	}
	temp := sReq.Size
	sReq.Size = sReq.Size + 1

	hasMore := false
	records, rErr := db.QuerySendRecordByCond(c, sReq)
	if rErr != nil {
		// todo
	}
	if len(records) > int(temp) {
		hasMore = true
	}
	// 返回
	result := make([]*model.RpSendRecord, 0)
	curInt := 0
	for i, record := range records {
		r := &model.RpSendRecord{
			UserId:      record.UserId,
			GroupChatId: record.GroupChatId,
			// todo 填充前端展示的数据
		}
		curInt = i
		result = append(result, r)
	}
	retVal := &model.QuerySendRecordResp{
		RpSendRecordList: result,
		HasMore:          hasMore,
		Cursor:           strconv.FormatInt(records[curInt].CreateTime.Unix(), 10),
	}
	utils.RetJsonWithData(c, utils.Json2String(retVal))
}

// 分页
func QuerySendRecordsByPage(c *gin.Context) {
	// 1. 参数绑定
	var sReq model.QuerySendRecordReqByPage
	err := c.BindJSON(&sReq)
	if err != nil {
		utils.Error("[QuerySendRecords] bind req json error")
		utils.RetErrJson(c, consts.BindError)
		return
	}
	// 2. 参数检查
	//ok := checkSendRecordParams(sReq)
	//if !ok {
	//	utils.Errorf("[ReceiveRedPacket] check params error, rReq: %v", utils.Json2String(sReq))
	//	utils.RetErrJson(c, consts.ParamError)
	//	return
	//}

	sReq.Page = sReq.Page - 1

	records, rErr := db.QuerySendRecordByCondPage(c, sReq)
	if rErr != nil {
		// todo
	}
	count, cErr := db.CountSendRecordByCondPage(c, sReq)
	if cErr != nil {

	}

	// 返回数据
	result := make([]*model.RpSendRecord, 0)
	for _, record := range records {
		r := &model.RpSendRecord{
			UserId:      record.UserId,
			GroupChatId: record.GroupChatId,
			// todo 填充前端展示的数据
		}
		result = append(result, r)
	}
	retVal := &model.QuerySendRecordRespByPage{
		RpSendRecordList: result,
		Total:            count,
	}
	utils.RetJsonWithData(c, utils.Json2String(retVal))
}

func checkSendRecordParams(req model.QuerySendRecordReq) bool {
	if req.UserId == "" {
		return false
	}
	return true
}
