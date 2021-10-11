package gzip

import (
	"reflect"
	"testing"
)

func TestZipCompressAndDecompress(t *testing.T) {
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
			out, err := ZipCompress(tt.args.data)
			if err != nil {
				t.Errorf("ZipCompress() error = %v, wantErr %v", err, tt.wantErr)
				return
			}

			got, err := ZipDecompress(out)
			if (err != nil) != tt.wantErr {
				t.Errorf("ZipDecompress() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("ZipDecompress() got = %v, want %v", got, tt.want)
			}
		})
	}
}
