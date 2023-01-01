package main

import (
	"ginDemo/handler"
	"github.com/gin-gonic/gin"
)

func register(r *gin.Engine) {

	r.GET("/ping", handler.Demo)

	r.GET("/gin_demo/package_infos/:user_id", handler.QueryByUserId)

	r.POST("/gin_demo/package_infos", handler.InsertRecord)

}
