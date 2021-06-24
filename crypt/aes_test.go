package crypt

import (
	"reflect"
	"testing"
)

// 12fffsagjseoprefdfjdasdsooekalushds
var commonOriginData = []byte{49, 50, 102, 102, 102, 115, 97, 103, 106, 115, 101, 111, 112, 114, 101, 102, 100,
	102, 106, 100, 97, 115, 100, 115, 111, 111, 101, 107, 97, 108, 117, 115, 104, 100, 115}

// thisisakeythisisthisisakeythisis
var commonKey32 = []byte{116, 104, 105, 115, 105, 115, 97, 107, 101, 121, 116, 104, 105, 115, 105, 115, 116, 104, 105,
	115, 105, 115, 97, 107, 101, 121, 116, 104, 105, 115, 105, 115}

// thisisakeythisisaassddww
var commonKey24 = []byte{116, 104, 105, 115, 105, 115, 97, 107, 101, 121, 116, 104, 105, 115, 105, 115, 97, 97, 115,
	115, 100, 100, 119, 119}

// thisisakeythisis
var commonKey16 = []byte{116, 104, 105, 115, 105, 115, 97, 107, 101, 121, 116, 104, 105, 115, 105, 115}

// thisisakeythis
var invalidKey = []byte{116, 104, 105, 115, 105, 115, 97, 107, 101, 121, 116, 104, 105, 115}

func TestAesCBCDecrypt(t *testing.T) {
	type args struct {
		data []byte
		key  []byte
	}
	tests := []struct {
		name         string
		args         args
		wantOrigData []byte
		wantErr      bool
	}{
		{
			name: "KeyInvalidLen",
			args: args{
				data: []byte{127, 112, 89, 164, 136, 229, 232, 48, 142, 96, 51, 204, 63, 152, 61, 240},
				key:  invalidKey,
			},
			wantErr: true,
		},
		{
			name: "OriginDataEmpty",
			args: args{
				data: []byte{138, 136, 199, 43, 145, 112, 90, 206, 204, 182, 35, 44, 246, 46, 26, 101, 249, 2, 83,
					220, 202, 215, 31, 124, 249, 114, 39, 90, 73, 104, 20, 114},
				key: commonKey16,
			},
			wantOrigData: []byte{},
		},
		{
			name: "AesCBC16",
			args: args{
				data: []byte{193, 218, 180, 253, 33, 56, 45, 71, 51, 25, 4, 223, 49, 246, 89, 8, 71, 177, 61, 16, 145,
					253, 8, 199, 7, 170, 211, 132, 147, 195, 149, 231, 120, 26, 130, 10, 167, 139, 117, 217, 223, 68,
					216, 226, 12, 157, 38, 227, 16, 206, 50, 60, 95, 227, 210, 139, 155, 111, 136, 151, 37, 180, 71,
					178},
				key: commonKey16,
			},
			wantOrigData: commonOriginData,
		},
		{
			name: "AesCBC32",
			args: args{
				data: []byte{8, 29, 223, 161, 242, 152, 176, 249, 183, 102, 145, 42, 30, 246, 196, 160, 83, 119, 246,
					211, 182, 10, 132, 92, 56, 156, 21, 145, 202, 67, 132, 147, 94, 198, 124, 180, 176, 194, 9, 33,
					204, 74, 36, 167, 165, 234, 78, 5, 87, 7, 225, 248, 56, 43, 138, 130, 115, 218, 119, 60, 251, 182,
					217, 63},
				key: commonKey32,
			},
			wantOrigData: commonOriginData,
		},
		{
			name: "DataNotFullBlock",
			args: args{
				data: []byte{8, 29, 223, 161, 242, 152, 176, 249, 183, 102, 145, 42, 30, 246, 196, 160, 83, 119, 246,
					211, 182, 10, 132, 92, 56, 156, 21, 145, 202, 67, 132, 147, 94, 198, 124, 180, 176, 194, 9, 33,
					204, 74, 36, 167, 165, 234, 78, 5, 87, 7, 225, 248, 56, 43, 138, 130, 115, 218, 119, 60, 251, 182},
				key: commonKey32,
			},
			wantErr: true,
		},
		{
			name: "DataEmpty",
			args: args{
				data: nil,
				key:  commonKey32,
			},
			wantErr: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			gotOrigData, err := aesCBCDecrypt(tt.args.data, tt.args.key)
			if (err != nil) != tt.wantErr {
				t.Errorf("aesCBCDecrypt() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(gotOrigData, tt.wantOrigData) {
				t.Errorf("aesCBCDecrypt() gotOrigData = %v, want %v", string(gotOrigData), string(tt.wantOrigData))
				t.Errorf("aesCBCDecrypt() gotOrigData = %v, want %v", gotOrigData, tt.wantOrigData)
			}
		})
	}
}

func TestAesCBCEncrypt(t *testing.T) {
	type args struct {
		data []byte
		key  []byte
	}
	tests := []struct {
		name    string
		args    args
		wantErr bool
	}{
		{
			name: "KeyInvalidLen",
			args: args{
				data: commonOriginData,
				key:  invalidKey,
			},
			wantErr: true,
		},
		{
			name: "DataEmpty",
			args: args{
				data: []byte{},
				key:  commonKey16,
			},
		},
		{
			name: "AesCBC16",
			args: args{
				data: commonOriginData,
				key:  commonKey16,
			},
		},
		{
			name: "AesCBC32",
			args: args{
				data: commonOriginData,
				key:  commonKey32,
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := aesCBCEncrypt(tt.args.data, tt.args.key)
			t.Logf("%v", got)
			if (err != nil) != tt.wantErr {
				t.Errorf("aesCBCEncrypt() error = %v, wantErr %v got %v", err, tt.wantErr, got)
				return
			}
		})
	}
}

func TestAesECBDecrypt(t *testing.T) {
	type args struct {
		ciphertext []byte
		key        []byte
	}
	tests := []struct {
		name         string
		args         args
		wantOrigData []byte
		wantErr      bool
	}{
		{
			name: "KeyInvalidLen",
			args: args{
				ciphertext: commonOriginData,
				key:        invalidKey,
			},
			wantOrigData: nil,
			wantErr:      true,
		},
		{
			name: "OriginDataEmpty",
			args: args{
				ciphertext: []byte{117, 194, 79, 252, 43, 33, 176, 117, 171, 17, 155, 181, 237, 8, 196, 60},
				key:        commonKey16,
			},
			wantOrigData: []byte{},
		},
		{
			name: "AesECB16",
			args: args{
				ciphertext: []byte{3, 198, 29, 111, 229, 172, 215, 137, 79, 42, 78, 212, 226, 4, 58, 139, 237, 81,
					107, 115, 98, 121, 143, 82, 209, 224, 255, 253, 237, 96, 133, 248, 103, 152, 111, 226, 245, 253,
					11, 203, 170, 4, 14, 10, 46, 215, 150, 163},
				key: commonKey16,
			},
			wantOrigData: commonOriginData,
		},
		{
			name: "DataNotFullBlock",
			args: args{
				ciphertext: []byte{3, 198, 29, 111, 229, 172, 215, 137, 79, 42, 78, 212, 226, 4, 58, 139, 237, 81,
					107, 115, 98, 121, 143, 82, 209, 224, 255, 253, 237, 96, 133, 248, 103, 152, 111, 226, 245, 253,
					11, 203, 170, 4, 14, 10, 46, 215, 150, 163, 163},
				key: commonKey16,
			},
			wantOrigData: nil,
			wantErr:      true,
		},
		{
			name: "DataEmpty",
			args: args{
				ciphertext: []byte{},
				key:        commonKey16,
			},
			wantOrigData: nil,
			wantErr:      true,
		},
		{
			name: "AesECB32",
			args: args{
				ciphertext: []byte{155, 177, 132, 128, 162, 246, 140, 57, 45, 161, 118, 50, 202, 138, 148, 141, 29,
					26, 224, 108, 156, 173, 244, 192, 45, 83, 171, 226, 35, 235, 58, 179, 184, 53, 218, 225, 89, 172,
					252, 193, 219, 191, 123, 182, 115, 93, 194, 245},
				key: commonKey32,
			},
			wantOrigData: commonOriginData,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			gotOrigData, err := aesECBDecrypt(tt.args.ciphertext, tt.args.key)
			if (err != nil) != tt.wantErr {
				t.Errorf("aesECBDecrypt() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(gotOrigData, tt.wantOrigData) {
				t.Errorf("aesECBDecrypt() gotOrigData = %v, want %v", gotOrigData, tt.wantOrigData)
			}
		})
	}
}

func TestAesECBEncrypt(t *testing.T) {
	type args struct {
		plaintext []byte
		key       []byte
	}
	tests := []struct {
		name        string
		args        args
		wantCrypted []byte
		wantErr     bool
	}{
		{
			name: "KeyInvalidLen",
			args: args{
				plaintext: commonOriginData,
				key:       invalidKey,
			},
			wantErr: true,
		},
		{
			name: "OriginDataEmpty",
			args: args{
				plaintext: []byte{},
				key:       commonKey16,
			},
			wantCrypted: []byte{117, 194, 79, 252, 43, 33, 176, 117, 171, 17, 155, 181, 237, 8, 196, 60},
		},
		{
			name: "AesECB16",
			args: args{
				plaintext: commonOriginData,
				key:       commonKey16,
			},
			wantCrypted: []byte{3, 198, 29, 111, 229, 172, 215, 137, 79, 42, 78, 212, 226, 4, 58, 139, 237, 81, 107,
				115, 98, 121, 143, 82, 209, 224, 255, 253, 237, 96, 133, 248, 103, 152, 111, 226, 245, 253, 11, 203,
				170, 4, 14, 10, 46, 215, 150, 163},
		},
		{
			name: "AesECB32",
			args: args{
				plaintext: commonOriginData,
				key:       commonKey32,
			},
			wantCrypted: []byte{155, 177, 132, 128, 162, 246, 140, 57, 45, 161, 118, 50, 202, 138, 148, 141, 29, 26,
				224, 108, 156, 173, 244, 192, 45, 83, 171, 226, 35, 235, 58, 179, 184, 53, 218, 225, 89, 172, 252, 193,
				219, 191, 123, 182, 115, 93, 194, 245},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			gotCrypted, err := aesECBEncrypt(tt.args.plaintext, tt.args.key)
			if (err != nil) != tt.wantErr {
				t.Errorf("aesECBEncrypt() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(gotCrypted, tt.wantCrypted) {
				t.Errorf("aesECBEncrypt() gotCrypted = %v, want %v", gotCrypted, tt.wantCrypted)
			}
		})
	}
}

func TestAesStreamEncrypt(t *testing.T) {
	type args struct {
		plaintext []byte
		key       []byte
		aw        AesMode
	}
	tests := []struct {
		name    string
		args    args
		wantErr bool
	}{
		{
			name: "KeyInvalidLen",
			args: args{
				plaintext: commonOriginData,
				key:       invalidKey,
			},
			wantErr: true,
		},
		{
			name: "OFBOriginDataEmpty",
			args: args{
				plaintext: []byte{},
				key:       commonKey16,
				aw:        AesModeOFB,
			},
		},
		{
			name: "AesModeInvalid",
			args: args{
				plaintext: commonOriginData,
				key:       commonKey16,
			},
			wantErr: true,
		},
		{
			name: "CFBOriginDataEmpty",
			args: args{
				plaintext: []byte{},
				key:       commonKey16,
				aw:        AesModeCFB,
			},
		},
		{
			name: "CTROriginDataEmpty",
			args: args{
				plaintext: []byte{},
				key:       commonKey16,
				aw:        AesModeCTR,
			},
		},
		{
			name: "AesOFB16",
			args: args{
				plaintext: commonOriginData,
				key:       commonKey16,
				aw:        AesModeOFB,
			},
		},
		{
			name: "AesOFB24",
			args: args{
				plaintext: commonOriginData,
				key:       commonKey24,
				aw:        AesModeOFB,
			},
		},
		{
			name: "AesOFB32",
			args: args{
				plaintext: commonOriginData,
				key:       commonKey32,
				aw:        AesModeOFB,
			},
		},
		{
			name: "AesCTR16",
			args: args{
				plaintext: commonOriginData,
				key:       commonKey16,
				aw:        AesModeCTR,
			},
		},
		{
			name: "AesCTR24",
			args: args{
				plaintext: commonOriginData,
				key:       commonKey24,
				aw:        AesModeCTR,
			},
		},
		{
			name: "AesCTR32",
			args: args{
				plaintext: commonOriginData,
				key:       commonKey32,
				aw:        AesModeCTR,
			},
		},
		{
			name: "AesCFB16",
			args: args{
				plaintext: commonOriginData,
				key:       commonKey16,
				aw:        AesModeCFB,
			},
		},
		{
			name: "AesCFB24",
			args: args{
				plaintext: commonOriginData,
				key:       commonKey24,
				aw:        AesModeCFB,
			},
		},
		{
			name: "AesCFB32",
			args: args{
				plaintext: commonOriginData,
				key:       commonKey32,
				aw:        AesModeCFB,
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := aesStreamEncrypt(tt.args.plaintext, tt.args.key, tt.args.aw)
			if (err != nil) != tt.wantErr {
				t.Errorf("aesStreamEncrypt() error = %v, wantErr %v, got %v", err, tt.wantErr, got)
				return
			}
		})
	}
}

func TestAesStreamDecrypt(t *testing.T) {
	type args struct {
		ciphertext []byte
		key        []byte
		aw         AesMode
	}
	tests := []struct {
		name         string
		args         args
		wantOrigData []byte
		wantErr      bool
	}{
		{
			name: "DataLenInvalid",
			args: args{
				ciphertext: []byte{97, 12, 203, 228, 173, 199, 170, 35, 48, 205, 214, 87, 109, 107, 20},
				key:        commonKey16,
			},
			wantErr: true,
		},
		{
			name: "KeyInvalidLen",
			args: args{
				ciphertext: commonOriginData,
				key:        invalidKey,
			},
			wantErr: true,
		},
		{
			name: "AesModeInvalid",
			args: args{
				ciphertext: []byte{100, 91, 8, 148, 175, 99, 227, 251, 149, 110, 190, 15, 127, 159, 139, 171},
				key:        commonKey16,
			},
			wantErr: true,
		},
		{
			name: "OFBOriginDataEmpty",
			args: args{
				ciphertext: []byte{100, 91, 8, 148, 175, 99, 227, 251, 149, 110, 190, 15, 127, 159, 139, 171},
				key:        commonKey16,
				aw:         AesModeOFB,
			},
			wantOrigData: []byte{},
		},
		{
			name: "AesOFB16",
			args: args{
				ciphertext: []byte{252, 139, 244, 172, 162, 83, 166, 189, 208, 42, 244, 149, 135, 114, 13, 191, 54,
					255, 48, 170, 173, 26, 212, 152, 133, 9, 232, 40, 89, 23, 55, 81, 6, 183, 209, 32, 14, 170, 205,
					195, 95, 77, 33, 167, 78, 48, 24, 197, 214, 130, 152},
				key: commonKey16,
				aw:  AesModeOFB,
			},
			wantOrigData: commonOriginData,
		},
		{
			name: "AesOFB24",
			args: args{
				ciphertext: []byte{59, 191, 131, 232, 140, 136, 41, 183, 120, 190, 68, 220, 103, 229, 202, 218, 103,
					200, 224, 107, 88, 7, 79, 114, 121, 77, 26, 215, 81, 193, 13, 143, 255, 5, 219, 87, 130, 93, 56,
					95, 225, 17, 193, 11, 114, 60, 129, 105, 167, 12, 192},
				key: commonKey24,
				aw:  AesModeOFB,
			},
			wantOrigData: commonOriginData,
		},
		{
			name: "AesOFB32",
			args: args{
				ciphertext: []byte{196, 130, 117, 32, 137, 135, 143, 111, 55, 192, 1, 201, 84, 224, 42, 184, 129, 137,
					36, 247, 64, 185, 144, 94, 211, 46, 89, 76, 167, 201, 246, 247, 189, 94, 216, 218, 101, 113, 53,
					189, 147, 124, 167, 196, 79, 183, 159, 1, 153, 121, 245},
				key: commonKey32,
				aw:  AesModeOFB,
			},
			wantOrigData: commonOriginData,
		},
		{
			name: "AesCTR16",
			args: args{
				ciphertext: []byte{199, 226, 133, 56, 97, 240, 0, 186, 250, 63, 175, 94, 112, 132, 144, 163, 115, 73,
					175, 21, 219, 229, 153, 31, 205, 207, 70, 172, 40, 220, 40, 197, 178, 189, 106, 222, 123, 15, 183,
					108, 161, 161, 93, 209, 118, 251, 44, 57, 72, 6, 172},
				key: commonKey16,
				aw:  AesModeCTR,
			},
			wantOrigData: commonOriginData,
		},
		{
			name: "AesCTR24",
			args: args{
				ciphertext: []byte{30, 223, 26, 31, 232, 147, 6, 234, 149, 215, 114, 58, 48, 18, 250, 83, 87, 40, 245,
					138, 66, 97, 89, 140, 14, 50, 54, 53, 126, 205, 172, 81, 13, 87, 151, 37, 229, 23, 135, 57, 113,
					175, 123, 25, 197, 101, 255, 141, 150, 127, 140},
				key: commonKey24,
				aw:  AesModeCTR,
			},
			wantOrigData: commonOriginData,
		},
		{
			name: "AesCTR32",
			args: args{
				ciphertext: []byte{103, 101, 46, 183, 60, 5, 216, 3, 159, 8, 6, 211, 193, 160, 202, 71, 207, 88, 90,
					27, 19, 47, 237, 229, 85, 218, 58, 40, 144, 93, 9, 94, 86, 248, 13, 186, 15, 237, 28, 106, 86, 179,
					244, 244, 64, 180, 185, 152, 129, 19, 88},
				key: commonKey32,
				aw:  AesModeCTR,
			},
			wantOrigData: commonOriginData,
		},
		{
			name: "AesCFB16",
			args: args{
				ciphertext: []byte{224, 167, 75, 239, 91, 111, 18, 207, 61, 122, 19, 194, 209, 89, 162, 233, 136, 235,
					228, 8, 56, 236, 11, 117, 131, 242, 181, 145, 45, 68, 123, 83, 180, 252, 152, 196, 71, 189, 59, 29,
					98, 251, 117, 59, 187, 125, 9, 55, 219, 45, 133},
				key: commonKey16,
				aw:  AesModeCFB,
			},
			wantOrigData: commonOriginData,
		},
		{
			name: "AesCFB24",
			args: args{
				ciphertext: []byte{36, 99, 255, 149, 9, 202, 154, 12, 173, 225, 65, 216, 5, 133, 95, 131, 33, 198, 135,
					57, 215, 25, 233, 179, 142, 204, 185, 162, 180, 132, 58, 245, 209, 100, 164, 2, 230, 196, 229, 207,
					192, 146, 131, 254, 33, 227, 66, 64, 13, 119, 25},
				key: commonKey24,
				aw:  AesModeCFB,
			},
			wantOrigData: commonOriginData,
		},
		{
			name: "AesCFB32",
			args: args{
				ciphertext: []byte{124, 0, 9, 111, 113, 188, 168, 255, 108, 63, 135, 90, 247, 27, 195, 178, 76, 253,
					65, 64, 40, 105, 108, 17, 255, 120, 74, 173, 16, 212, 236, 63, 245, 166, 71, 76, 142, 46, 131, 115,
					64, 191, 113, 15, 195, 29, 45, 31, 92, 60, 27},
				key: commonKey32,
				aw:  AesModeCFB,
			},
			wantOrigData: commonOriginData,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := aesStreamDecrypt(tt.args.ciphertext, tt.args.key, tt.args.aw)
			if (err != nil) != tt.wantErr {
				t.Errorf("aesStreamDecrypt() error = %v, wantErr %v, got %v", err, tt.wantErr, got)
				return
			}
			if !reflect.DeepEqual(got, tt.wantOrigData) {
				t.Errorf("aesStreamDecrypt() gotOrigData = %v, want %v", got, tt.wantOrigData)
			}
		})
	}
}

func TestAESDecrypt(t *testing.T) {
	type args struct {
		ciphertext []byte
		key        []byte
		aw         AesMode
	}
	tests := []struct {
		name         string
		args         args
		wantOrigData []byte
		wantErr      bool
	}{
		{
			name: "CBC",
			args: args{
				ciphertext: []byte{8, 29, 223, 161, 242, 152, 176, 249, 183, 102, 145, 42, 30, 246, 196, 160, 83, 119,
					246, 211, 182, 10, 132, 92, 56, 156, 21, 145, 202, 67, 132, 147, 94, 198, 124, 180, 176, 194, 9,
					33, 204, 74, 36, 167, 165, 234, 78, 5, 87, 7, 225, 248, 56, 43, 138, 130, 115, 218, 119, 60, 251,
					182, 217, 63},
				key: commonKey32,
				aw:  AesModeCBC,
			},
			wantOrigData: commonOriginData,
		},
		{
			name: "ECB",
			args: args{
				ciphertext: []byte{3, 198, 29, 111, 229, 172, 215, 137, 79, 42, 78, 212, 226, 4, 58, 139, 237, 81,
					107, 115, 98, 121, 143, 82, 209, 224, 255, 253, 237, 96, 133, 248, 103, 152, 111, 226, 245, 253,
					11, 203, 170, 4, 14, 10, 46, 215, 150, 163},
				key: commonKey16,
				aw:  AesModeECB,
			},
			wantOrigData: commonOriginData,
		},
		{
			name: "CFB",
			args: args{
				ciphertext: []byte{124, 0, 9, 111, 113, 188, 168, 255, 108, 63, 135, 90, 247, 27, 195, 178, 76, 253,
					65, 64, 40, 105, 108, 17, 255, 120, 74, 173, 16, 212, 236, 63, 245, 166, 71, 76, 142, 46, 131,
					115, 64, 191, 113, 15, 195, 29, 45, 31, 92, 60, 27},
				key: commonKey32,
				aw:  AesModeCFB,
			},
			wantOrigData: commonOriginData,
		},
		{
			name: "CTR",
			args: args{
				ciphertext: []byte{30, 223, 26, 31, 232, 147, 6, 234, 149, 215, 114, 58, 48, 18, 250, 83, 87, 40, 245,
					138, 66, 97, 89, 140, 14, 50, 54, 53, 126, 205, 172, 81, 13, 87, 151, 37, 229, 23, 135, 57, 113,
					175, 123, 25, 197, 101, 255, 141, 150, 127, 140},
				key: commonKey24,
				aw:  AesModeCTR,
			},
			wantOrigData: commonOriginData,
		},
		{
			name: "OFB",
			args: args{
				ciphertext: []byte{59, 191, 131, 232, 140, 136, 41, 183, 120, 190, 68, 220, 103, 229, 202, 218, 103,
					200, 224, 107, 88, 7, 79, 114, 121, 77, 26, 215, 81, 193, 13, 143, 255, 5, 219, 87, 130, 93, 56,
					95, 225, 17, 193, 11, 114, 60, 129, 105, 167, 12, 192},
				key: commonKey24,
				aw:  AesModeOFB,
			},
			wantOrigData: commonOriginData,
		},
		{
			name: "unknown",
			args: args{
				ciphertext: []byte{17, 150, 44, 116, 120, 91, 240, 219, 26, 199, 152, 128, 184, 118, 155, 65, 240, 4,
					240, 71, 14, 25, 39, 152, 160, 242, 248, 23, 203, 110, 18, 74, 41, 212, 105, 27, 174, 197, 129,
					237, 197, 134, 206, 246, 109, 233, 85, 47},
				key: commonKey16,
			},
			wantOrigData: nil,
			wantErr:      true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			gotOrigData, err := AESDecrypt(tt.args.ciphertext, tt.args.key, tt.args.aw)
			if (err != nil) != tt.wantErr {
				t.Errorf("AESDecrypt() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(gotOrigData, tt.wantOrigData) {
				t.Errorf("AESDecrypt() gotOrigData = %v, want %v", gotOrigData, tt.wantOrigData)
			}
		})
	}
}

func TestAESEncrypt(t *testing.T) {
	type args struct {
		plaintext []byte
		key       []byte
		am        AesMode
	}
	tests := []struct {
		name         string
		args         args
		compareValue bool
		wantCrypted  []byte
		wantErr      bool
	}{
		{
			name: "KeyInvalidLen",
			args: args{
				plaintext: commonOriginData,
				key:       invalidKey,
				am:        AesModeCBC,
			},
			wantErr: true,
		},
		{
			name: "ECB",
			args: args{
				plaintext: commonOriginData,
				key:       commonKey32,
				am:        AesModeECB,
			},
			compareValue: true,
			wantCrypted: []byte{155, 177, 132, 128, 162, 246, 140, 57, 45, 161, 118, 50, 202, 138, 148, 141, 29, 26,
				224, 108, 156, 173, 244, 192, 45, 83, 171, 226, 35, 235, 58, 179, 184, 53, 218, 225, 89, 172, 252,
				193, 219, 191, 123, 182, 115, 93, 194, 245},
		},
		{
			name: "CTR",
			args: args{
				plaintext: commonOriginData,
				key:       commonKey32,
				am:        AesModeCTR,
			},
		},
		{
			name: "unknown",
			args: args{
				plaintext: commonOriginData,
				key:       commonKey32,
			},
			wantErr: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := AESEncrypt(tt.args.plaintext, tt.args.key, tt.args.am)
			if (err != nil) != tt.wantErr {
				t.Errorf("AESEncrypt() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if tt.compareValue && !reflect.DeepEqual(got, tt.wantCrypted) {
				t.Errorf("AESEncrypt() gotCrypted = %v, want %v", got, tt.wantCrypted)
			}
		})
	}
}
