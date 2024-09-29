package xerr

var codeText = map[int]string{
	SERVER_COMMON_ERROR: "server fatal, try again later",
	REQUEST_PARAM_ERROR: "params wrong",
	DB_ERROR:            "database busy, try again later",
}

func ErrMsg(code int) string {
	if msg, ok := codeText[code]; ok {
		return msg
	}
	return codeText[SERVER_COMMON_ERROR]
}
