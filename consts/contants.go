package consts

import "time"

var (
	ExpireTime24 = time.Hour * 24
	ExpireTime12 = time.Hour * 12
)

var (
	RpStatusSend    = 5
	RpStatusOver    = 10
	RpStatusExpired = 15
)
