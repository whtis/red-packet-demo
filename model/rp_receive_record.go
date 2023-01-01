package model

import "time"

type RpReceiveRecord struct {
	Id          int64
	UserId      string
	GroupChatId string
	RpId        string
	Amount      int64
	CreateTime  time.Time
	ModifyTime  time.Time
}
