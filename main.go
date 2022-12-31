package main

import (
	"github.com/gin-gonic/gin"
	"log"
	"net/http"
)

type PackageInfo struct {
	UserId        string `json:"user_id"`
	ReceiveAmount int64  `json:"receive_amount"`
}

func main() {
	m := map[string]int64{
		"u001": 10,
		"u002": 5,
		"u003": 20,
	}

	r := gin.Default()

	r.GET("/ping", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{
			"message": "pong",
		})
	})

	r.GET("/gin_demo/package_infos/:user_id", func(c *gin.Context) {
		userId := c.Param("user_id")
		log.Printf("get userid from request %v", userId)

		if value, ok := m[userId]; ok {
			log.Printf("find userId amount, userId: %v, amount: %v", userId, value)
			c.JSON(http.StatusOK, gin.H{
				"receive_amount": value,
			})
			return
		} else {
			log.Printf("can not find userId amount, userId: %v", userId)
			//c.JSON(http.StatusNotFound, gin.H{
			//	"message": "bad request",
			//})
			c.JSON(http.StatusOK, gin.H{
				"code":    "-1",
				"message": "record not found",
			})
		}
	})

	r.POST("/gin_demo/package_infos", func(c *gin.Context) {
		var packageInfo PackageInfo
		err := c.BindJSON(&packageInfo)
		if err != nil {
			log.Printf("bind package info error %v", err)
			c.JSON(http.StatusOK, gin.H{
				"code":    "-2",
				"message": "bind error",
			})
			return
		}
		m[packageInfo.UserId] = packageInfo.ReceiveAmount
		c.JSON(http.StatusOK, gin.H{
			"code":    "0",
			"message": "success",
		})
	})

	r.Run() // listen and serve on 0.0.0.0:8080
}
