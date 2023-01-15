package model

import "time"

type RpSendRecord struct {
	Id            int64
	UserId        string
	GroupChatId   string
	RpId          string
	BizOutNo      string
	Amount        int64
	ReceiveAmount int64
	Number        int64
	Status        int
	ExpireTime    time.Time
	SendTime      time.Time
	CreateTime    time.Time
	ModifyTime    time.Time
}
