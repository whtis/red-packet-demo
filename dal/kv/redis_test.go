package kv

import (
	"context"
	"testing"
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
