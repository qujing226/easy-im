package logic

import "easy-chat/pkg/xerr"

var (
	ErrPhoneNotFound = xerr.New(xerr.ServerCommonError, "user not found")
	ErrIDNotFound    = xerr.New(xerr.ServerCommonError, "user id not found")
	ErrUserPwdErr    = xerr.New(xerr.ServerCommonError, "password is wrong")

	ErrParamError = xerr.New(xerr.RequestParamError, "params error")
)
