package convert

import (
	"errors"
	"reflect"
	"testing"
)

func TestToIntSlice(t *testing.T) {
	s := []int{1, 2}
	ds := [][]int{s}
	type args struct {
		i interface{}
	}
	tests := []struct {
		name    string
		args    args
		want    []int
		wantErr bool
	}{
		{name: "nil", args: args{nil}, want: nil, wantErr: true},
		{name: "[]interface{}", args: args{[]interface{}{1, 2, 3, 4}}, want: []int{1, 2, 3, 4}},
		{name: "[]int", args: args{[]int{1, 2, 3, 4}}, want: []int{1, 2, 3, 4}},
		{name: "[]int8", args: args{[]int8{1, 2, 3, 4}}, want: []int{1, 2, 3, 4}},
		{name: "[]int16", args: args{[]int16{1, 2, 3, 4}}, want: []int{1, 2, 3, 4}},
		{name: "[]int32", args: args{[]int32{1, 2, 3}}, want: []int{1, 2, 3}},
		{name: "[]int64", args: args{[]int64{1, 2, 3}}, want: []int{1, 2, 3}},
		{name: "[]uint", args: args{[]uint{1, 2, 3, 4}}, want: []int{1, 2, 3, 4}},
		{name: "[]uint8", args: args{[]uint8{1, 2, 3}}, want: []int{1, 2, 3}},
		{name: "[]uint16", args: args{[]uint16{1, 2, 3}}, want: []int{1, 2, 3}},
		{name: "[]uint32", args: args{[]uint32{1, 2, 3}}, want: []int{1, 2, 3}},
		{name: "[]uint64", args: args{[]uint64{1, 2, 3}}, want: []int{1, 2, 3}},
		{name: "[]float32", args: args{[]float32{1.4, 2.222}}, want: []int{1, 2}},
		{name: "[]float64", args: args{[]float64{1.2, 2.2222}}, want: []int{1, 2}},
		{name: "[]bool", args: args{[]bool{true, false, false, true}}, want: []int{1, 0, 0, 1}},
		{name: "[]string", args: args{[]string{"11", "22", "33"}}, want: []int{11, 22, 33}},
		{name: "[]byte", args: args{[]byte("he")}, want: []int{104, 101}},
		{name: "[]rune", args: args{[]rune("我爱你")}, want: []int{25105, 29233, 20320}},
		{name: "string", args: args{"hello"}, want: nil, wantErr: true},
		{name: "ptr", args: args{&s}, want: []int{1, 2}},
		{name: "slice ptr", args: args{&ds}, want: nil, wantErr: true},
		{name: "[]string empty", args: args{""}, want: nil, wantErr: true},
		{name: "[...]string empty", args: args{[...]string{}}, want: []int{}},
		{name: "[...]string", args: args{[...]string{"11"}}, want: []int{11}},
		{name: "int", args: args{11}, want: []int{11}},
		{name: "unknown", args: args{errors.New("error")}, want: nil, wantErr: true},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := ToIntSlice(tt.args.i)
			if (err != nil) != tt.wantErr {
				t.Errorf("ToIntSlice() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("ToIntSlice() got = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestToSlice(t *testing.T) {
	s := []int{1, 2}
	ds := [][]int{s}
	type args struct {
		i interface{}
	}
	tests := []struct {
		name    string
		args    args
		want    []interface{}
		wantErr bool
	}{
		{name: "nil", args: args{nil}, want: nil, wantErr: true},
		{name: "[]interface{}", args: args{[]interface{}{1, 2, 3, 4}}, want: []interface{}{1, 2, 3, 4}},
		{name: "[]int", args: args{[]int{1, 2, 3, 4}}, want: []interface{}{1, 2, 3, 4}},
		{name: "[]int8", args: args{[]int8{1, 2, 3, 4}}, want: []interface{}{int8(1), int8(2), int8(3), int8(4)}},
		{name: "[]int16", args: args{[]int16{1, 2, 3, 4}}, want: []interface{}{int16(1), int16(2), int16(3), int16(4)}},
		{name: "[]int32", args: args{[]int32{1, 2, 3}}, want: []interface{}{int32(1), int32(2), int32(3)}},
		{name: "[]int64", args: args{[]int64{1, 2, 3}}, want: []interface{}{int64(1), int64(2), int64(3)}},
		{name: "[]uint", args: args{[]uint{1, 2, 3, 4}}, want: []interface{}{uint(1), uint(2), uint(3), uint(4)}},
		{name: "[]uint8", args: args{[]uint8{1, 2, 3}}, want: []interface{}{uint8(1), uint8(2), uint8(3)}},
		{name: "[]uint16", args: args{[]uint16{1, 2, 3}}, want: []interface{}{uint16(1), uint16(2), uint16(3)}},
		{name: "[]uint32", args: args{[]uint32{1, 2, 3}}, want: []interface{}{uint32(1), uint32(2), uint32(3)}},
		{name: "[]uint64", args: args{[]uint64{1, 2, 3}}, want: []interface{}{uint64(1), uint64(2), uint64(3)}},
		{name: "[]float32", args: args{[]float32{1.4, 2.222}}, want: []interface{}{float32(1.4), float32(2.222)}},
		{name: "[]float64", args: args{[]float64{1.2, 2.2222}}, want: []interface{}{1.2, 2.2222}},
		{name: "[]bool", args: args{[]bool{true, false, false, true}}, want: []interface{}{true, false, false, true}},
		{name: "[]string", args: args{[]string{"11", "22", "33"}}, want: []interface{}{"11", "22", "33"}},
		{name: "[]byte", args: args{[]byte("he")}, want: []interface{}{byte('h'), byte('e')}},
		{name: "[]rune", args: args{[]rune("我爱你")}, want: []interface{}{'我', '爱', '你'}},
		{name: "string", args: args{"hello"}, want: []interface{}{'h', 'e', 'l', 'l', 'o'}},
		{name: "ptr", args: args{&s}, want: []interface{}{1, 2}},
		{name: "slice ptr", args: args{&ds}, want: []interface{}{[]int{1, 2}}},
		{name: "[]string empty", args: args{""}, want: []interface{}{}},
		{name: "[...]string empty", args: args{[...]string{}}, want: []interface{}{}},
		{name: "[...]string", args: args{[...]string{"arr"}}, want: []interface{}{"arr"}},
		{name: "unknown", args: args{errors.New("error")}, want: nil, wantErr: true},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := ToSlice(tt.args.i)
			if (err != nil) != tt.wantErr {
				t.Errorf("ToSlice() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("ToSlice() got = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestToStringSlice(t *testing.T) {
	type args struct {
		i interface{}
	}
	tests := []struct {
		name    string
		args    args
		want    []string
		wantErr bool
	}{
		{name: "[]string", args: args{[]string{"1", "2"}}, want: []string{"1", "2"}},
		{name: "string", args: args{"1"}, want: []string{"1"}},
		{name: "[]error", args: args{[]error{errors.New("err1"), errors.New("err2")}}, want: []string{"err1", "err2"}},
		{name: "int8", args: args{int8(8)}, want: []string{"8"}},
		{name: "[]interface{}", args: args{[]interface{}{1, 2}}, want: []string{"1", "2"}},
		{name: "[]int", args: args{[]int{1, 2}}, want: []string{"1", "2"}},
		{name: "[]int32", args: args{[]int32{1, 2}}, want: []string{"1", "2"}},
		{name: "[]int64", args: args{[]int64{1, 2}}, want: []string{"1", "2"}},
		{name: "[]float64", args: args{[]float64{1, 2.222}}, want: []string{"1", "2.222"}},
		{name: "[]uint64", args: args{[]uint64{1, 2}}, want: []string{"1", "2"}},
		{name: "nil", args: args{nil}, want: nil, wantErr: true},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := ToStringSlice(tt.args.i)
			if (err != nil) != tt.wantErr {
				t.Errorf("ToStringSlice() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("ToStringSlice() got = %v, want %v", got, tt.want)
			}
		})
	}
}
