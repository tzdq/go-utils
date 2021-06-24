package crypt

import (
	"reflect"
	"testing"
)

func Test_ISO10126Padding(t *testing.T) {
	type args struct {
		ciphertext []byte
		blockSize  int
	}
	tests := []struct {
		name string
		args args
		want []byte
	}{
		{
			name: "Normal",
			args: args{
				ciphertext: []byte{0x99, 0x98, 0x97, 0x96, 0x95, 0x94, 0x93, 0x92, 0x91, 0x90},
				blockSize:  8,
			},
			// 由于是随机的不能直接对比值，只需要判断范围即可
			want: []byte{0x99, 0x98, 0x97, 0x96, 0x95, 0x94, 0x93, 0x92, 0x91, 0x90, 0x0, 0x0, 0x0, 0x0, 0x0, 0x6},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := ISO10126Padding(tt.args.ciphertext, tt.args.blockSize)
			if !compareISO10126Padding(got, tt.want) {
				t.Errorf("ISO10126Padding() = %v, want %v", got, tt.want)
			}
		})
	}
}

func compareISO10126Padding(src, dst []byte) bool {
	srcLen := len(src)
	dstLen := len(dst)
	if srcLen != dstLen {
		return false
	}

	srcPadding := src[srcLen-1]
	dstPadding := dst[dstLen-1]
	if srcPadding != dstPadding {
		return false
	}
	for i := 0; i < srcLen-int(srcPadding); i++ {
		if src[i] != dst[i] {
			return false
		}
	}
	for i := srcLen - int(srcPadding); i < srcLen; i++ {
		if src[i] > 255 || src[i] < 0 {
			return false
		}
	}
	return true
}

func Test_ISO10126UnPadding(t *testing.T) {
	type args struct {
		origData []byte
	}
	tests := []struct {
		name     string
		args     args
		wantData []byte
		wantErr  bool
	}{
		{
			name:     "DataEmpty",
			args:     args{[]byte{}},
			wantErr:  true,
			wantData: nil,
		},
		{
			name: "LengthInvalid",
			args: args{[]byte{0x99, 0x98, 0x97, 0x96, 0x95, 0x94, 0x93, 0x92, 0x91, 0x90, 0x0, 0x0, 0x0,
				0x0, 0x0, 0x18}},
			wantErr:  true,
			wantData: nil,
		},
		{
			name: "Normal",
			args: args{[]byte{0x99, 0x98, 0x97, 0x96, 0x95, 0x94, 0x93, 0x92, 0x91, 0x90, 0x04, 0x40, 0x1A,
				0x22, 0x34, 0x6}},
			wantErr:  false,
			wantData: []byte{0x99, 0x98, 0x97, 0x96, 0x95, 0x94, 0x93, 0x92, 0x91, 0x90},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			gotData, err := ISO10126UnPadding(tt.args.origData)
			if (err != nil) != tt.wantErr {
				t.Errorf("ISO10126UnPadding() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(gotData, tt.wantData) {
				t.Errorf("ISO10126UnPadding() gotData = %v, want %v", gotData, tt.wantData)
			}
		})
	}
}

func Test_PKCS7Padding(t *testing.T) {
	type args struct {
		data      []byte
		blockSize int
	}
	tests := []struct {
		name string
		args args
		want []byte
	}{
		{
			name: "Normal",
			args: args{
				data:      []byte{0x99, 0x98, 0x97, 0x96, 0x95, 0x94, 0x93, 0x92, 0x91, 0x90},
				blockSize: 8,
			},
			want: []byte{0x99, 0x98, 0x97, 0x96, 0x95, 0x94, 0x93, 0x92, 0x91, 0x90, 0x6, 0x6, 0x6, 0x6, 0x6, 0x6},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := PKCS7Padding(tt.args.data, tt.args.blockSize); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("pkcs5Padding() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_PKCS7UnPadding(t *testing.T) {
	type args struct {
		origData []byte
	}
	tests := []struct {
		name     string
		args     args
		wantData []byte
		wantErr  bool
	}{
		{
			name:     "DataEmpty",
			args:     args{[]byte{}},
			wantErr:  true,
			wantData: nil,
		},
		{
			name: "LengthInvalid",
			args: args{[]byte{0x99, 0x98, 0x97, 0x96, 0x95, 0x94, 0x93, 0x92, 0x91, 0x90, 0x6, 0x6, 0x6,
				0x6, 0x6, 0x18}},
			wantErr:  true,
			wantData: nil,
		},
		{
			name: "LengthInvalid",
			args: args{[]byte{0x99, 0x98, 0x97, 0x96, 0x95, 0x94, 0x93, 0x92, 0x91, 0x90, 0x1, 0x6, 0x6,
				0x6, 0x6, 0x6}},
			wantErr:  true,
			wantData: nil,
		},
		{
			name: "Normal",
			args: args{[]byte{0x99, 0x98, 0x97, 0x96, 0x95, 0x94, 0x93, 0x92, 0x91, 0x90, 0x6, 0x6, 0x6,
				0x6, 0x6, 0x6}},
			wantErr:  false,
			wantData: []byte{0x99, 0x98, 0x97, 0x96, 0x95, 0x94, 0x93, 0x92, 0x91, 0x90},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			gotData, err := PKCS7UnPadding(tt.args.origData)
			if (err != nil) != tt.wantErr {
				t.Errorf("PKCS7UnPadding() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(gotData, tt.wantData) {
				t.Errorf("PKCS7UnPadding() gotData = %v, want %v", gotData, tt.wantData)
			}
		})
	}
}

func Test_X923Padding(t *testing.T) {
	type args struct {
		data      []byte
		blockSize int
	}
	tests := []struct {
		name string
		args args
		want []byte
	}{
		{
			name: "Normal",
			args: args{
				data:      []byte{0x99, 0x98, 0x97, 0x96, 0x95, 0x94, 0x93, 0x92, 0x91, 0x90},
				blockSize: 8,
			},
			want: []byte{0x99, 0x98, 0x97, 0x96, 0x95, 0x94, 0x93, 0x92, 0x91, 0x90, 0x0, 0x0, 0x0, 0x0, 0x0, 0x6},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := X923Padding(tt.args.data, tt.args.blockSize); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("X923Padding() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_X923UnPadding(t *testing.T) {
	type args struct {
		origData []byte
	}
	tests := []struct {
		name     string
		args     args
		wantData []byte
		wantErr  bool
	}{
		{
			name:     "DataEmpty",
			args:     args{[]byte{}},
			wantErr:  true,
			wantData: nil,
		},
		{
			name: "LengthInvalid",
			args: args{[]byte{0x99, 0x98, 0x97, 0x96, 0x95, 0x94, 0x93, 0x92, 0x91, 0x90, 0x0, 0x0, 0x0,
				0x0, 0x0, 0x16}},
			wantErr:  true,
			wantData: nil,
		},
		{
			name: "DataInvalid",
			args: args{[]byte{0x99, 0x98, 0x97, 0x96, 0x95, 0x94, 0x93, 0x92, 0x91, 0x90, 0x1, 0x0, 0x0,
				0x0, 0x0, 0x6}},
			wantErr:  true,
			wantData: nil,
		},
		{
			name: "Normal",
			args: args{[]byte{0x99, 0x98, 0x97, 0x96, 0x95, 0x94, 0x93, 0x92, 0x91, 0x90, 0x0, 0x0, 0x0,
				0x0, 0x0, 0x6}},
			wantErr:  false,
			wantData: []byte{0x99, 0x98, 0x97, 0x96, 0x95, 0x94, 0x93, 0x92, 0x91, 0x90},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			gotData, err := X923UnPadding(tt.args.origData)
			if (err != nil) != tt.wantErr {
				t.Errorf("X923UnPadding() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(gotData, tt.wantData) {
				t.Errorf("X923UnPadding() gotData = %v, want %v", gotData, tt.wantData)
			}
		})
	}
}

func Test_ZeroPadding(t *testing.T) {
	type args struct {
		data      []byte
		blockSize int
	}
	tests := []struct {
		name string
		args args
		want []byte
	}{
		{
			name: "Normal",
			args: args{
				data:      []byte{0x99, 0x98, 0x97, 0x96, 0x95, 0x94, 0x93, 0x92, 0x91, 0x90},
				blockSize: 8,
			},
			want: []byte{0x99, 0x98, 0x97, 0x96, 0x95, 0x94, 0x93, 0x92, 0x91, 0x90, 0x0, 0x0, 0x0, 0x0, 0x0, 0x0},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := ZeroPadding(tt.args.data, tt.args.blockSize); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("ZeroPadding() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_ZeroUnPadding(t *testing.T) {
	type args struct {
		origData []byte
	}
	tests := []struct {
		name     string
		args     args
		wantErr  bool
		wantData []byte
	}{
		{
			name:     "DataEmpty",
			args:     args{[]byte{}},
			wantData: nil,
			wantErr:  true,
		},
		{
			name: "Normal",
			args: args{[]byte{0x99, 0x98, 0x97, 0x96, 0x95, 0x94, 0x93, 0x92, 0x91, 0x90, 0x0, 0x0, 0x0,
				0x0, 0x0, 0x0}},
			wantData: []byte{0x99, 0x98, 0x97, 0x96, 0x95, 0x94, 0x93, 0x92, 0x91, 0x90},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			gotData, err := ZeroUnPadding(tt.args.origData)
			if (err != nil) != tt.wantErr {
				t.Errorf("ZeroUnPadding() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(gotData, tt.wantData) {
				t.Errorf("ZeroUnPadding() gotData = %v, want %v", gotData, tt.wantData)
			}
		})
	}
}

func Test_ISO7816Padding(t *testing.T) {
	type args struct {
		data      []byte
		blockSize int
	}
	tests := []struct {
		name string
		args args
		want []byte
	}{
		{
			name: "Normal",
			args: args{
				data:      []byte{0x99, 0x98, 0x97, 0x96, 0x95, 0x94, 0x93, 0x92, 0x91, 0x90},
				blockSize: 8,
			},
			want: []byte{0x99, 0x98, 0x97, 0x96, 0x95, 0x94, 0x93, 0x92, 0x91, 0x90, 0x80, 0x0, 0x0, 0x0, 0x0, 0x0},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := ISO7816Padding(tt.args.data, tt.args.blockSize); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("ISO7816Padding() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_ISO7816UnPadding(t *testing.T) {
	type args struct {
		origData []byte
	}
	tests := []struct {
		name     string
		args     args
		wantData []byte
		wantErr  bool
	}{
		{
			name:     "DataEmpty",
			args:     args{[]byte{}},
			wantErr:  true,
			wantData: nil,
		},
		{
			name: "LengthInvalid",
			args: args{[]byte{0x99, 0x98, 0x97, 0x96, 0x95, 0x94, 0x93, 0x92, 0x91, 0x90, 0x0, 0x0, 0x0,
				0x0, 0x0, 0x0}},
			wantErr:  true,
			wantData: nil,
		},
		{
			name: "Normal",
			args: args{[]byte{0x99, 0x98, 0x97, 0x96, 0x95, 0x94, 0x93, 0x92, 0x91, 0x90, 0x80, 0x0, 0x0,
				0x0, 0x0, 0x0}},
			wantErr:  false,
			wantData: []byte{0x99, 0x98, 0x97, 0x96, 0x95, 0x94, 0x93, 0x92, 0x91, 0x90},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			gotData, err := ISO7816UnPadding(tt.args.origData)
			if (err != nil) != tt.wantErr {
				t.Errorf("ISO7816UnPadding() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(gotData, tt.wantData) {
				t.Errorf("ISO7816UnPadding() gotData = %v, want %v", gotData, tt.wantData)
			}
		})
	}
}
