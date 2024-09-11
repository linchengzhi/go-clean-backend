package cerror

import (
	"fmt"
)

type CustomError struct {
	code int
	msg  string
	err  error
}

func (e *CustomError) Error() string {
	if e.err != nil {
		return fmt.Sprintf("code: %d, msg: %s, err: %v", e.code, e.msg, e.err)
	}
	return fmt.Sprintf("code: %d, msg: %s", e.code, e.msg)
}

func NewError(code int, msg string) *CustomError {
	return &CustomError{
		code: code,
		msg:  msg,
	}
}

func (e *CustomError) WithErr(err error) *CustomError {
	return &CustomError{
		code: e.code,
		msg:  e.msg,
		err:  err,
	}
}

func (e *CustomError) AddMsg(msg string) *CustomError {
	return &CustomError{
		code: e.code,
		msg:  e.msg + " " + msg,
		err:  e.err,
	}
}

func (e *CustomError) GetCode() int {
	return e.code
}

func (e *CustomError) GetMsg() string {
	return e.msg
}

func (e *CustomError) GetErr() error {
	return e.err
}
