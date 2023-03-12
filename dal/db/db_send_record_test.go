package db

import (
	"context"
	"testing"
)

func TestQuerySendRecordByBizOutNoAndUserId(t *testing.T) {
	InitDB()

	id, err := QuerySendRecordByBizOutNoAndUserId(context.Background(), "123", "234")
	if err != nil {
		return
	}
	t.Log(id)
}
