package db

import (
	"ginDemo/model"
	"github.com/go-playground/assert/v2"
	"github.com/google/uuid"
	"testing"
	"time"
)

func TestInsertSendRecord(t *testing.T) {
	InitDB()
	m := &model.RpSendRecord{
		UserId:      "9527",
		GroupChatId: "g001",
		BizOutNo:    uuid.New().String(),
		Amount:      1000, //10元
		Number:      10,
		ExpireTime:  time.Now().Add(1 * time.Hour),
		SendTime:    time.Now(),
		CreateTime:  time.Now(),
		ModifyTime:  time.Now(),
	}
	id, err := InsertSendRecord(nil, m)
	t.Log(id)
	assert.Equal(t, err, nil)

}
