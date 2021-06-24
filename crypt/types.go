package crypt

// error string
const (
	sErrDataInvalid    = "data is invalid"
	sErrDataEmpty      = "data is empty"
	sErrDataLenInvalid = "data padding len is invalid"
	sErrBlockNotFull   = "input not full blocks"
	sErrAesModeInvalid = "aes work mode invalid"
	sErrPublicKeyErr   = "public key error"
	sErrPrivateKeyErr  = "private key error"
)

// -------------------------------------------------------------------------------------

// HashType Hash algorithm type. use in hash.go
type HashType int

const (
	HtMD5 HashType = iota + 1
	HtSha1
	HtSha224
	HtSha256
	HtSha384
	HtSha512
	HtFnv32
	HtFnvA32
	HtFnv64
	HtFnvA64
	HtFnv128
	HtFnvA128
	HtTime33
	HtAdler32
	HtCrc32
	HtCrc64ISO
	HtCrc64ECMA
)

// -------------------------------------------------------------------------------------

// ScopeType Source type definition for generating random string. use in random.go
type ScopeType int

const (
	OnlyLowerCase       ScopeType = iota + 1 // contain only lowercase letters
	OnlyUpperCase                            // contain only uppercase letters
	OnlyNumber                               // contains only numbers
	AlphaLetter                              // uppercase and lowercase letters
	LowerCaseAndNumber                       // lowercase letters and numbers
	LowerCaseAndSpecial                      // lowercase letters and special characters
	UpperCaseAndNumber                       // uppercase letters and numbers
	UpperCaseAndSpecial                      // uppercase letters and special characters
	NumberAndSpecial                         // numbers and special characters
	AlphaAndNumber                           // uppercase and lowercase letters and numbers
	AlphaAndSpecial                          // uppercase and lowercase letters and special characters
	NoUpperCase                              // lowercase letters and special characters and numbers
	NoLowerCase                              // uppercase letters and special characters and numbers
	AllLetter                                // uppercase and lowercase letters and special characters and numbers
)

const (
	lowerAlphaCharset = "abcdefghijklmnopqrstuvwxyz"
	upperAlphaCharset = "ABCDEFGHIJKLMNOPQRSTUVWXYZ"
	numberCharset     = "0123456789"
	specialCharset    = "~!@#$%^&*()_,.<>{}=-+"
)

// -------------------------------------------------------------------------------------

// AesMode AES encryption algorithm working mode. use in aes.go
type AesMode int32

const (
	AesModeCBC AesMode = iota + 1 // default mode
	AesModeCFB
	AesModeECB
	AesModeCTR
	AesModeOFB
)
