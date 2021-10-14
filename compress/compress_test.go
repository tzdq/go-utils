package compress

import (
	"reflect"
	"testing"
)

func TestGzipCompressAndDecompress(t *testing.T) {
	type args struct {
		data []byte
	}
	tests := []struct {
		name    string
		args    args
		want    []byte
		wantErr bool
	}{
		{
			name:    "#1",
			args:    args{[]byte("this is a test string")},
			want:    []byte("this is a test string"),
			wantErr: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			out, err := GzipCompress(tt.args.data)
			if err != nil {
				t.Errorf("GzipCompress() error = %v, wantErr %v", err, tt.wantErr)
				return
			}

			got, err := GzipDecompress(out)
			if (err != nil) != tt.wantErr {
				t.Errorf("GzipDecompress() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("GzipDecompress() got = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestZlibCompressAndDecompress(t *testing.T) {
	type args struct {
		data []byte
	}
	tests := []struct {
		name    string
		args    args
		want    []byte
		wantErr bool
	}{
		{
			name:    "#1",
			args:    args{[]byte("this is a test string")},
			want:    []byte("this is a test string"),
			wantErr: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			out, err := ZlibCompress(tt.args.data)
			if err != nil {
				t.Errorf("ZlibCompress() error = %v, wantErr %v", err, tt.wantErr)
				return
			}

			got, err := ZlibDecompress(out)
			if (err != nil) != tt.wantErr {
				t.Errorf("ZlibDecompress() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("ZlibDecompress() got = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestFlateCompressAndDecompress(t *testing.T) {
	type args struct {
		data []byte
	}
	tests := []struct {
		name    string
		args    args
		want    []byte
		wantErr bool
	}{
		{
			name:    "#1",
			args:    args{[]byte("this is a test string")},
			want:    []byte("this is a test string"),
			wantErr: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			out, err := FlateCompress(tt.args.data)
			if err != nil {
				t.Errorf("FlateCompress() error = %v, wantErr %v", err, tt.wantErr)
				return
			}

			got, err := FlateDecompress(out)
			if (err != nil) != tt.wantErr {
				t.Errorf("FlateDecompress() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("FlateDecompress() got = %v, want %v", got, tt.want)
			}
		})
	}
}
