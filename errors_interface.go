package errors

import (
	"fmt"

	"golang.org/x/xerrors"
)

type AppError interface {
	Error() string
	Unwrap() error
	Format(s fmt.State, v rune)
	FormatError(p xerrors.Printer) error
	AddData(field string, data interface{}) AppError

	BadRequest() AppError
	Unauthorized() AppError
	NotFound() AppError
	InternalServerError() AppError

	Panic() AppError
	Crit() AppError
	Warn() AppError
	Info() AppError
	IsPanic() bool
	IsCrit() bool
	IsWarn() bool
	IsInfo() bool

	New(msg ...string) AppError
	Errorf(format string, args ...interface{}) AppError
	Wrap(err error, msg ...string) AppError
	Wrapf(err error, format string, args ...interface{}) AppError
	Code() string
	Status() int
	InfoMessage() string
	Messagef(args ...interface{}) AppError
	Is(err error) bool
}
