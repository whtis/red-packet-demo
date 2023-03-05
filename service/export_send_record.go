package service

import (
	"fmt"
	"ginDemo/consts"
	"ginDemo/dal/db"
	"ginDemo/model"
	"ginDemo/utils"
	"github.com/gin-gonic/gin"
	"github.com/sirupsen/logrus"
	"github.com/xuri/excelize/v2"
	"time"
)

func ExportSendRecords(c *gin.Context) {
	// 1. 参数绑定
	var sReq model.ExportSendRecordReq
	err := c.BindJSON(&sReq)
	if err != nil {
		logrus.Error("[QuerySendRecords] bind req json error")
		utils.RetErrJson(c, consts.BindError)
		return
	}
	// 2. 参数检查
	//ok := checkSendRecordParams(sReq)
	//if !ok {
	//	logrus.Errorf("[ReceiveRedPacket] check params error, rReq: %v", utils.Json2String(sReq))
	//	utils.RetErrJson(c, consts.ParamError)
	//	return
	//}

	records, rErr := db.ExportSendRecords(c, sReq)
	if rErr != nil {
		// todo
	}

	eErr := exportRecords(c, records)
	if eErr != nil {
		utils.RetErrJson(c, consts.ServiceBusy)
	}
	utils.RetJson(c)
}

func exportRecords(c *gin.Context, records []*model.RpSendRecord) error {
	f := excelize.NewFile()
	defer func() {
		if err := f.Close(); err != nil {
			fmt.Println(err)
		}
	}()
	// Create a new sheet.
	sheetTitle := "send_record"
	index, err := f.NewSheet(sheetTitle)
	if err != nil {
		fmt.Println(err)
		return err
	}
	// 设置表头
	_ = f.SetCellValue(sheetTitle, "A1", "用户id")
	_ = f.SetCellValue(sheetTitle, "B1", "群聊id")
	_ = f.SetCellValue(sheetTitle, "C1", "发放金额")
	_ = f.SetCellValue(sheetTitle, "D1", "发放时间")
	// 设置数据
	dataIndex := 2
	for _, record := range records {
		_ = f.SetCellValue(sheetTitle, fmt.Sprintf("A%d", dataIndex), record.UserId)
		_ = f.SetCellValue(sheetTitle, fmt.Sprintf("B%d", dataIndex), record.GroupChatId)
		_ = f.SetCellValue(sheetTitle, fmt.Sprintf("C%d", dataIndex), record.Amount)
		_ = f.SetCellValue(sheetTitle, fmt.Sprintf("D%d", dataIndex), record.CreateTime) // todo 时间需要转换成固定格式
		dataIndex = dataIndex + 1
	}
	// Set active sheet of the workbook.
	f.SetActiveSheet(index)
	// Save spreadsheet by the given path.
	if err := f.SaveAs(fmt.Sprintf("./../sendRerecord_%d.xlsx", time.Now().Unix())); err != nil {
		return err
	}
	return nil
}
