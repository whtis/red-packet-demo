package main

import (
	"database/sql"
	"fmt"
	"github.com/gin-gonic/gin"
	_ "github.com/go-sql-driver/mysql"
	"log"
	"net/http"
	"time"
)

type PackageInfo struct {
	UserId        string `json:"user_id"`
	ReceiveAmount int64  `json:"receive_amount"`
}

type RpReceiveRecord struct {
	Id          int64
	UserId      string
	GroupChatId string
	RpId        string
	Amount      int64
	CreateTime  string
	ModifyTime  string
}

func main() {

	db, err := sql.Open("mysql", "root:root2023@tcp(127.0.0.1:3306)/tech")
	if err != nil {
		panic(err)
	}
	// See "Important settings" section.
	db.SetConnMaxLifetime(time.Minute * 3)
	db.SetMaxOpenConns(10)
	db.SetMaxIdleConns(10)

	r := gin.Default()

	r.GET("/ping/u001", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{
			"receive_amount": "10",
		})
	})

	// /gin_demo/package_infos/u001
	r.GET("/gin_demo/package_infos/:user_id", func(c *gin.Context) {
		userId := c.Param("user_id")
		log.Printf("get userid from request %v", userId)

		var record RpReceiveRecord
		s := fmt.Sprintf("select * from rp_receive_record where user_id = '%s'", userId)
		err := db.QueryRow(s).Scan(&record.Id, &record.UserId, &record.GroupChatId, &record.RpId, &record.Amount, &record.CreateTime, &record.ModifyTime)
		//s := "select * from rp_receive_record where user_id = ? and id = ?"
		//err := db.QueryRow(s, userId, 1).Scan(&record.Id, &record.UserId, &record.GroupChatId, &record.RpId, &record.Amount, &record.CreateTime, &record.ModifyTime)
		if err != nil {
			log.Printf("can not find userId amount, userId: %v, err: %v", userId, err)
			c.JSON(http.StatusOK, gin.H{
				"code":    "-1",
				"message": "record not found",
			})
		} else {
			//log.Printf("find userId amount, userId: %v, amount: %v", userId, value)
			c.JSON(http.StatusOK, gin.H{
				"receive_amount": record.Amount,
				"code":           "0",
				"message":        "success",
			})
			return
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

		s := "insert into rp_receive_record (user_id,group_chat_id,rp_id,amount,create_time) values(?,?,?,?,?)"

		r, err := db.Exec(s, packageInfo.UserId, "insert_group001", "insert_rp_001", packageInfo.ReceiveAmount, time.Now())
		if err != nil {
			log.Printf("insert data err: %v\n", err)
			c.JSON(http.StatusOK, gin.H{
				"code":    "-3",
				"message": "insert data error",
			})
			return
		} else {
			i, _ := r.LastInsertId()
			log.Printf("i: %v\n", i)
			c.JSON(http.StatusOK, gin.H{
				"code":       "0",
				"message":    "success",
				"primary_id": i,
			})
		}

	})

	_ = r.Run() // listen and serve on 0.0.0.0:8080
}
