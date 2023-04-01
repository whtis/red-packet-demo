package kv

import (
	"context"
	"fmt"
	"ginDemo/utils"
	"github.com/redis/go-redis/v9"
)

var redisCli *redis.Client

func InitRedis(ctx context.Context) {
	redisCli = redis.NewClient(&redis.Options{
		Addr:     "127.0.0.1:6379",
		Password: "redispw", // 没有密码，默认值
		DB:       0,         // 默认DB 0
	})

	_, err := redisCli.Ping(ctx).Result()
	if err != nil {
		panic(fmt.Sprintf("连接redis出错，错误信息：%v", err))
	}
	utils.Info("成功连接redis")

}
