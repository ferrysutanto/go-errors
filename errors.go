package errors

import (
	"fmt"
	"strings"
)

type Error struct {
	message string
	code    ErrCode

	fileName     string
	functionName string
	line         int

	stacks []*Error
}

func (e *Error) Code() ErrCode {
	return e.code
}

func (e *Error) Error() string {
	return e.message
}

func (e *Error) Stacks() []string {
	if e == nil {
		return make([]string, 0)
	}

	resp := make([]string, len(e.stacks))
	for k, err := range e.stacks {
		resp[k] = fmt.Sprintf("%s:%s:%d %s", err.fileName, err.functionName, err.line, err.message)
	}

	return resp
}

func (e *Error) Errors() string {
	if e == nil {
		return ""
	}

	return strings.Join(e.Stacks(), "\n")
}

func New(msg string) *Error {
	err := Error{
		message: msg,
	}

	return &err
}

func NewWithCode(msg string, code ErrCode) *Error {
	t := getTrace()

	err := Error{
		message:      msg,
		code:         code,
		fileName:     t.FileName,
		functionName: t.FunctionName,
		line:         t.Line,
	}

	return &err
}

func Newf(format string, args ...interface{}) *Error {
	return New(fmt.Sprintf(format, args...))
}

func NewfWithCode(code ErrCode, format string, args ...interface{}) *Error {
	return NewWithCode(fmt.Sprintf(format, args...), code)
}

func Wrap(err error, msg string) *Error {
	return WrapWithCode(err, msg, ErrUnknown)
}

func WrapWithCode(err error, msg string, code ErrCode) *Error {
	t := getTrace()

	resp := Error{
		message:      msg,
		code:         code,
		fileName:     t.FileName,
		functionName: t.FunctionName,
		line:         t.Line,
		stacks:       make([]*Error, 0),
	}

	if err == nil {
		return &resp
	}

	if ce, ok := err.(*Error); ok {
		if ce == nil {
			rootCause := &Error{
				message:      msg,
				code:         code,
				fileName:     t.FileName,
				functionName: t.FunctionName,
				line:         t.Line,
				stacks:       make([]*Error, 0),
			}

			resp.stacks = append(resp.stacks, rootCause)

			return &resp
		}

		resp.stacks = append([]*Error{&resp}, ce.stacks...)
	} else {
		resp.stacks = append(resp.stacks, &resp, NewWithCode(err.Error(), code))
	}

	return &resp
}

func Errorf(format string, args ...interface{}) error {
	return ErrorfWithCode(ErrUnexpected, format, args...)
}

func ErrorfWithCode(code ErrCode, format string, args ...interface{}) error {
	return &Error{
		message: fmt.Sprintf(format, args...),
		code:    code,
	}
}

func Is(err, target error) bool {
	return false
}

func As(err error, target interface{}) bool {
	return false
}

func Unwrap(err error) error {
	return nil
}

func IsNotFound(err error) bool {
	return false
}

func RootCause(err error) *Error {
	if t, ok := err.(*Error); ok {
		ls := len(t.stacks)

		if ls == 0 {
			return t
		}

		rc := t.stacks[ls-1]

		return rc
	}

	return &Error{
		message: err.Error(),
		code:    ErrUnexpected,
	}
}

func Client(err error) *Error {
	if t, ok := err.(*Error); ok {
		return t
	}

	return &Error{
		message: err.Error(),
		code:    ErrUnexpected,
	}
}
