package main

import (
	"context"
	"ginDemo/dal/db"
	"ginDemo/dal/kv"
	"ginDemo/handler"
	"ginDemo/service"
	"ginDemo/utils"
	"github.com/gin-gonic/gin"
)

func main() {
	db.InitDB()
	kv.InitRedis(context.Background())
	r := gin.New()
	r.Use(gin.Recovery())
	r.Use(utils.GinLogMiddleware())
	//r.Use(utils.GinConsoleLogMiddleware())
	register(r)
	_ = r.Run() // listen and serve on 0.0.0.0:8080
}

func register(r *gin.Engine) {
	// 路由

	r.GET("/s/wd=?", handler.Demo)

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
	// 导出发放记录
	r.POST("/red-packet/send/export", service.ExportSendRecords)

}
