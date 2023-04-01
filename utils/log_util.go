package utils

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
	"github.com/sirupsen/logrus"
	"io"
	"os"
	"path"
	"time"
)

// 定义一个实例
var RpLog logrus.Logger

// 设置用户uid，hook中使用的上
var uid string

func init() {
	setUid()
	Logger()
}

func Logger() {
	//当前时间
	nowTime := time.Now()
	//获取日志文件存储的目录，这里我采用的是自己封装的一个获取配置文件的方法,可以看我上一篇viper获取配置信息的文章
	logFilePath := "./log"
	//创建文件夹
	if err := os.MkdirAll(logFilePath, os.ModePerm); err != nil {
		Error(err.Error())
	}
	//文件名称
	logFileName := nowTime.Format("2006-01-02") + ".log"
	//日志文件地址拼接
	fileName := path.Join(logFilePath, logFileName)
	//fmt.Println("文件名称："+fileName)
	if _, err := os.Stat(fileName); err != nil {
		Error("检测文件：" + err.Error())
		_, err := os.Create(fileName)
		if err != nil {
			Error(err.Error())
		}
	}
	//打开文件
	src, err := os.OpenFile(fileName, os.O_CREATE|os.O_RDWR|os.O_APPEND, os.ModeAppend|os.ModePerm)
	if err != nil {
		Error("write file log error", err)
	}
	//实例化
	RpLog = *logrus.New()

	writers := []io.Writer{
		src,
		os.Stdout}
	//同时写文件和屏幕
	fileAndStdoutWriter := io.MultiWriter(writers...)
	//设置输出
	RpLog.Out = fileAndStdoutWriter

	//这里我觉得应该是交给需要封装的方法去确认使用什么等级的日志和什么格式
	//设置日志级别
	//logger.SetLevel(logrus.InfoLevel)
	////设置日志格式 json格式
	//logger.SetFormatter(&logrus.JSONFormatter{
	// TimestampFormat: "2006-01-02 15:04:05",
	//})
	RpLog.AddHook(&LogrusHook{})
	RpLog.SetFormatter(&logrus.JSONFormatter{
		TimestampFormat: "2006-01-02 15:04:05",
	})
}

// 获取gin日志的中间件
func GinLogMiddleware() gin.HandlerFunc {

	RpLog.SetFormatter(&logrus.JSONFormatter{})
	RpLog.SetLevel(logrus.InfoLevel)

	return func(c *gin.Context) {
		c.Next()
		method := c.Request.Method
		reqUrl := c.Request.RequestURI
		statusCode := c.Writer.Status()
		clientIP := c.ClientIP()
		RpLog.WithFields(logrus.Fields{
			"method":      method,
			"uri":         reqUrl,
			"status_code": statusCode,
			"client_ip":   clientIP,
		}).Info()
		setUid()
	}
}

// 获取gin日志的中间件
func GinConsoleLogMiddleware() gin.HandlerFunc {
	loggerInfo := *logrus.New()
	loggerInfo.SetFormatter(&logrus.JSONFormatter{})
	loggerInfo.SetLevel(logrus.InfoLevel)

	return func(c *gin.Context) {
		c.Next()
		method := c.Request.Method
		reqUrl := c.Request.RequestURI
		statusCode := c.Writer.Status()
		clientIP := c.ClientIP()
		loggerInfo.WithFields(logrus.Fields{
			"method":      method,
			"uri":         reqUrl,
			"status_code": statusCode,
			"client_ip":   clientIP,
		}).Info()
		setUid()
	}
}

func Info(args ...interface{}) {
	RpLog.WithFields(logrus.Fields{}).Info(args)
}

func Infof(format string, args ...interface{}) {
	RpLog.WithFields(logrus.Fields{}).Infof(format, args)
}

func Warn(args ...interface{}) {
	RpLog.WithFields(logrus.Fields{}).Warn(args)
}

func Warnf(format string, args ...interface{}) {
	RpLog.WithFields(logrus.Fields{
		"data": Json2String(fmt.Sprintf(format, args)),
	}).Warn()
}

func Error(args ...interface{}) {
	RpLog.WithFields(logrus.Fields{}).Error(args)
}

func Errorf(format string, args ...interface{}) {
	RpLog.WithFields(logrus.Fields{}).Errorf(format, args)
}

func setUid() {
	uid = uuid.New().String()
}

func GetNewUid() string {
	return uid
}
