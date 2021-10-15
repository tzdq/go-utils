package convert

import (
	"math"
	"testing"
)

func TestToBool(t *testing.T) {
	s := "true"
	type args struct {
		i interface{}
	}
	tests := []struct {
		name    string
		args    args
		want    bool
		wantErr bool
	}{
		{name: "bool", args: args{false}, want: false},
		{name: "nil", args: args{nil}, want: false},
		{name: "int", args: args{1}, want: true},
		{name: "int8", args: args{int8(1)}, want: true},
		{name: "int16", args: args{int16(1)}, want: true},
		{name: "int32", args: args{int32(1)}, want: true},
		{name: "int64", args: args{int64(2)}, want: true},
		{name: "uint", args: args{uint(1)}, want: true},
		{name: "uint8", args: args{uint8(1)}, want: true},
		{name: "uint16", args: args{uint16(1)}, want: true},
		{name: "uint32", args: args{uint32(1)}, want: true},
		{name: "uint64", args: args{uint64(10)}, want: true},
		{name: "string", args: args{"true"}, want: true},
		{name: "bytes", args: args{[]byte{0x32}}, want: false, wantErr: true},
		{name: "bytes", args: args{[]byte{'f', 'a', 'l', 's', 'e'}}, want: false},
		{name: "ptr_string", args: args{&s}, want: true},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := ToBool(tt.args.i)
			if (err != nil) != tt.wantErr {
				t.Errorf("ToBool() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if got != tt.want {
				t.Errorf("ToBool() got = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestToFloat64(t *testing.T) {
	s := "15.263"
	type args struct {
		i interface{}
	}
	tests := []struct {
		name    string
		args    args
		want    float64
		wantErr bool
	}{
		{name: "float64", args: args{1.232222}, want: 1.232222},
		//{name: "float32", args: args{float32(12.4332)}, want: 12.433199882507324},
		{name: "float32", args: args{float32(12.4332)}, want: 12.4332},
		{name: "int", args: args{1}, want: 1},
		{name: "int8", args: args{int8(1)}, want: 1},
		{name: "int16", args: args{int16(1)}, want: 1},
		{name: "int32", args: args{int32(1)}, want: 1},
		{name: "int64", args: args{int64(2)}, want: 2},
		{name: "uint", args: args{uint(1)}, want: 1},
		{name: "uint8", args: args{uint8(1)}, want: 1},
		{name: "uint16", args: args{uint16(1)}, want: 1},
		{name: "uint32", args: args{uint32(1)}, want: 1},
		{name: "uint64", args: args{uint64(10)}, want: 10},
		{name: "string", args: args{"-10.22"}, want: -10.22},
		{name: "bool", args: args{true}, want: 1},
		{name: "ptr_string", args: args{&s}, want: 15.263},
		{name: "bytes", args: args{[]byte{'1', '5', '.', '2', '6', '3'}}, want: 15.263},
		{name: "unknown", args: args{[]int32{3}}, want: 0, wantErr: true},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := ToFloat64(tt.args.i)
			if (err != nil) != tt.wantErr {
				t.Errorf("ToFloat64() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if got != tt.want {
				t.Errorf("ToFloat64() got = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestToFloat32(t *testing.T) {
	s := "15.263"
	type args struct {
		i interface{}
	}
	tests := []struct {
		name    string
		args    args
		want    float32
		wantErr bool
	}{
		//{name: "float64", args: args{math.MaxFloat32 * 10}, want: math.MaxFloat32 * 10},
		{name: "float64", args: args{math.MaxFloat32 / 10}, want: math.MaxFloat32 / 10},
		{name: "float32", args: args{float32(12.4332)}, want: 12.4332},
		{name: "int", args: args{1}, want: 1},
		{name: "int8", args: args{int8(1)}, want: 1},
		{name: "int16", args: args{int16(1)}, want: 1},
		{name: "int32", args: args{int32(1)}, want: 1},
		{name: "int64", args: args{int64(2)}, want: 2},
		{name: "uint", args: args{uint(1)}, want: 1},
		{name: "uint8", args: args{uint8(1)}, want: 1},
		{name: "uint16", args: args{uint16(1)}, want: 1},
		{name: "uint32", args: args{uint32(1)}, want: 1},
		{name: "uint64", args: args{uint64(10)}, want: 10},
		{name: "string", args: args{"-10.22"}, want: -10.22},
		{name: "bool", args: args{true}, want: 1},
		{name: "ptr_string", args: args{&s}, want: 15.263},
		{name: "bytes", args: args{[]byte{'1', '2'}}, want: 12},
		{name: "unknown", args: args{[]int32{2}}, want: 0, wantErr: true},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := ToFloat32(tt.args.i)
			if (err != nil) != tt.wantErr {
				t.Errorf("ToFloat32() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if got != tt.want {
				t.Errorf("ToFloat32() got = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestToInt(t *testing.T) {
	s := "15"
	type args struct {
		i interface{}
	}
	tests := []struct {
		name    string
		args    args
		want    int
		wantErr bool
	}{
		{name: "float64", args: args{1.222}, want: 1},
		{name: "float32", args: args{float32(12.4332)}, want: 12},
		{name: "int", args: args{1}, want: 1},
		{name: "int8", args: args{int8(1)}, want: 1},
		{name: "int16", args: args{int16(1)}, want: 1},
		{name: "int32", args: args{int32(1)}, want: 1},
		{name: "int64", args: args{int64(2)}, want: 2},
		{name: "uint", args: args{uint(1)}, want: 1},
		{name: "uint8", args: args{uint8(1)}, want: 1},
		{name: "uint16", args: args{uint16(1)}, want: 1},
		{name: "uint32", args: args{uint32(1)}, want: 1},
		{name: "uint64", args: args{uint64(10)}, want: 10},
		{name: "string", args: args{"-10"}, want: -10},
		{name: "bool", args: args{true}, want: 1},
		{name: "ptr_string", args: args{&s}, want: 15},
		{name: "nil", args: args{nil}, want: 0},
		{name: "bytes", args: args{[]byte{0x32}}, want: 2},
		{name: "unknown", args: args{[]int{2}}, want: 0, wantErr: true},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := ToInt(tt.args.i)
			if (err != nil) != tt.wantErr {
				t.Errorf("ToInt() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if got != tt.want {
				t.Errorf("ToInt() got = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestToInt8(t *testing.T) {
	s := "15"
	type args struct {
		i interface{}
	}
	tests := []struct {
		name    string
		args    args
		want    int8
		wantErr bool
	}{
		{name: "float64", args: args{1.222}, want: 1},
		{name: "float32", args: args{float32(12.4332)}, want: 12},
		{name: "int", args: args{1}, want: 1},
		{name: "int8", args: args{int8(1)}, want: 1},
		{name: "int16", args: args{int16(1)}, want: 1},
		{name: "int32", args: args{int32(1)}, want: 1},
		{name: "int64", args: args{int64(2)}, want: 2},
		{name: "uint", args: args{uint(1)}, want: 1},
		{name: "uint8", args: args{uint8(1)}, want: 1},
		{name: "uint16", args: args{uint16(1)}, want: 1},
		{name: "uint32", args: args{uint32(1)}, want: 1},
		{name: "uint64", args: args{uint64(10)}, want: 10},
		{name: "string", args: args{"-10"}, want: -10},
		{name: "bool", args: args{true}, want: 1},
		{name: "ptr_string", args: args{&s}, want: 15},
		{name: "nil", args: args{nil}, want: 0},
		{name: "bytes", args: args{[]byte{0x32}}, want: 2},
		{name: "unknown", args: args{[]int{0x32}}, want: 0, wantErr: true},
		{name: "out range", args: args{int32(1222)}, want: -58},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := ToInt8(tt.args.i)
			if (err != nil) != tt.wantErr {
				t.Errorf("ToInt8() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if got != tt.want {
				t.Errorf("ToInt8() got = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestToInt16(t *testing.T) {
	s := "15"
	type args struct {
		i interface{}
	}
	tests := []struct {
		name    string
		args    args
		want    int16
		wantErr bool
	}{
		{name: "float64", args: args{1.222}, want: 1},
		{name: "float32", args: args{float32(12.4332)}, want: 12},
		{name: "int", args: args{1}, want: 1},
		{name: "int8", args: args{int8(1)}, want: 1},
		{name: "int16", args: args{int16(1)}, want: 1},
		{name: "int32", args: args{int32(1)}, want: 1},
		{name: "int64", args: args{int64(2)}, want: 2},
		{name: "uint", args: args{uint(1)}, want: 1},
		{name: "uint8", args: args{uint8(1)}, want: 1},
		{name: "uint16", args: args{uint16(1)}, want: 1},
		{name: "uint32", args: args{uint32(1)}, want: 1},
		{name: "uint64", args: args{uint64(10)}, want: 10},
		{name: "string", args: args{"-10"}, want: -10},
		{name: "bool", args: args{true}, want: 1},
		{name: "ptr_string", args: args{&s}, want: 15},
		{name: "nil", args: args{nil}, want: 0},
		{name: "bytes", args: args{[]byte{0x32}}, want: 2},
		{name: "unknown", args: args{[]int{0x32}}, want: 0, wantErr: true},
		{name: "out range", args: args{int32(12222222)}, want: 32526},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := ToInt16(tt.args.i)
			if (err != nil) != tt.wantErr {
				t.Errorf("ToInt16() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if got != tt.want {
				t.Errorf("ToInt16() got = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestToInt32(t *testing.T) {
	s := "15"
	type args struct {
		i interface{}
	}
	tests := []struct {
		name    string
		args    args
		want    int32
		wantErr bool
	}{
		{name: "float64", args: args{1.222}, want: 1},
		{name: "float32", args: args{float32(12.4332)}, want: 12},
		{name: "int", args: args{1}, want: 1},
		{name: "int8", args: args{int8(1)}, want: 1},
		{name: "int16", args: args{int16(1)}, want: 1},
		{name: "int32", args: args{int32(1)}, want: 1},
		{name: "int64", args: args{int64(2)}, want: 2},
		{name: "uint", args: args{uint(1)}, want: 1},
		{name: "uint8", args: args{uint8(1)}, want: 1},
		{name: "uint16", args: args{uint16(1)}, want: 1},
		{name: "uint32", args: args{uint32(1)}, want: 1},
		{name: "uint64", args: args{uint64(10)}, want: 10},
		{name: "string", args: args{"-10"}, want: -10},
		{name: "bool", args: args{true}, want: 1},
		{name: "ptr_string", args: args{&s}, want: 15},
		{name: "nil", args: args{nil}, want: 0},
		{name: "bytes", args: args{[]byte{0x32}}, want: 2},
		{name: "unknown", args: args{[]int{0x32}}, want: 0, wantErr: true},
		{name: "out range", args: args{int64(12222222222)}, want: -662679666},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := ToInt32(tt.args.i)
			if (err != nil) != tt.wantErr {
				t.Errorf("ToInt32() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if got != tt.want {
				t.Errorf("ToInt32() got = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestToInt64(t *testing.T) {
	s := "15"
	type args struct {
		i interface{}
	}
	tests := []struct {
		name    string
		args    args
		want    int64
		wantErr bool
	}{
		{name: "float64", args: args{float64(1)}, want: 1},
		{name: "float32", args: args{float32(12)}, want: 12},
		{name: "int", args: args{1}, want: 1},
		{name: "int8", args: args{int8(1)}, want: 1},
		{name: "int16", args: args{int16(1)}, want: 1},
		{name: "int32", args: args{int32(1)}, want: 1},
		{name: "int64", args: args{int64(2)}, want: 2},
		{name: "uint", args: args{uint(1)}, want: 1},
		{name: "uint8", args: args{uint8(1)}, want: 1},
		{name: "uint16", args: args{uint16(1)}, want: 1},
		{name: "uint32", args: args{uint32(math.MaxUint32)}, want: math.MaxUint32},
		{name: "uint64", args: args{uint64(math.MaxUint64)}, want: -1},
		{name: "string", args: args{"-93109800000000000"}, want: -93109800000000000},
		{name: "bool", args: args{true}, want: 1},
		{name: "ptr_string", args: args{&s}, want: 15},
		{name: "nil", args: args{nil}, want: 0},
		{name: "bytes", args: args{[]byte{0x32}}, want: 2},
		{name: "unknown", args: args{[]int{2}}, want: 0, wantErr: true},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := ToInt64(tt.args.i)
			if (err != nil) != tt.wantErr {
				t.Errorf("ToInt64() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if got != tt.want {
				t.Errorf("ToInt64() got = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestToUint(t *testing.T) {
	s := "15"
	type args struct {
		i interface{}
	}
	tests := []struct {
		name    string
		args    args
		want    uint
		wantErr bool
	}{
		{name: "float64", args: args{1.222}, want: 1},
		{name: "float32", args: args{float32(12.4332)}, want: 12},
		{name: "int", args: args{1}, want: 1},
		{name: "int8", args: args{int8(1)}, want: 1},
		{name: "int16", args: args{int16(1)}, want: 1},
		{name: "int32", args: args{int32(1)}, want: 1},
		{name: "int64", args: args{int64(2)}, want: 2},
		{name: "float64(-x)", args: args{-1.222}, want: 0, wantErr: true},
		{name: "float32(-x)", args: args{float32(-12.4332)}, want: 0, wantErr: true},
		{name: "int(-x)", args: args{-1}, want: 0, wantErr: true},
		{name: "int8(-x)", args: args{int8(-1)}, want: 0, wantErr: true},
		{name: "int16(-x)", args: args{int16(-1)}, want: 0, wantErr: true},
		{name: "int32(-x)", args: args{int32(-1)}, want: 0, wantErr: true},
		{name: "int64(-x)", args: args{int64(-2)}, want: 0, wantErr: true},
		{name: "uint", args: args{uint(1)}, want: 1},
		{name: "uint8", args: args{uint8(1)}, want: 1},
		{name: "uint16", args: args{uint16(1)}, want: 1},
		{name: "uint32", args: args{uint32(1)}, want: 1},
		{name: "uint64", args: args{uint64(10)}, want: 10},
		{name: "string", args: args{"10"}, want: 10},
		{name: "bool", args: args{true}, want: 1},
		{name: "ptr_string", args: args{&s}, want: 15},
		{name: "nil", args: args{nil}, want: 0},
		{name: "bytes", args: args{[]byte{0x32}}, want: 2},
		{name: "unknown", args: args{[]int32{2}}, want: 0, wantErr: true},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := ToUint(tt.args.i)
			if (err != nil) != tt.wantErr {
				t.Errorf("ToUint() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if got != tt.want {
				t.Errorf("ToUint() got = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestToUint8(t *testing.T) {
	s := "15"
	type args struct {
		i interface{}
	}
	tests := []struct {
		name    string
		args    args
		want    uint8
		wantErr bool
	}{
		{name: "float64", args: args{1.222}, want: 1},
		{name: "float32", args: args{float32(12.4332)}, want: 12},
		{name: "int", args: args{1}, want: 1},
		{name: "int8", args: args{int8(1)}, want: 1},
		{name: "int16", args: args{int16(1)}, want: 1},
		{name: "int32", args: args{int32(1)}, want: 1},
		{name: "int64", args: args{int64(2)}, want: 2},
		{name: "float64(-x)", args: args{-1.222}, want: 0, wantErr: true},
		{name: "float32(-x)", args: args{float32(-12.4332)}, want: 0, wantErr: true},
		{name: "int(-x)", args: args{-1}, want: 0, wantErr: true},
		{name: "int8(-x)", args: args{int8(-1)}, want: 0, wantErr: true},
		{name: "int16(-x)", args: args{int16(-1)}, want: 0, wantErr: true},
		{name: "int32(-x)", args: args{int32(-1)}, want: 0, wantErr: true},
		{name: "int64(-x)", args: args{int64(-2)}, want: 0, wantErr: true},
		{name: "uint", args: args{uint(1)}, want: 1},
		{name: "uint8", args: args{uint8(1)}, want: 1},
		{name: "uint16", args: args{uint16(1)}, want: 1},
		{name: "uint32", args: args{uint32(1)}, want: 1},
		{name: "uint64", args: args{uint64(10)}, want: 10},
		{name: "string", args: args{"10"}, want: 10},
		{name: "bool", args: args{true}, want: 1},
		{name: "ptr_string", args: args{&s}, want: 15},
		{name: "nil", args: args{nil}, want: 0},
		{name: "bytes", args: args{[]byte{0x32}}, want: 2},
		{name: "unknown", args: args{[]int{0x32}}, want: 0, wantErr: true},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := ToUint8(tt.args.i)
			if (err != nil) != tt.wantErr {
				t.Errorf("ToUint8() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if got != tt.want {
				t.Errorf("ToUint8() got = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestToUint16(t *testing.T) {
	s := "15"
	type args struct {
		i interface{}
	}
	tests := []struct {
		name    string
		args    args
		want    uint16
		wantErr bool
	}{
		{name: "float64", args: args{1.222}, want: 1},
		{name: "float32", args: args{float32(12.4332)}, want: 12},
		{name: "int", args: args{1}, want: 1},
		{name: "int8", args: args{int8(1)}, want: 1},
		{name: "int16", args: args{int16(1)}, want: 1},
		{name: "int32", args: args{int32(1)}, want: 1},
		{name: "int64", args: args{int64(2)}, want: 2},
		{name: "float64(-x)", args: args{-1.222}, want: 0, wantErr: true},
		{name: "float32(-x)", args: args{float32(-12.4332)}, want: 0, wantErr: true},
		{name: "int(-x)", args: args{-1}, want: 0, wantErr: true},
		{name: "int8(-x)", args: args{int8(-1)}, want: 0, wantErr: true},
		{name: "int16(-x)", args: args{int16(-1)}, want: 0, wantErr: true},
		{name: "int32(-x)", args: args{int32(-1)}, want: 0, wantErr: true},
		{name: "int64(-x)", args: args{int64(-2)}, want: 0, wantErr: true},
		{name: "uint", args: args{uint(1)}, want: 1},
		{name: "uint8", args: args{uint8(1)}, want: 1},
		{name: "uint16", args: args{uint16(1)}, want: 1},
		{name: "uint32", args: args{uint32(1)}, want: 1},
		{name: "uint64", args: args{uint64(10)}, want: 10},
		{name: "string", args: args{"10"}, want: 10},
		{name: "bool", args: args{true}, want: 1},
		{name: "ptr_string", args: args{&s}, want: 15},
		{name: "nil", args: args{nil}, want: 0},
		{name: "bytes", args: args{[]byte{0x32}}, want: 2},
		{name: "unknown", args: args{[]int32{0x32}}, want: 0, wantErr: true},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := ToUint16(tt.args.i)
			if (err != nil) != tt.wantErr {
				t.Errorf("ToUint16() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if got != tt.want {
				t.Errorf("ToUint16() got = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestToUint32(t *testing.T) {
	s := "15"
	type args struct {
		i interface{}
	}
	tests := []struct {
		name    string
		args    args
		want    uint32
		wantErr bool
	}{
		{name: "float64", args: args{1.222}, want: 1},
		{name: "float32", args: args{float32(12.4332)}, want: 12},
		{name: "int", args: args{1}, want: 1},
		{name: "int8", args: args{int8(1)}, want: 1},
		{name: "int16", args: args{int16(1)}, want: 1},
		{name: "int32", args: args{int32(1)}, want: 1},
		{name: "int64", args: args{int64(2)}, want: 2},
		{name: "float64(-x)", args: args{-1.222}, want: 0, wantErr: true},
		{name: "float32(-x)", args: args{float32(-12.4332)}, want: 0, wantErr: true},
		{name: "int(-x)", args: args{-1}, want: 0, wantErr: true},
		{name: "int8(-x)", args: args{int8(-1)}, want: 0, wantErr: true},
		{name: "int16(-x)", args: args{int16(-1)}, want: 0, wantErr: true},
		{name: "int32(-x)", args: args{int32(-1)}, want: 0, wantErr: true},
		{name: "int64(-x)", args: args{int64(-2)}, want: 0, wantErr: true},
		{name: "uint", args: args{uint(1)}, want: 1},
		{name: "uint8", args: args{uint8(1)}, want: 1},
		{name: "uint16", args: args{uint16(1)}, want: 1},
		{name: "uint32", args: args{uint32(1)}, want: 1},
		{name: "uint64", args: args{uint64(10)}, want: 10},
		{name: "string", args: args{"10"}, want: 10},
		{name: "bool", args: args{true}, want: 1},
		{name: "ptr_string", args: args{&s}, want: 15},
		{name: "nil", args: args{nil}, want: 0},
		{name: "bytes", args: args{[]byte{0x32}}, want: 2},
		{name: "unknown", args: args{[]int32{0x32}}, want: 0, wantErr: true},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := ToUint32(tt.args.i)
			if (err != nil) != tt.wantErr {
				t.Errorf("ToUint32() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if got != tt.want {
				t.Errorf("ToUint32() got = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestToUint64(t *testing.T) {
	s := "15"
	type args struct {
		i interface{}
	}
	tests := []struct {
		name    string
		args    args
		want    uint64
		wantErr bool
	}{
		{name: "float64", args: args{1.222}, want: 1},
		{name: "float32", args: args{float32(12.4332)}, want: 12},
		{name: "int", args: args{1}, want: 1},
		{name: "int8", args: args{int8(1)}, want: 1},
		{name: "int16", args: args{int16(1)}, want: 1},
		{name: "int32", args: args{int32(1)}, want: 1},
		{name: "int64", args: args{int64(2)}, want: 2},
		{name: "float64(-x)", args: args{-1.222}, want: 0, wantErr: true},
		{name: "float32(-x)", args: args{float32(-12.4332)}, want: 0, wantErr: true},
		{name: "int(-x)", args: args{-1}, want: 0, wantErr: true},
		{name: "int8(-x)", args: args{int8(-1)}, want: 0, wantErr: true},
		{name: "int16(-x)", args: args{int16(-1)}, want: 0, wantErr: true},
		{name: "int32(-x)", args: args{int32(-1)}, want: 0, wantErr: true},
		{name: "int64(-x)", args: args{int64(-2)}, want: 0, wantErr: true},
		{name: "uint", args: args{uint(1)}, want: 1},
		{name: "uint8", args: args{uint8(1)}, want: 1},
		{name: "uint16", args: args{uint16(1)}, want: 1},
		{name: "uint32", args: args{uint32(1)}, want: 1},
		{name: "uint64", args: args{uint64(10)}, want: 10},
		{name: "string", args: args{"10"}, want: 10},
		{name: "bool", args: args{true}, want: 1},
		{name: "ptr_string", args: args{&s}, want: 15},
		{name: "nil", args: args{nil}, want: 0},
		{name: "bytes", args: args{[]byte{0x32, 0x35}}, want: 25},
		{name: "slice", args: args{[]int32{0}}, want: 0, wantErr: true},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := ToUint64(tt.args.i)
			if (err != nil) != tt.wantErr {
				t.Errorf("ToUint64() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if got != tt.want {
				t.Errorf("ToUint64() got = %v, want %v", got, tt.want)
			}
		})
	}
}
