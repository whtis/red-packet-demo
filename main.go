package main

import (
	"github.com/gin-gonic/gin"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
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
	CreateTime  time.Time
	ModifyTime  time.Time
}

func main() {

	dsn := "root:root2023@tcp(127.0.0.1:3306)/tech?charset=utf8mb4&parseTime=True&loc=Local"
	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})

	if err != nil {
		panic(err)
	}

	sqlDB, _ := db.DB()

	// SetMaxIdleConns 设置空闲连接池中连接的最大数量
	sqlDB.SetMaxIdleConns(10)

	// SetMaxOpenConns 设置打开数据库连接的最大数量。
	sqlDB.SetMaxOpenConns(100)

	// SetConnMaxLifetime 设置了连接可复用的最大时间。
	sqlDB.SetConnMaxLifetime(time.Hour)

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

		err := db.Table("rp_receive_record").Where("user_id = ?", userId).First(&record).Error
		//s := fmt.Sprintf("select * from rp_receive_record where user_id = '%s'", userId)
		//err := db.QueryRow(s).Scan(&record.Id, &record.UserId, &record.GroupChatId, &record.RpId, &record.Amount, &record.CreateTime, &record.ModifyTime)
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
		record := &RpReceiveRecord{
			UserId:      packageInfo.UserId,
			GroupChatId: "insert_group002",
			RpId:        "insert_rp_002",
			Amount:      packageInfo.ReceiveAmount,
			CreateTime:  time.Now(),
			ModifyTime:  time.Now(),
		}
		//r, err := db.Exec(s, packageInfo.UserId, "insert_group001", "insert_rp_001", packageInfo.ReceiveAmount, time.Now())
		//s := "insert into rp_receive_record (user_id,group_chat_id,rp_id,amount,create_time) values(?,?,?,?,?)"
		err = db.Table("rp_receive_record").Create(&record).Error
		if err != nil {
			log.Printf("insert data err: %v\n", err)
			c.JSON(http.StatusOK, gin.H{
				"code":    "-3",
				"message": "insert data error",
			})
			return
		} else {
			log.Printf("i: %v\n", record.Id)
			c.JSON(http.StatusOK, gin.H{
				"code":       "0",
				"message":    "success",
				"primary_id": record.Id,
			})
		}

	})

	_ = r.Run() // listen and serve on 0.0.0.0:8080
}
