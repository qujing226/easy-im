package xerr

import "github.com/zeromicro/x/errors"

func New(code int, msg string) error {
	return errors.New(code, msg)
}

func NewDBErr() error {
	return New(DB_ERROR, ErrMsg(DB_ERROR))
}

func NewServerCommonErr() error {
	return New(SERVER_COMMON_ERROR, ErrMsg(SERVER_COMMON_ERROR))
}
