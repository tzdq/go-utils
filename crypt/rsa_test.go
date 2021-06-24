package crypt

import (
	"go-utils/file"
	"reflect"
	"testing"
)

// "这是一个测试rsa的字符串，用于测试文件，rsa np!!!"
var rsaPlaintextTest = []byte{232, 191, 153, 230, 152, 175, 228, 184, 128, 228, 184, 170, 230, 181, 139, 232, 175,
	149, 114, 115, 97, 231, 154, 132, 229, 173, 151, 231, 172, 166, 228, 184, 178, 239, 188, 140, 231, 148, 168, 228,
	186, 142, 230, 181, 139, 232, 175, 149, 230, 150, 135, 228, 187, 182, 239, 188, 140, 114, 115, 97, 32, 110, 112,
	33, 33, 33}

func TestRSAEncryptAndDecrypt(t *testing.T) {
	type args struct {
		plaintext []byte
	}
	tests := []struct {
		name string
		args args
		want []byte
	}{
		{
			name: "Normal",
			args: args{
				plaintext: rsaPlaintextTest,
			},
			want: rsaPlaintextTest,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			pubKey, priKey, err := RSAGenKey(2048)
			if err != nil {
				t.Errorf("RSAGenKey() error = %v", err)
				return
			}

			ciphertext, err := RSAEncrypt(tt.args.plaintext, pubKey)
			if err != nil {
				t.Errorf("RSAEncrypt() error = %v", err)
				return
			}
			plaintext, err := RSADecrypt(ciphertext, priKey)
			if err != nil {
				t.Errorf("RSADecrypt() error = %v", err)
				return
			}

			if !reflect.DeepEqual(plaintext, tt.want) {
				t.Errorf("TestRSAEncryptAndDecrypt() got = %v, want %v", plaintext, tt.want)
			}
		})
	}
}

func TestRSAEncryptAndDecryptFromFile(t *testing.T) {
	type args struct {
		plaintext      []byte
		privateKeyPath string
		publicKeyPath  string
	}
	tests := []struct {
		name string
		args args
		want []byte
	}{
		{
			name: "#1",
			args: args{
				plaintext:      rsaPlaintextTest,
				privateKeyPath: "testdata/private.pem",
				publicKeyPath:  "testdata/public.pem",
			},
			want: rsaPlaintextTest,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if !file.IsExist(tt.args.publicKeyPath) || !file.IsExist(tt.args.privateKeyPath) {
				err := RSAGenKeyToFile(2048, tt.args.publicKeyPath, tt.args.privateKeyPath)
				if err != nil {
					t.Errorf("RSAGenKeyToFile() error = %v", err)
					return
				}
			}

			ciphertext, err := RSAEncryptFromFile(tt.args.plaintext, tt.args.publicKeyPath)
			if err != nil {
				t.Errorf("RSAEncryptFromFile() error = %v", err)
				return
			}
			plaintext, err := RSADecryptFromFile(ciphertext, tt.args.privateKeyPath)
			if err != nil {
				t.Errorf("RSADecryptFromFile() error = %v", err)
				return
			}
			if !reflect.DeepEqual(plaintext, tt.want) {
				t.Errorf("TestRSAEncryptAndDecryptFromFile() got = %v, want %v", plaintext, tt.want)
			}
		})
	}
}

func TestRSAGenKeyToFile(t *testing.T) {
	type args struct {
		bits           int
		privateKeyPath string
		publicKeyPath  string
	}
	tests := []struct {
		name    string
		args    args
		wantErr bool
	}{
		{
			name:    "#1",
			args:    args{2048, "testdata/private.pem", "testdata/public.pem"},
			wantErr: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			err := RSAGenKeyToFile(tt.args.bits, tt.args.publicKeyPath, tt.args.privateKeyPath)
			if (err != nil) != tt.wantErr {
				t.Errorf("RSAGenKeyToFile() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}

func TestRSASignAndVerySign(t *testing.T) {
	type args struct {
		data []byte
	}
	tests := []struct {
		name string
		args args
	}{
		{
			name: "#1",
			args: args{rsaPlaintextTest},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {

			pubKey, priKey, err := RSAGenKey(2048)
			if err != nil {
				t.Errorf("RSAGenKey() error = %v", err)
				return
			}

			signature, err := RSASign(tt.args.data, priKey)
			if err != nil {
				t.Errorf("RSASign() error = %v, priKey=%v,signature=%v", err, priKey, signature)
				return
			}
			if !RSAVerySign(tt.args.data, signature, pubKey) {
				t.Errorf("RSAVerySign() error = %v, pubKey=%v,signature=%v,data=%v",
					err, priKey, signature, tt.args.data)
				return
			}
		})
	}
}

func TestRSASignAndVerySignFromFile(t *testing.T) {
	type args struct {
		data []byte
	}
	tests := []struct {
		name string
		args args
	}{
		{
			name: "#1",
			args: args{rsaPlaintextTest},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			priKeyPath := "testdata/private.pem"
			pubKeyPath := "testdata/public.pem"
			if !file.IsExist(priKeyPath) || !file.IsExist(pubKeyPath) {
				err := RSAGenKeyToFile(2048, pubKeyPath, priKeyPath)
				if err != nil {
					t.Errorf("RSAGenKeyToFile() error = %v", err)
					return
				}
			}

			signature, err := RSASignFromFile(tt.args.data, priKeyPath)
			if err != nil {
				t.Errorf("RSASignFromFile() error = %v, priKeyPath %v", err, priKeyPath)
				return
			}
			if !RSAVerySignFromFile(tt.args.data, signature, pubKeyPath) {
				t.Errorf("RSASignFromFile() pubKeyPath = %v, signature %v", pubKeyPath, signature)
			}
		})
	}
}
