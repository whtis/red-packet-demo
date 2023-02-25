package consts

type RError struct {
	Code int64
	Msg  string
}

var (
	Success = buildErrorCode(0, "success")

	RecordNotFound = buildErrorCode(-1, "record not found")

	BindError = buildErrorCode(-2, "bind error")

	ParamError = buildErrorCode(-5, "param error")

	InsertError = buildErrorCode(-3, "insert data error")

	ServiceBusy = buildErrorCode(-110, "service busy")

	RpExpiredError = buildErrorCode(-6, "red packet expired")

	RpReceivedError = buildErrorCode(-4, "all the red packets are received")
)

func buildErrorCode(code int64, msg string) RError {
	return RError{
		Code: code,
		Msg:  msg,
	}
}
