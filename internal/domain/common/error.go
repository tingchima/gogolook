// Package common provides
package common

import (
	"errors"
	"fmt"
	"net/http"
	"strings"
)

// Error .
type Error struct {
	errCode   ErrCode
	causeErr  error
	clientMsg string
	details   []any
}

// ErrOption .
type ErrOption func(*Error)

func WithMsg(msg string) ErrOption {
	return func(e *Error) {
		e.clientMsg = msg
	}
}

func WithDetail(detail []any) ErrOption {
	return func(e *Error) {
		e.details = detail
	}
}

// NewError .
func NewError(errCode ErrCode, causeErr error, options ...ErrOption) *Error {
	err := Error{
		errCode:  errCode,
		causeErr: causeErr,
	}

	for _, o := range options {
		o(&err)
	}

	return &err
}

// AsErr .
func AsErr(err error, target any) bool {
	return errors.As(err, target)
}

// IsErrCode .
func IsErrCode(targetErr error, targetErrCode ErrCode) bool {
	var domainErr *Error
	ok := AsErr(targetErr, &domainErr)
	if ok {
		return domainErr.errCode == targetErrCode
	}
	return false
}

// Error .
func (e *Error) Error() string {

	var b strings.Builder
	_, _ = b.WriteRune('[')
	_, _ = b.WriteString(fmt.Sprintf("%d", e.HTTPStatus()))
	_, _ = b.WriteRune(']')

	_, _ = b.WriteRune(' ')
	_, _ = b.WriteString(fmt.Sprintf("client msg: %s", e.ClientMsg()))

	_, _ = b.WriteRune(' ')
	if e.causeErr != nil {
		_, _ = b.WriteString(fmt.Sprintf("cause err: %s", e.CauseMsg()))
	}

	return b.String()
}

// CauseErr .
func (e *Error) CauseErr() error {
	if e.causeErr == nil {
		e.causeErr = errors.New("cause error should be implemented")
	}
	return e.causeErr
}

// Name .
func (e *Error) Name() string {
	return e.errCode.Name
}

// ClientMsg .
func (e *Error) ClientMsg() string {
	if e.clientMsg == "" {
		return e.causeErr.Error()
	}
	return e.clientMsg
}

// CauseMsg .
func (e *Error) CauseMsg() string {
	return e.causeErr.Error()
}

// DetailMsg .
func (e *Error) DetailMsg() []any {
	return e.details
}

// HTTPStatus .
func (e *Error) HTTPStatus() int {
	if e.errCode.StatusCode == 0 {
		return http.StatusInternalServerError
	}
	return e.errCode.StatusCode
}

// WithDetails .
func (e *Error) WithDetails(errs ...any) {
	l := len(e.details)
	e.details = append(e.details[:l:l], errs...)
}
