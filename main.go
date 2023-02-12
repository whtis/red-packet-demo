package main

import (
	"ginDemo/dal"
	"ginDemo/handler"
	"github.com/gin-gonic/gin"
)

func main() {
	dal.InitDB()
	r := gin.Default()
	register(r)
	_ = r.Run() // listen and serve on 0.0.0.0:8080
}

func register(r *gin.Engine) {

	r.GET("/ping", handler.Demo)

	r.GET("/gin_demo/package_infos/:user_id", handler.QueryByUserId)

	r.POST("/gin_demo/package_infos", handler.InsertRecord)

	// 发放红包接口
	r.POST("/red-packet/send", handler.SendRedPacket)
	// 查询发放记录
	r.GET("/red-packet/send/query", handler.QuerySendRecords)
	// 领取红包接口
	r.POST("/red-packet/receive", handler.ReceiveRedPacket)
	// 查询领取红包记录
	r.POST("/red-packet/receive/query", handler.QueryReceiveRecords)

}
