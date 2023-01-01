package consts

type RError struct {
	Code int64
	Msg  string
}

var (
	Success = buildErrorCode(0, "success")

	RecordNotFound = buildErrorCode(-1, "record not found")

	ParamsError = buildErrorCode(-2, "bind error")

	InsertError = buildErrorCode(-3, "insert data error")
)

func buildErrorCode(code int64, msg string) RError {
	return RError{
		Code: code,
		Msg:  msg,
	}
}
