package handler

import (
	"encoding/json"
	"ginDemo/consts"
	"ginDemo/dal/db"
	"ginDemo/model"
	"ginDemo/utils"
	"github.com/gin-gonic/gin"
	"log"
	"net/http"
	"time"
)

func Demo(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{
		"message": "pong",
	})
}

func QueryByUserId(c *gin.Context) {
	userId := c.Param("user_id")
	log.Printf("get userid from request %v", userId)
	record, err := db.QueryByUserId(userId)
	if err != nil {
		log.Printf("can not find userId amount, userId: %v, err: %v", userId, err)
		utils.RetErrJson(c, consts.BindError)
	} else {
		//log.Printf("find userId amount, userId: %v, amount: %v", userId, value)
		rStr, _ := json.Marshal(record)
		utils.RetJsonWithData(c, string(rStr))
		return
	}
}

func InsertRecord(c *gin.Context) {
	var packageInfo model.PackageInfo
	err := c.BindJSON(&packageInfo)
	if err != nil {
		log.Printf("bind package info error %v", err)
		utils.RetErrJson(c, consts.BindError)
		return
	}
	record := &model.RpReceiveRecord{
		UserId:      packageInfo.UserId,
		GroupChatId: "insert_group002",
		RpId:        "insert_rp_002",
		Amount:      packageInfo.ReceiveAmount,
		CreateTime:  time.Now(),
		ModifyTime:  time.Now(),
	}
	id, iErr := db.InsertRecord(record)

	if iErr != nil {
		log.Printf("insert data err: %v\n", err)
		utils.RetErrJson(c, consts.InsertError)
		return
	} else {
		log.Printf("i: %v\n", id)
		c.JSON(http.StatusOK, gin.H{
			"code":       "0",
			"message":    "success",
			"primary_id": id,
		})
	}
}
