package framework

import (
	"fmt"
	"runtime/debug"
)

type ServiceError struct {
	error
	Code int
	Msg  string
}

func NewServiceError(code int, msg string, args ...any) error {
	msg = fmt.Sprintf(msg, args...)
	log.Errorf("err: %v\nstack:\n%s", msg, debug.Stack())
	return &ServiceError{
		Code: code,
		Msg:  msg,
	}
}
