package model

type PackageInfo struct {
	UserId        string `json:"user_id"`
	ReceiveAmount int64  `json:"receive_amount"`
}

type SendRpReq struct {
	UserId   string `json:"user_id"`
	GroupId  string `json:"group_id"`
	Amount   int64  `json:"amount"` // 传的是分
	Number   int64  `json:"number"`
	BizOutNo string `json:"biz_out_no"`
}

type ReceiveRpReq struct {
	UserId   string `json:"user_id"`
	GroupId  string `json:"group_id"`
	BizOutNo string `json:"biz_out_no"`
	RpId     string `json:"rp_id"`
}

type QuerySendRecordReq struct {
	UserId  string `json:"user_id"`
	GroupId string `json:"group_id"`

	Size   int64  `json:"size"`   // 默认10条
	Cursor string `json:"cursor"` // 默认从"0"
}

type QuerySendRecordReqByPage struct {
	UserId  string `json:"user_id"`
	GroupId string `json:"group_id"`

	Size  int64 `json:"size"` // 默认20条
	Page  int64 `json:"page"` // 从第1页开始，默认是1
	Total int64 `json:"total"`
}

type QuerySendRecordRespByPage struct {
	RpSendRecordList []*RpSendRecord `json:"rp_send_record_list"`
	Total            int64           `json:"total"`
}

type QueryReceiveRecordReq struct {
	UserId  string `json:"user_id"`
	GroupId string `json:"group_id"`

	Size  int64 `json:"size"`
	Page  int64 `json:"page"`
	Total int64 `json:"total"`
}

type QuerySendRecordResp struct {
	RpSendRecordList []*RpSendRecord `json:"rp_send_record_list"`
	HasMore          bool            `json:"has_more"`
	Cursor           string          `json:"cursor"`
}
