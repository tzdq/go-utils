package convert

import (
	"errors"
	"math"
	"testing"
)

type TStringer struct {
	name string
}

func (ts TStringer) String() string {
	return ts.name
}

func TestToString(t *testing.T) {
	i := uint64(64)
	type args struct {
		i interface{}
	}
	tests := []struct {
		name    string
		args    args
		want    string
		wantErr bool
	}{
		{name: "string", args: args{"123"}, want: "123"},
		{name: "bool", args: args{true}, want: "true"},
		{name: "float32", args: args{float32(1.2)}, want: "1.2"},
		{name: "float64", args: args{2.5}, want: "2.5"},
		{name: "int", args: args{1}, want: "1"},
		{name: "int8", args: args{int8(10)}, want: "10"},
		{name: "int16", args: args{int16(32)}, want: "32"},
		{name: "int32", args: args{math.MinInt32}, want: "-2147483648"},
		{name: "int64", args: args{math.MaxInt64}, want: "9223372036854775807"},
		{name: "uint", args: args{uint(12)}, want: "12"},
		{name: "uint8", args: args{uint8(255)}, want: "255"},
		{name: "uint16", args: args{uint16(12222)}, want: "12222"},
		{name: "uint32", args: args{uint32(12222)}, want: "12222"},
		{name: "uint64", args: args{uint64(12222)}, want: "12222"},
		{name: "bytes", args: args{[]byte{0x32, 0x33, 0x34}}, want: "234"},
		{name: "rune", args: args{[]rune{0x32, 0x33, 0x34}}, want: "234"},
		{name: "nil", args: args{nil}, want: ""},
		{name: "error", args: args{errors.New("123")}, want: "123"},
		{name: "fmt.Stringer", args: args{TStringer{"122223"}}, want: "122223"},
		{name: "complex(a)", args: args{complex(1, 0)}, want: "1"},
		{name: "complex(bi)", args: args{complex(0, 2)}, want: "2i"},
		{name: "complex(a+bi)", args: args{complex(1, 1)}, want: "1+1i"},
		{name: "complex(0)", args: args{complex(0, 0)}, want: "0"},
		{name: "complex(1e-17)", args: args{complex(1e-17, 0)}, want: "0"},
		{name: "map", args: args{map[string]string{"k": "v"}}, want: "", wantErr: true},
		{name: "uint64", args: args{&i}, want: "64"},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := ToString(tt.args.i)
			if (err != nil) != tt.wantErr {
				t.Errorf("ToString() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if got != tt.want {
				t.Errorf("ToString() got = %v, want %v", got, tt.want)
			}
		})
	}
}
