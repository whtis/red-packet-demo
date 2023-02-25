package mq

import (
	"context"
	"testing"
	"time"
)

func TestInitProduct(t *testing.T) {
	ctx := context.Background()
	InitConsumer(ctx)
	time.Sleep(10 * time.Second)
	InitProduct(ctx)

	for i := 0; i < 10; i++ {
		_ = SendRpDelay(ctx, "hello world!", 0)
	}
}
