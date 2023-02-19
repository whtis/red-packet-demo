package kv

import (
	"context"
	"fmt"
	"github.com/redis/go-redis/v9"
	"github.com/sirupsen/logrus"
)

var redisCli *redis.Client

func InitRedis(ctx context.Context) {
	redisCli = redis.NewClient(&redis.Options{
		Addr:     "192.168.2.88:6379",
		Password: "", // 没有密码，默认值
		DB:       0,  // 默认DB 0
	})

	_, err := redisCli.Ping(ctx).Result()
	if err != nil {
		panic(fmt.Sprintf("连接redis出错，错误信息：%v", err))
	}
	logrus.Info("成功连接redis")

}
