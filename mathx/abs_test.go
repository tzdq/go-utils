package mathx

import (
	"math"
	"testing"
)

func TestAbsFloat32(t *testing.T) {
	type args struct {
		n float32
	}
	tests := []struct {
		name string
		args args
		want float32
	}{
		{"positiveNumber", args{6.3}, 6.3},
		{"negativeNumber", args{-6.9}, 6.9},
		{"Zero", args{0}, 0},
		{"positiveNumberMax", args{math.MaxFloat32}, math.MaxFloat32},
		{"negativeNumberMin", args{math.SmallestNonzeroFloat32}, math.SmallestNonzeroFloat32},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := AbsFloat32(tt.args.n); got != tt.want {
				t.Errorf("AbsFloat32() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestAbsInt8(t *testing.T) {
	type args struct {
		n int8
	}
	tests := []struct {
		name string
		args args
		want int8
	}{
		{"positiveNumber", args{6}, 6},
		{"negativeNumber", args{-6}, 6},
		{"Zero", args{0}, 0},
		{"positiveNumberMax", args{math.MaxInt8}, math.MaxInt8},
		{"negativeNumberMin", args{math.MinInt8}, -128},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := AbsInt8(tt.args.n); got != tt.want {
				t.Errorf("AbsInt8() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestAbsInt16(t *testing.T) {
	type args struct {
		n int16
	}
	tests := []struct {
		name string
		args args
		want int16
	}{
		{"positiveNumber", args{6}, 6},
		{"negativeNumber", args{-6}, 6},
		{"Zero", args{0}, 0},
		{"positiveNumberMax", args{math.MaxInt16}, math.MaxInt16},
		{"negativeNumberMin", args{math.MinInt16}, -32768},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := AbsInt16(tt.args.n); got != tt.want {
				t.Errorf("AbsInt16() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestAbsInt32(t *testing.T) {
	type args struct {
		n int32
	}
	tests := []struct {
		name string
		args args
		want int32
	}{
		{"positiveNumber", args{6}, 6},
		{"negativeNumber", args{-6}, 6},
		{"Zero", args{0}, 0},
		{"positiveNumberMax", args{math.MaxInt32}, math.MaxInt32},
		{"negativeNumberMin", args{math.MinInt32}, -2147483648},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := AbsInt32(tt.args.n); got != tt.want {
				t.Errorf("AbsInt32() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestAbsInt64(t *testing.T) {
	type args struct {
		n int64
	}
	tests := []struct {
		name string
		args args
		want int64
	}{
		{"positiveNumber", args{6}, 6},
		{"negativeNumber", args{-6}, 6},
		{"Zero", args{0}, 0},
		{"positiveNumberMax", args{math.MaxInt64}, math.MaxInt64},
		{"negativeNumberMin", args{math.MinInt64}, -9223372036854775808},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := AbsInt64(tt.args.n); got != tt.want {
				t.Errorf("AbsInt64() = %v, want %v", got, tt.want)
			}
		})
	}
}
