package errors

import (
	"fmt"
	"runtime"
	"strings"
)

type (
	Error interface {
		error
		Code() string
		Message() string
		File() string
		FuncName() string
		Line() int
		Pkg() string
		ErrorDetail() string
	}
	err struct {
		err     error
		code    string
		message string
		file    string
		pkg     string
		line    int
		fun     string
	}
)

func New(code, message string) Error {
	file, pkg, funcName, line := getPositionInfo()
	return err{
		code:    code,
		message: message,
		file:    file,
		pkg:     pkg,
		line:    line,
		fun:     funcName,
	}
}

func Wrap(innerErr error, code, message string) Error {
	file, pkg, funcName, line := getPositionInfo()
	return err{
		err:     innerErr,
		code:    code,
		message: message,
		file:    file,
		pkg:     pkg,
		line:    line,
		fun:     funcName,
	}
}

func getPositionInfo() (file, pkg, funcName string, line int) {
	var pc uintptr
	pc, file, line, _ = runtime.Caller(2)
	pkg = ""
	funcName = runtime.FuncForPC(pc).Name()
	tmp := strings.Split(funcName, ".")
	if len(tmp) >= 2 {
		funcName = tmp[len(tmp)-1]
		pkg = strings.Join(tmp[:len(tmp)-1], ".")
	} else if len(tmp) == 1 {
		funcName = tmp[0]
		pkg = tmp[0]
	}
	return file, pkg, funcName, line
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

func (self err) File() string {
	return self.file
}

func (self err) Pkg() string {
	return self.pkg
}

func (self err) FuncName() string {
	return self.fun
}

func (self err) Line() int {
	return self.line
}

func (self err) ErrorDetail() string {
	return fmt.Sprintf("Pkg:%s Func:%s Line:%d Code:%s  Message:%s", self.Pkg(), self.FuncName(), self.Line(), self.Code(), self.Message())
}
