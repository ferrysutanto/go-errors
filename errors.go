package errors

import (
	"fmt"
	"strings"
)

type Error struct {
	message string
	code    ErrCode

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
		resp[k] = err.message
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
	err := Error{
		message: msg,
		code:    code,
	}

	return &err
}

type ErrProps struct {
	ClientError     *Error
	ClientErrorCode *ErrCode
	RootCause       *Error
	RootCauseCode   *ErrCode
}

func Wrap(err error, msg string) *Error {
	return WrapWithCode(err, msg, ErrUnexpected)
}

func WrapWithCode(err error, msg string, code ErrCode) *Error {
	resp := Error{
		message: msg,
		code:    code,
		stacks:  make([]*Error, 0),
	}

	if err == nil {
		return &resp
	}

	if t, ok := err.(*Error); ok {
		if t == nil {
			rootCause := &Error{
				message: msg,
				code:    code,
				stacks:  make([]*Error, 0),
			}

			resp.stacks = append(resp.stacks, rootCause)

			return &resp
		}

		resp.stacks = append([]*Error{&resp}, t.stacks...)
	} else {
		resp.stacks = append(resp.stacks, New(err.Error()))
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
