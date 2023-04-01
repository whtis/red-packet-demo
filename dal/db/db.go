package db

import (
	"fmt"
	"ginDemo/utils"
	"github.com/sirupsen/logrus"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
)

var Rdb *gorm.DB

// 定义自己的Writer
type MyWriter struct {
	mlog *logrus.Logger
}

// 实现gorm/logger.Writer接口
func (m *MyWriter) Printf(format string, v ...interface{}) {
	logstr := fmt.Sprintf(format, v...)
	//利用loggus记录日志
	m.mlog.Info(logstr)
}

func NewMyWriter() *MyWriter {
	//配置logrus
	utils.RpLog.SetFormatter(&logrus.JSONFormatter{
		TimestampFormat: "2006-01-02 15:04:05",
	})

	return &MyWriter{mlog: &utils.RpLog}
}

func InitDB() {
	dsn := "root:jrttroot@tcp(127.0.0.1:3306)/tech?charset=utf8mb4&parseTime=True&loc=Local"
	var err error
	ormLogger := logger.New(
		//设置Logger
		NewMyWriter(),
		logger.Config{
			//设置日志级别，只有Warn以上才会打印sql
			LogLevel: logger.Info,
		},
	)

	Rdb, err = gorm.Open(mysql.Open(dsn), &gorm.Config{
		Logger: ormLogger,
	})
	if err != nil {
		panic(err)
	}
}
