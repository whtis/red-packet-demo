package mq

import (
	"context"
	"github.com/apache/rocketmq-client-go/v2"
	"github.com/apache/rocketmq-client-go/v2/consumer"
)

var co rocketmq.PushConsumer

func InitConsumer(c context.Context) {
	co, _ = rocketmq.NewPushConsumer(
		consumer.WithNameServer([]string{"localhost:9876"}), // 接入点地址
		consumer.WithConsumerModel(consumer.Clustering),
		consumer.WithGroupName("ConsumerGroup"), // 分组名称
	)
	err := co.Start()
	if err != nil {
		panic("")
	}
}
