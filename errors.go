package errors

import (
	"fmt"

	"golang.org/x/xerrors"
)

var Nothing bool

func create(msg string) *appError {
	var e appError
	e.message = msg
	e.frame = xerrors.Caller(2)
	return &e
}

func New(msg string) AppError {
	return create(msg)
}

func Errorf(format string, args ...interface{}) AppError {
	return create(fmt.Sprintf(format, args...))
}

func Wrap(err error, msg ...string) AppError {
	if err == nil {
		return nil
	}

	var m string
	if len(msg) != 0 {
		m = msg[0]
	}
	e := create(m)
	e.next = err
	return e
}

func Wrapf(err error, format string, args ...interface{}) AppError {
	e := create(fmt.Sprintf(format, args...))
	e.next = err
	return e
}

func As(err error, target interface{}) bool {
	return xerrors.As(err, target)
}

func AsAppError(err error) *appError {
	if err == nil {
		return nil
	}

	var e *appError
	if xerrors.As(err, &e) {
		return e
	}
	return nil
}

type appError struct {
	// 標準のエラー仕様を満たす変数
	next    error
	message string
	frame   xerrors.Frame

	// 独自の拡張仕様の変数
	data  []map[string]interface{}
	level level

	// APIエラー
	code        string
	infoMessage string
	status      int
}

func (e *appError) Error() string {
	// 一番下位層のメッセージを取り出す
	next := AsAppError(e.next)
	if next != nil {
		return next.Error()
	}
	if e.next == nil {
		if e.message != `` {
			return e.message
		}
		return `no message`
	}
	return e.next.Error()
}

func (e *appError) Is(err error) bool {
	if er := AsAppError(err); er != nil {
		return e.Code() == er.Code()
	}
	return false
}

func (e *appError) Unwrap() error              { return e.next }
func (e *appError) Format(s fmt.State, v rune) { xerrors.FormatError(e, s, v) }

func (e *appError) FormatError(p xerrors.Printer) error {
	var message string
	if e.level != "" {
		message += fmt.Sprintf("[%s] ", e.level)
	}
	if e.code != "" {
		message += fmt.Sprintf("[%s] ", e.code)
	}
	if e.message != "" {
		message += fmt.Sprintf("%s", e.message)
	}
	if len(e.data) != 0 {
		if message != "" {
			message += "\n"
		}
		message += fmt.Sprintf("data: %+v", e.data)
	}

	p.Print(message)
	e.frame.Format(p)
	return e.next
}

func (e *appError) AddData(field string, data interface{}) AppError {
	if e.data == nil {
		e.data = make([]map[string]interface{}, 0)
	}
	e.data = append(e.data, map[string]interface{}{field: data})
	return e
}
