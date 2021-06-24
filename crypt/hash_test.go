package crypt

import (
	"crypto/sha1"
	"crypto/sha256"
	"hash"
	"reflect"
	"testing"
)

// "this is a test!!!"
var hashCommonTest = []byte{116, 104, 105, 115, 32, 105, 115, 32, 97, 32, 116, 101, 115, 116, 33, 33, 33}

// password
var hashPasswordTest = []byte{112, 97, 115, 115, 119, 111, 114, 100}

// salt
var hashSaltTest = []byte{115, 97, 108, 116}

// thisisakey
var hashKeyTest = []byte{116, 104, 105, 115, 105, 115, 97, 107, 101, 121}

func TestHmacBytes(t *testing.T) {
	type args struct {
		data []byte
		key  []byte
		ht   HashType
	}
	tests := []struct {
		name string
		args args
		want []byte
	}{
		{
			name: "HMacSha1",
			args: args{hashCommonTest, hashKeyTest, HtSha1},
			want: []byte{196, 56, 40, 166, 234, 47, 203, 30, 216, 209, 161, 209, 251, 113, 209, 135, 108,
				199, 232, 65},
		},
		{
			name: "HMacSha224",
			args: args{hashCommonTest, hashKeyTest, HtSha224},
			want: []byte{71, 225, 255, 16, 177, 183, 248, 3, 7, 46, 57, 29, 16, 115, 52, 189, 84, 145, 27, 180, 21,
				230, 54, 67, 240, 128, 135, 255},
		},
		{
			name: "HMacSha256",
			args: args{hashCommonTest, hashKeyTest, HtSha256},
			want: []byte{113, 96, 62, 157, 48, 74, 139, 34, 233, 156, 187, 143, 130, 2, 207, 11, 3, 248, 235, 210,
				191, 142, 176, 47, 31, 52, 167, 98, 207, 2, 67, 155},
		},
		{
			name: "HMacSha384",
			args: args{hashCommonTest, hashKeyTest, HtSha384},
			want: []byte{184, 150, 141, 178, 179, 83, 223, 130, 33, 85, 114, 58, 180, 252, 31, 249, 28, 97, 250, 215,
				170, 209, 138, 106, 22, 131, 44, 11, 110, 125, 182, 57, 249, 82, 83, 115, 118, 181, 57, 248, 21, 46,
				102, 165, 135, 237, 9, 49},
		},
		{
			name: "HMacSha512",
			args: args{hashCommonTest, hashKeyTest, HtSha512},
			want: []byte{149, 225, 126, 79, 32, 12, 30, 132, 65, 136, 192, 96, 252, 171, 242, 220, 23, 0, 146, 16, 116,
				2, 236, 100, 166, 191, 97, 55, 145, 151, 173, 81, 11, 170, 13, 59, 173, 26, 81, 125, 54, 241, 12, 169,
				155, 238, 86, 157, 177, 5, 163, 203, 48, 170, 237, 111, 163, 129, 235, 48, 83, 100, 216, 217},
		},
		{
			name: "HmacMd5",
			args: args{hashCommonTest, hashKeyTest, HtMD5},
			want: []byte{186, 189, 237, 130, 87, 77, 16, 127, 113, 146, 132, 21, 98, 84, 69, 13},
		},
		{
			name: "3Des",
			args: args{hashCommonTest, hashKeyTest, 0},
			want: []byte{},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			t.Logf("%v", tt.args.key)
			if got := HmacBytes(tt.args.data, tt.args.key, tt.args.ht); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("HmacBytes(%d) = %v, want %v", tt.args.ht, got, tt.want)
			}
		})
	}
}

func TestHashBytes(t *testing.T) {
	type args struct {
		data []byte
		ht   HashType
	}
	tests := []struct {
		name string
		args args
		want []byte
	}{
		{
			name: "MD5",
			args: args{hashCommonTest, HtMD5},
			want: []byte{220, 83, 226, 101, 251, 176, 123, 179, 201, 217, 254, 204, 120, 117, 234, 190},
		},
		{
			name: "Sha1",
			args: args{hashCommonTest, HtSha1},
			want: []byte{102, 166, 120, 54, 150, 117, 116, 109, 13, 45, 153, 167, 124, 206, 170, 173, 0, 29, 74, 24},
		},
		{
			name: "Sha224",
			args: args{hashCommonTest, HtSha224},
			want: []byte{159, 197, 240, 87, 242, 105, 146, 248, 173, 234, 72, 119, 221, 158, 149, 178,
				155, 88, 95, 23, 249, 219, 93, 213, 220, 155, 144, 7},
		},
		{
			name: "Sha256",
			args: args{hashCommonTest, HtSha256},
			want: []byte{223, 132, 47, 106, 85, 24, 126, 72, 183, 230, 118, 175, 14, 123, 252, 221, 5, 83,
				143, 199, 34, 24, 221, 194, 33, 14, 152, 120, 35, 76, 229, 203},
		},
		{
			name: "Sha384",
			args: args{hashCommonTest, HtSha384},
			want: []byte{173, 27, 110, 51, 144, 40, 119, 48, 232, 229, 160, 2, 17, 213, 197, 182, 43, 132, 227, 247,
				34, 48, 77, 83, 200, 23, 242, 237, 118, 184, 152, 101, 114, 5, 84, 195, 59, 67, 174, 87, 13, 28, 124,
				209, 226, 14, 55, 213},
		},
		{
			name: "Sha512",
			args: args{hashCommonTest, HtSha512},
			want: []byte{23, 201, 62, 103, 146, 19, 144, 146, 89, 68, 176, 251, 149, 13, 175, 207, 182, 163, 192, 207,
				128, 140, 184, 48, 149, 247, 126, 90, 140, 13, 86, 93, 248, 139, 109, 210, 86, 232, 159, 227, 183, 149,
				33, 139, 194, 245, 43, 4, 152, 172, 200, 149, 124, 50, 116, 135, 73, 255, 86, 55, 61, 226, 140, 200},
		},
		{
			name: "FNV1-32",
			args: args{hashCommonTest, HtFnv32},
			want: []byte{137, 105, 135, 181},
		},
		{
			name: "FNV1A-32",
			args: args{hashCommonTest, HtFnvA32},
			want: []byte{76, 159, 215, 85},
		},
		{
			name: "FNV1-64",
			args: args{hashCommonTest, HtFnv64},
			want: []byte{154, 127, 81, 0, 106, 136, 175, 181},
		},
		{
			name: "FNV1A-64",
			args: args{hashCommonTest, HtFnvA64},
			want: []byte{115, 87, 204, 74, 175, 33, 44, 85},
		},
		{
			name: "FNV1-128",
			args: args{hashCommonTest, HtFnv128},
			want: []byte{133, 233, 249, 51, 164, 174, 191, 185, 248, 227, 255, 92, 158, 147, 159, 149},
		},
		{
			name: "FNV1A-128",
			args: args{hashCommonTest, HtFnvA128},
			want: []byte{36, 154, 242, 233, 197, 38, 16, 196, 23, 67, 220, 235, 28, 169, 151, 245},
		},
		{
			name: "3Des",
			args: args{hashCommonTest, 0},
			want: []byte{},
		},
		{
			name: "Adler32",
			args: args{hashCommonTest, HtAdler32},
			want: []byte{54, 59, 5, 121},
		},
		{
			name: "Crc32",
			args: args{hashCommonTest, HtCrc32},
			want: []byte{73, 125, 253, 75},
		},
		{
			name: "Crc64ISO",
			args: args{hashCommonTest, HtCrc64ISO},
			want: []byte{94, 104, 166, 5, 217, 188, 197, 247},
		},
		{
			name: "Crc64ECMA",
			args: args{hashCommonTest, HtCrc64ECMA},
			want: []byte{159, 71, 192, 210, 75, 227, 214, 169},
		},
		{
			name: "Time33",
			args: args{hashCommonTest, HtTime33},
			want: []byte{120, 28, 75, 93},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			t.Logf("%v", ToHexString(HashBytes(tt.args.data, tt.args.ht)))
			if got := HashBytes(tt.args.data, tt.args.ht); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("HashBytes(%d) = %v, want %v", tt.args.ht, got, tt.want)
			}
		})
	}
}

func TestToHexString(t *testing.T) {
	type args struct {
		data []byte
	}
	tests := []struct {
		name string
		args args
		want string
	}{

		{
			name: "#1", args: args{[]byte{220, 83, 226, 101, 251, 176, 123, 179, 201, 217, 254, 204, 120,
				117, 234, 190}},
			want: "dc53e265fbb07bb3c9d9fecc7875eabe",
		},
		{
			name: "#2", args: args{[]byte{}},
			want: "",
		},
		{
			name: "#3", args: args{[]byte{133, 233, 249, 51, 164, 174, 191, 185, 248, 227, 255, 92, 158, 147,
				159, 149}},
			want: "85e9f933a4aebfb9f8e3ff5c9e939f95",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := ToHexString(tt.args.data); got != tt.want {
				t.Errorf("ToHexString() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestMd5File(t *testing.T) {
	type args struct {
		path string
	}
	tests := []struct {
		name string
		args args
		want string
	}{
		{
			name: "Md5File",
			args: args{"./random.go"},
			want: "d5ded8075552012f35ed146087255492",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := Md5File(tt.args.path); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("Md5File() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestPBKDF2(t *testing.T) {
	type args struct {
		password []byte
		salt     []byte
		iterCnt  uint32
		keyLen   uint32
		fn       func() hash.Hash
	}
	tests := []struct {
		name string
		args args
		want []byte
	}{
		{
			name: "sha1",
			args: args{hashPasswordTest, hashSaltTest, 1, 20, sha1.New},
			want: []byte{0x0c, 0x60, 0xc8, 0x0f, 0x96, 0x1f, 0x0e, 0x71, 0xf3, 0xa9, 0xb5, 0x24, 0xaf, 0x60, 0x12,
				0x06, 0x2f, 0xe0, 0x37, 0xa6},
		},
		{
			name: "sha1#4096",
			args: args{hashPasswordTest, hashSaltTest, 4096, 20, sha1.New},
			want: []byte{0x4b, 0x00, 0x79, 0x01, 0xb7, 0x65, 0x48, 0x9a, 0xbe, 0xad, 0x49, 0xd9, 0x26, 0xf7, 0x21,
				0xd0, 0x65, 0xa4, 0x29, 0xc1},
		},
		{
			name: "FuncNil",
			args: args{hashPasswordTest, hashSaltTest, 4096, 20, nil},
			want: []byte{0x4b, 0x00, 0x79, 0x01, 0xb7, 0x65, 0x48, 0x9a, 0xbe, 0xad, 0x49, 0xd9, 0x26, 0xf7, 0x21,
				0xd0, 0x65, 0xa4, 0x29, 0xc1},
		},
		{
			name: "sha256#4096",
			args: args{[]byte("passwordPASSWORDpassword"), []byte("saltSALTsaltSALTsaltSALTsaltSALTsalt"),
				4096, 25, sha256.New},
			want: []byte{0x34, 0x8c, 0x89, 0xdb, 0xcb, 0xd3, 0x2b, 0x2f, 0x32, 0xd8, 0x14, 0xb8, 0x11, 0x6e, 0x84,
				0xcf, 0x2b, 0x17, 0x34, 0x7e, 0xbc, 0x18, 0x00, 0x18, 0x1c},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := PBKDF2(tt.args.password, tt.args.salt, tt.args.iterCnt, tt.args.keyLen, tt.args.fn)
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("PBKDF2() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestTime33(t *testing.T) {
	type args struct {
		data []byte
	}
	tests := []struct {
		name string
		args args
		want uint32
	}{
		{
			name: "Time33",
			args: args{hashCommonTest},
			want: 2015120221,
		},
		{
			name: "Time33",
			args: args{[]byte("this is a test string")},
			want: 1917303537,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := Time33(tt.args.data); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("Time33() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestHashUInt64(t *testing.T) {
	type args struct {
		data []byte
		ht   HashType
	}
	tests := []struct {
		name string
		args args
		want uint64
	}{
		{
			name: "Fnv64",
			args: args{hashCommonTest, HtFnv64},
			want: 11132705866112348085,
		},
		{
			name: "Fnv64A",
			args: args{hashCommonTest, HtFnvA64},
			want: 8311336258473372757,
		},
		{
			name: "Crc64ISO",
			args: args{hashCommonTest, HtCrc64ISO},
			want: 6802869781201208823,
		},
		{
			name: "Crc64ECMA",
			args: args{hashCommonTest, HtCrc64ECMA},
			want: 11477354184825886377,
		},
		{
			name: "3Des",
			args: args{hashCommonTest, 0},
			want: 0,
		},
		{
			name: "Time33",
			args: args{hashCommonTest, HtTime33},
			want: 2015120221,
		},
		{
			name: "Fnv32",
			args: args{hashCommonTest, HtFnv32},
			want: 2305394613,
		},
		{
			name: "Fnv32A",
			args: args{hashCommonTest, HtFnvA32},
			want: 1285543765,
		},
		{
			name: "Crc32",
			args: args{hashCommonTest, HtCrc32},
			want: 1232993611,
		},
		{
			name: "Adler32",
			args: args{hashCommonTest, HtAdler32},
			want: 909837689,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := HashUInt64(tt.args.data, tt.args.ht); got != tt.want {
				t.Errorf("HashUInt64(%d) = %v, want %v", tt.args.ht, got, tt.want)
			}
		})
	}
}

func TestHashUInt32(t *testing.T) {
	type args struct {
		data []byte
		ht   HashType
	}
	tests := []struct {
		name string
		args args
		want uint32
	}{
		{
			name: "Fnv32",
			args: args{hashCommonTest, HtFnv32},
			want: 2305394613,
		},
		{
			name: "Fnv32A",
			args: args{hashCommonTest, HtFnvA32},
			want: 1285543765,
		},
		{
			name: "Crc32",
			args: args{hashCommonTest, HtCrc32},
			want: 1232993611,
		},
		{
			name: "Adler32",
			args: args{hashCommonTest, HtAdler32},
			want: 909837689,
		},
		{
			name: "Time33",
			args: args{hashCommonTest, HtTime33},
			want: 2015120221,
		},
		{
			name: "3Des",
			args: args{hashCommonTest, 0},
			want: 0,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := HashUInt32(tt.args.data, tt.args.ht); got != tt.want {
				t.Errorf("HashUInt32(%d) = %v, want %v", tt.args.ht, got, tt.want)
			}
		})
	}
}
