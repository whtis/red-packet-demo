package kv

import (
	"context"
	"testing"
)

func TestInitRedis(t *testing.T) {
	InitRedis(context.Background())
}
