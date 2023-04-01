package utils

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
	"github.com/sirupsen/logrus"
	"os"
	"path"
	"time"
)

var RpLog = *logrus.New()

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
	logFilePath := "./logs"
	// 如果没有获取到配置文件的话那就是直接代码写死一个文件地址
	//if len(logFilePath) <= 0 {
	//	//获取当前目前的地址，也就是项目的根目录
	//	if dir, err := os.Getwd(); err == nil {
	//		logFilePath = dir + "/logs/"
	//	}
	//}
	//创建文件夹
	if err := os.MkdirAll(logFilePath, os.ModePerm); err != nil {
		fmt.Println(err.Error())
	}
	//文件名称
	logFileName := nowTime.Format("2006-01-02") + ".log"
	//日志文件地址拼接
	fileName := path.Join(logFilePath, logFileName)
	//fmt.Println("文件名称："+fileName)
	if _, err := os.Stat(fileName); err != nil {
		fmt.Println("检测文件：" + err.Error())
		_, err := os.Create(fileName)
		if err != nil {
			fmt.Println(err.Error())
		}
	}
	//打开文件
	src, err := os.OpenFile(fileName, os.O_CREATE|os.O_RDWR|os.O_APPEND, os.ModeAppend|os.ModePerm)
	if err != nil {
		fmt.Println("write file log error", err)
	}
	//设置输出
	RpLog.Out = src
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
	RpLog.Out = os.Stdout

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

func Info(args ...interface{}) {
	RpLog.WithFields(logrus.Fields{
		"data": Json2String(fmt.Sprint(args...)),
	}).Info()
}

func Infof(format string, args ...interface{}) {
	RpLog.WithFields(logrus.Fields{
		"data": Json2String(fmt.Sprintf(format, args)),
	}).Info()
}

func Warn(args ...interface{}) {
	RpLog.WithFields(logrus.Fields{
		"data": Json2String(fmt.Sprint(args...)),
	}).Warn()
}

func Warnf(format string, args ...interface{}) {
	RpLog.WithFields(logrus.Fields{
		"data": Json2String(fmt.Sprintf(format, args)),
	}).Warn()
}

func Error(args ...interface{}) {
	RpLog.WithFields(logrus.Fields{
		"data": Json2String(fmt.Sprint(args...)),
	}).Error()
}

func Errorf(format string, args ...interface{}) {
	RpLog.WithFields(logrus.Fields{
		"data": Json2String(fmt.Sprintf(format, args)),
	}).Error()
}

func setUid() {
	uid = uuid.New().String()
}

func GetNewUid() string {
	return uid
}
