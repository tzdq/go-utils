package crypt

import (
	"crypto/hmac"
	"crypto/md5"
	"crypto/sha1"
	"crypto/sha256"
	"crypto/sha512"
	"encoding/hex"
	"hash"
	"hash/adler32"
	"hash/crc32"
	"hash/crc64"
	"hash/fnv"
	"io"
	"log"
	"os"
)

// Md5File md5 the content of the file and return the result as a hexadecimal string
// You can also use shaX series functions instead of md5,such as sha1....
func Md5File(path string) string {
	f, err := os.Open(path)
	if err != nil {
		log.Fatal(err)
	}
	defer f.Close()

	h := md5.New()
	if _, err := io.Copy(h, f); err != nil {
		log.Fatal(err)
	}
	return hex.EncodeToString(h.Sum(nil))
}

// HashBytes return the checksum raw buffer of the specified hash algorithm
// ht: md5(16bytes) 、sha1(20bytes)、sha224(28bytes)、sha256(32bytes)、sha384(48bytes)、sha512(64bytes)
func HashBytes(data []byte, ht HashType) []byte {
	var h hash.Hash
	switch ht {
	case HtMD5:
		h = md5.New()
	case HtSha1:
		h = sha1.New()
	case HtSha224:
		h = sha256.New224()
	case HtSha256:
		h = sha256.New()
	case HtSha384:
		h = sha512.New384()
	case HtSha512:
		h = sha512.New()
	case HtFnv32:
		h = fnv.New32()
	case HtFnvA32:
		h = fnv.New32a()
	case HtFnv64:
		h = fnv.New64()
	case HtFnvA64:
		h = fnv.New64a()
	case HtFnv128:
		h = fnv.New128()
	case HtFnvA128:
		h = fnv.New128a()
	case HtCrc32:
		h = crc32.NewIEEE()
	case HtAdler32:
		h = adler32.New()
	case HtCrc64ISO:
		h = crc64.New(crc64.MakeTable(crc64.ISO))
	case HtCrc64ECMA:
		h = crc64.New(crc64.MakeTable(crc64.ECMA))
	case HtTime33:
		return toBytes(Time33(data))
	default:
		return []byte{}
	}
	h.Write(data)
	return h.Sum(nil)
}

// HmacBytes return the authentication code raw buffer of the specified hash algorithm.
// ht only support HtMD5、HtSha1、HtSha224、HtSha384、HtSha512
func HmacBytes(data, key []byte, ht HashType) []byte {
	var h hash.Hash

	switch ht {
	case HtMD5:
		h = hmac.New(md5.New, key)
	case HtSha1:
		h = hmac.New(sha1.New, key)
	case HtSha224:
		h = hmac.New(sha256.New224, key)
	case HtSha256:
		h = hmac.New(sha256.New, key)
	case HtSha384:
		h = hmac.New(sha512.New384, key)
	case HtSha512:
		h = hmac.New(sha512.New, key)
	default:
		return []byte{}
	}
	h.Write(data)
	return h.Sum(nil)
}

// ToHexString convert bytes to hexadecimal string
func ToHexString(src []byte) string {
	return hex.EncodeToString(src)
}

// PBKDF2 return the encrypted passwords
// PBKDF2(Password-Based Key Derivation Function) is a function used to generate keys and is often used to generate
// encrypted passwords. Its basic principle is to pass a pseudo-random function (such as the HMAC function),
// take plaintext and a salt value as input parameters, and then repeat the operation, and finally generate a key.
//
// password: The original password used to generate the key
// salt    : Salt value for encryption,its recommended to use random numbers
// iterCnt : The number of iterations, the more the number, the longer the time required for encryption and decryption
// keyLen  : The desired bit-length of the derived key
// fn      : The hash function used for encryption, sha1 is used by default
func PBKDF2(password, salt []byte, iterCnt, keyLen uint32, fn func() hash.Hash) []byte {
	if fn == nil {
		fn = sha1.New
	}
	prf := hmac.New(fn, password)
	hashLen := prf.Size()
	numBlocks := (int(keyLen) + hashLen - 1) / hashLen

	var buf [4]byte
	dk := make([]byte, 0, numBlocks*hashLen)
	U := make([]byte, hashLen)
	for block := 1; block <= numBlocks; block++ {
		// || means concatenation, ^ means XOR
		// for each block T_i = U_1 ^ U_2 ^ ... ^ U_iter
		// U_1 = PRF(password, salt || uint(i))
		prf.Reset()
		prf.Write(salt)
		buf[0] = byte(block >> 24)
		buf[1] = byte(block >> 16)
		buf[2] = byte(block >> 8)
		buf[3] = byte(block)
		prf.Write(buf[:4])
		dk = prf.Sum(dk)
		T := dk[len(dk)-hashLen:]
		copy(U, T)

		// U_n = PRF(password, U_(n-1))
		for n := uint32(2); n <= iterCnt; n++ {
			prf.Reset()
			prf.Write(U)
			U = U[:0]
			U = prf.Sum(U)
			for x := range U {
				T[x] ^= U[x]
			}
		}
	}
	return dk[:keyLen]
}

// Time33 return the hash value of the plaintext through the Time33 hash function
// Time33(Daniel J. Bernstein, Times 33 with Addition) ,Its basic principle is to iteratively multiply each character
// of the string by 33.
// hash(i) = hash(i-1) * 33 + str[i]
func Time33(data []byte) uint32 {
	length := len(data)
	h := uint32(5381)

	for i := 0; i < length; i++ {
		h += ((h << 5) & 0x7FFFFFFF) + uint32(data[i]) // & 0x7FFFFFFF ensure that its value is in the int32 range
	}
	return h & 0x7FFFFFFF
}

// HashUInt32 return a hash value of uint32 type through a specific hash function
// ht only support HtFnv32、HtFnvA32、 HtAdler32、HtCrc32、HtTime33
func HashUInt32(data []byte, ht HashType) uint32 {
	switch ht {
	case HtFnv32:
		h := fnv.New32()
		h.Write(data)
		return h.Sum32()
	case HtFnvA32:
		h := fnv.New32a()
		h.Write(data)
		return h.Sum32()
	case HtAdler32:
		return adler32.Checksum(data)
	case HtCrc32:
		return crc32.ChecksumIEEE(data)
	case HtTime33:
		return Time33(data)
	default:
		return 0
	}
}

// HashUInt64 return a hash value of uint64 type through a specific hash function
// ht only support HtFnv32、HtFnvA32、 HtAdler32、HtCrc32、HtTime33、HtFnv64、HtFnvA64、HtCrc64ISO、HtCrc64ECMA
func HashUInt64(data []byte, ht HashType) uint64 {
	switch ht {
	case HtFnv32, HtFnvA32, HtAdler32, HtCrc32, HtTime33:
		return uint64(HashUInt32(data, ht))
	case HtFnv64:
		h := fnv.New64()
		h.Write(data)
		return h.Sum64()
	case HtFnvA64:
		h := fnv.New64a()
		h.Write(data)
		return h.Sum64()
	case HtCrc64ISO:
		crc64t := crc64.MakeTable(crc64.ISO)
		return crc64.Checksum(data, crc64t)
	case HtCrc64ECMA:
		crc64t := crc64.MakeTable(crc64.ECMA)
		return crc64.Checksum(data, crc64t)
	default:
		return 0
	}
}

// uint32 convert uint32 to byte slice
func toBytes(i uint32) []byte {
	var in []byte
	return append(in, byte(i>>24), byte(i>>16), byte(i>>8), byte(i))
}

// JumpConsistentHash return a hash value in the range of [0,numBuckets) through jump consistent hash algorithm
// If numBuckets is less than or equal to 0,the default value of 1 is used
func JumpConsistentHash(key uint64, numBuckets int32) int32 {
	if numBuckets <= 0 {
		numBuckets = 1
	}
	b, j := int64(-1), int64(0)
	for j < int64(numBuckets) {
		b = j
		key = key*2862933555777941757 + 1
		j = int64(float64(b+1) * (float64(int64(1)<<31) / float64((key>>33)+1)))
	}

	return int32(b)
}
