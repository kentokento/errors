package errors

import (
	"reflect"
	"testing"
)

func TestIs(t *testing.T) {
	err := InvalidParameter.New()
	if !InvalidParameter.Is(err) {
		t.Error("Isの値が一致しません")
	}
	if !InvalidParameter.Is(Wrap(err)) {
		t.Error("Isの値が一致しません")
	}
}

func TestAs(t *testing.T) {
	var err error
	err0 := Wrap(nil)
	err1 := InvalidParameter.New("test")
	err2 := InvalidParameter.New("test").Crit()
	err3 := Wrap(InvalidParameter.New("test"), "wrap")
	err4 := Wrap(Wrap(InvalidParameter.New("test"), "wrap"), "wrap2")
	err5 := InvalidParameter.New("test").AddData("user_id", 123)

	errConv := func(e AppError) error {
		return e
	}
	errConvApp := func(e AppError) *appError {
		if e == nil {
			return nil
		}
		return e.(*appError)
	}

	tests := []struct {
		name string
		err  error
		want *appError
	}{
		{"nil1 ", err, nil},
		{"nil2 ", errConv(nil), nil},
		{"nil2 ", errConvApp(nil), nil},
		{"nil3 ", errConv(errConvApp(nil)), nil},
		{"w nil", errConv(err0), nil},
		{"new  ", errConv(err1), errConvApp(err1)},
		{"level", errConv(err2), errConvApp(err2)},
		{"wrap ", errConv(err3), errConvApp(err3)},
		{"wraps", errConv(err4), errConvApp(err4)},
		{"data ", errConv(err5), errConvApp(err5)},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := AsAppError(tt.err); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("\nerror: %+v\nwant : %+v", got, tt.want)
			}
		})
	}
}
