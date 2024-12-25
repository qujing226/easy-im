package xerr

var codeText = map[int]string{
	ServerCommonError: "server fatal, try again later",
	RequestParamError: "params wrong",
	DbError:           "database busy, try again later",
}

func ErrMsg(code int) string {
	if msg, ok := codeText[code]; ok {
		return msg
	}
	return codeText[ServerCommonError]
}
