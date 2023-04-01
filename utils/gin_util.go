package utils

import (
	"ginDemo/consts"
	"github.com/gin-gonic/gin"
	"github.com/sirupsen/logrus"
	"net/http"
)

func RetJson(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{
		"code": consts.Success.Code,
		"msg":  consts.Success.Msg,
	})
}

func RetJsonWithData(c *gin.Context, data string) {
	c.JSON(http.StatusOK, gin.H{
		"code": consts.Success.Code,
		"msg":  consts.Success.Msg,
		"data": data,
	})
}

func RetErrJson(c *gin.Context, rErr consts.RError) {
	c.JSON(http.StatusOK, gin.H{
		"code": rErr.Code,
		"msg":  rErr.Msg,
	})
}

type LogrusHook struct {
}

// Levels 设置所有的日志等级都走这个钩子
func (hook *LogrusHook) Levels() []logrus.Level {
	return logrus.AllLevels
}

// Fire 修改其中的数据，或者进行其他操作
func (hook *LogrusHook) Fire(entry *logrus.Entry) error {
	entry.Data["request_id"] = GetNewUid()
	return nil
}
