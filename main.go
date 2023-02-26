package main

import (
	"ginDemo/dal/db"
	"ginDemo/handler"
	"ginDemo/service"
	"github.com/gin-gonic/gin"
)

func main() {
	db.InitDB()
	r := gin.Default()
	register(r)
	_ = r.Run() // listen and serve on 0.0.0.0:8080
}

func register(r *gin.Engine) {

	r.GET("/ping", handler.Demo)

	r.GET("/gin_demo/package_infos/:user_id", handler.QueryByUserId)

	r.POST("/gin_demo/package_infos", handler.InsertRecord)

	// 发放红包接口
	r.POST("/red-packet/send", service.SendRedPacket)
	// 查询发放记录
	r.GET("/red-packet/send/query", service.QuerySendRecords)
	// 领取红包接口
	r.POST("/red-packet/receive", service.ReceiveRedPacket)
	// 查询领取红包记录
	r.POST("/red-packet/receive/query", service.QueryReceiveRecords)

}
