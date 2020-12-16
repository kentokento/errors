package errors

import (
	"fmt"
	"net/http"

	"golang.org/x/xerrors"
)

func newError(code string, msg string) *appError {
	e := &appError{
		code:        code,
		infoMessage: msg,
	}
	return e
}

func newBadRequest(code, msg string) *appError {
	e := newError(code, msg)
	e.status = http.StatusBadRequest
	e.Info()
	return e
}

func newUnauthorized(code, msg string) *appError {
	e := newError(code, msg)
	e.status = http.StatusUnauthorized
	e.Info()
	return e
}

func newForbidden(code, msg string) *appError {
	e := newError(code, msg)
	e.status = http.StatusForbidden
	e.Info()
	return e
}

func newConflict(code, msg string) *appError {
	e := newError(code, msg)
	e.status = http.StatusConflict
	e.Info()
	return e
}
func newNotFound(code, msg string) *appError {
	e := newError(code, msg)
	e.status = http.StatusNotFound
	e.Info()
	return e
}

func newInternalServerError(code, msg string) *appError {
	e := newError(code, msg)
	e.status = http.StatusInternalServerError
	e.Crit()
	return e
}

func (e appError) new(a string) *appError {
	e.message = a
	e.frame = xerrors.Caller(2)
	return &e
}

func (e appError) New(msg ...string) AppError {
	var m string
	if len(msg) == 0 {
		m = e.Code()
	} else {
		m = msg[0]
	}
	return e.new(m)
}

func (e appError) Errorf(format string, args ...interface{}) AppError {
	return e.new(fmt.Sprintf(format, args...))
}

func (e appError) Wrap(err error, msg ...string) AppError {
	var m string
	if len(msg) == 0 {
		m = e.Code()
	} else {
		m = msg[0]
	}
	ne := e.new(m)
	ne.next = err
	return ne
}

func (e appError) Wrapf(err error, format string, args ...interface{}) AppError {
	ne := e.new(fmt.Sprintf(format, args...))
	ne.next = err
	return ne
}

// Messagef: ユーザー向けメッセージのフォーマットにパラメータを適用する
func (e appError) Messagef(args ...interface{}) AppError {
	e.infoMessage = fmt.Sprintf(e.infoMessage, args...)
	return &e
}

func (e *appError) Code() string {
	if e.code != `` {
		return e.code
	}
	next := AsAppError(e.next)
	if next != nil {
		return next.Code()
	}
	return `not_defined`
}

func (e *appError) Status() int {
	if e.status != 0 {
		return e.status
	}
	next := AsAppError(e.next)
	if next != nil {
		return next.Status()
	}
	return http.StatusInternalServerError
}

func (e *appError) InfoMessage() string {
	if e.infoMessage != `` {
		return e.infoMessage
	}
	next := AsAppError(e.next)
	if next != nil {
		return next.InfoMessage()
	}
	return "unknown info message"
}
