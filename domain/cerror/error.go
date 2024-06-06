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
	e.err = err
	return e
}

func (e *CustomError) AddMsg(msg string) *CustomError {
	e.msg = e.msg + " " + msg
	return e
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
