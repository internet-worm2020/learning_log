package pkg

import (
	"fmt"
)

// Error 表示一个错误的详细信息.
type Error struct {
	BusinessCode ResCode `json:"code"`
	Message      string  `json:"message"`
}

// NewError 根据状态码、错误码、错误描述创建一个Error.
func NewError(businessCode ResCode, msg string) *Error {
	return &Error{
		BusinessCode: businessCode,
		Message:      msg,
	}
}

// NewErrorAutoMsg 根据状态码、错误码创建一个Error.
func NewErrorAutoMsg(businessCode ResCode) *Error {
	msg := businessCode.Msg()
	return NewError(businessCode, msg)
}

// WithErr 把内部的err放到 Error 中.
func (e *Error) WithErr(err error) *Error {
	if err == nil {
		return e
	}
	return &Error{
		BusinessCode: e.BusinessCode,
		Message:      fmt.Sprintf("%s: %s", e.Message, err.Error()),
	}
}
