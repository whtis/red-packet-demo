package kv

import (
	"context"
	"fmt"
	"github.com/sirupsen/logrus"
)

func getKey(oKey string) string {
	return fmt.Sprintf("%s:%s", "gen:demo:", oKey)
}

func LPushRp(c context.Context, key string, val []string) error {
	lKey := getKey(key)
	//返回值是当前列表元素的数量
	num, err := redisCli.LPush(c, lKey, val).Result()
	if err != nil {
		logrus.Errorf("redis: LPushRp error %v", err)
	} else {
		logrus.Infof("lpush success, size:%v", num)
	}
	return err
}

func LLenRp(c context.Context, key string) (*int64, error) {
	lKey := getKey(key)

	rLen, rErr := redisCli.LLen(c, lKey).Result()
	if rErr != nil {
		logrus.Errorf("redis: get key error %v", lKey)
		return nil, rErr
	}
	return &rLen, nil
}

// LPop 领红包的时候要用
func LPop(c context.Context, key string) (int64, error) {
	lKey := getKey(key)
	val, err := redisCli.LPop(c, lKey).Int64()
	if err != nil {
		logrus.Errorf("redis: lPop error %v", lKey)
		return 0, err
	}
	return val, nil
}
