package mq

import (
	"context"
	"encoding/json"
	"github.com/apache/rocketmq-client-go/v2"
	"github.com/apache/rocketmq-client-go/v2/primitive"
	"github.com/apache/rocketmq-client-go/v2/producer"
	"github.com/sirupsen/logrus"
)

var p rocketmq.Producer

var topic = "gen.demo.redPacket"

func InitProduct(c context.Context) {
	p, _ = rocketmq.NewProducer(
		producer.WithNameServer([]string{"localhost:9876"}), // 接入点地址
		producer.WithRetry(2),                               // 重试次数
		producer.WithGroupName("ProductGroup"),              // 分组名称
	)
	err := p.Start()
	if err != nil {
		panic("")
	}
}

func SendRpDelay(c context.Context, val interface{}, retryTimes int64) error {
	if retryTimes > 3 {
		logrus.Errorf("SendRpDelay error, reason: reTry times max")
		// 出问题（没发出去的消息，我得记录下来----mysql----id）--db出问题-》写redis--local file
		return nil
	}
	// 发送的消息
	mResult, _ := json.Marshal(val)

	msg := &primitive.Message{
		Topic: topic,
		Body:  mResult,
	}
	msg.WithTag("my-tag")
	msg.WithKeys([]string{"my-key"})
	// 发送消息
	// 作业，看下延迟消息怎么发
	result, err := p.SendSync(c, msg)
	if err != nil {
		logrus.Errorf("SendRpDelay: send message error %v", err)
		retryTimes = retryTimes + 1
		return SendRpDelay(c, val, retryTimes)
	}
	logrus.Infof("send rp delay message success! msg:%v", result.MsgID)
	return nil
}
