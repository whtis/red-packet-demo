package mq

import (
	"context"
	"fmt"
	"github.com/apache/rocketmq-client-go/v2"
	"github.com/apache/rocketmq-client-go/v2/consumer"
	"github.com/apache/rocketmq-client-go/v2/primitive"
)

var co rocketmq.PushConsumer

func InitConsumer(c context.Context) {
	co, _ = rocketmq.NewPushConsumer(
		consumer.WithNameServer([]string{"192.168.2.88:9876"}), // 接入点地址
		consumer.WithConsumerModel(consumer.Clustering),
		consumer.WithGroupName(group), // 分组名称
	)
	co.Subscribe(topic, consumer.MessageSelector{}, func(ctx context.Context, msg ...*primitive.MessageExt) (consumer.ConsumeResult, error) {
		for _, v := range msg {
			fmt.Println(string(v.Body)) // v.Body : 消息主体
		}
		return consumer.ConsumeSuccess, nil
	})
	err := co.Start()
	if err != nil {
		panic(err)
	}

}
