package crypt

import (
	"reflect"
	"testing"
)

func TestBase64Decode(t *testing.T) {
	type args struct {
		str string
	}
	tests := []struct {
		name    string
		args    args
		want    []byte
		wantErr bool
	}{
		{
			"#1",
			args{str: "dGhpcyBpcyDllYogdGVzdCBjYXNl"},
			[]byte("this is 啊 test case"),
			false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := Base64Decode(tt.args.str)
			if (err != nil) != tt.wantErr {
				t.Errorf("Base64Decode() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("Base64Decode() got = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestBase64Encode(t *testing.T) {
	type args struct {
		data []byte
	}
	tests := []struct {
		name string
		args args
		want string
	}{
		{
			"#1",
			args{data: []byte("this is 啊 test case")},
			"dGhpcyBpcyDllYogdGVzdCBjYXNl",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := Base64Encode(tt.args.data); got != tt.want {
				t.Errorf("Base64Encode() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestBase64UrlDecode(t *testing.T) {
	type args struct {
		str string
	}
	tests := []struct {
		name    string
		args    args
		want    []byte
		wantErr bool
	}{
		{
			"#1",
			args{"aHR0cHM6Ly90b29sLmNoaW5hei5jb20vdG9vbHMvdXJsY3J5cHQuYXNweA=="},
			[]byte("https://tool.chinaz.com/tools/urlcrypt.aspx"),
			false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := Base64UrlDecode(tt.args.str)
			if (err != nil) != tt.wantErr {
				t.Errorf("Base64UrlDecode() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("Base64UrlDecode() got = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestBase64UrlEncode(t *testing.T) {
	type args struct {
		data []byte
	}
	tests := []struct {
		name string
		args args
		want string
	}{
		{
			"#1",
			args{[]byte("https://tool.chinaz.com/tools/urlcrypt.aspx")},
			"aHR0cHM6Ly90b29sLmNoaW5hei5jb20vdG9vbHMvdXJsY3J5cHQuYXNweA==",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := Base64UrlEncode(tt.args.data); got != tt.want {
				t.Errorf("Base64UrlEncode() = %v, want %v", got, tt.want)
			}
		})
	}
}
