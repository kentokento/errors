package errors

import (
	"testing"
)

func Test_appError_IsCrit(t *testing.T) {
	tests := []struct {
		name string
		err  AppError
		want bool
	}{
		{"400 error", InvalidParameter.New(), false},
		{"500 error", SystemDefault.New(), true},
		{"400 error Crit", InvalidParameter.New().Crit(), true},
		{"400 error WrapCrit", Wrap(Wrap(InvalidParameter.New().Crit())), true},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := tt.err.IsCrit(); got != tt.want {
				t.Errorf("appError.IsCrit() = %v, want %v", got, tt.want)
			}
		})
	}
}
