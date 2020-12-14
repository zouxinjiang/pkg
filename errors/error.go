package errors

import (
	"fmt"
)

type (
	Error interface {
		error
		Code() string
		Message() string
	}
	err struct {
		err     error
		code    string
		message string
	}
)

func New(code, message string) Error {
	return err{
		code:    code,
		message: message,
	}
}

func Wrap(innerErr error, code, message string) Error {
	return err{
		err:     innerErr,
		code:    code,
		message: message,
	}
}

func (self err) Code() string {
	return self.code
}

func (self err) Message() string {
	return self.message
}

func (self err) Error() string {
	return fmt.Sprintf("Code:%s  Message:%s", self.Code(), self.Message())
}
