package kv

import (
	"context"
	"fmt"
	"testing"
	"time"
)

func TestInitRedis(t *testing.T) {
	InitRedis(context.Background())
}

func TestLPushRp(t *testing.T) {
	ctx := context.Background()
	InitRedis(ctx)
	key := "lllll"
	vals := []string{"1"}
	_ = LPushRp(ctx, key, vals)
	intVal1, rErr1 := LPop(ctx, key)
	intVal2, rErr2 := LPop(ctx, key)
	t.Log(intVal1, intVal2, rErr1, rErr2)

}

func TestLLen(t *testing.T) {
	nowTime := time.Date(2999, 1, 1, 0, 0, 0, 0, time.Local).Unix()
	fmt.Print(nowTime)
	ctx := context.Background()
	InitRedis(ctx)

	rpId := "c32a253d7c5d4966b572dc0c4ef1aa17"

	count, _ := LLenRp(ctx, rpId)
	t.Log(count)
}
