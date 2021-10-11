package crypt

import (
	"math/rand"
	"time"
)

func init() {
	rand.Seed(time.Now().UnixNano())
}

// RandInt return a non_negative random integer,range:[0, ∞)
func RandInt() int {
	return rand.Int()
}

// RandIntN return a non_negative random integer,range:[0, n-1]
func RandIntN(n int) int {
	if n <= 0 {
		return 0
	}
	return rand.Intn(n)
}

// RandIntRange return a non_negative random integer,range:[min, max-1]
// min and max are not necessarily in ascending order
func RandIntRange(min, max int) int {
	tmpMin, tmpMax := min, max
	if max < min {
		tmpMin, tmpMax = max, min
	}

	if tmpMax <= 0 {
		return 0
	} else if tmpMin < 0 {
		tmpMin = 0
	}

	return RandIntN(tmpMax-tmpMin) + tmpMin
}

// RandFloats return a non_negative random float,range:[min, max)
// min and max are not necessarily in ascending order
func RandFloats(min, max float64) float64 {
	tmpMin, tmpMax := min, max
	if max < min {
		tmpMin, tmpMax = max, min
	}
	if tmpMax <= 0 {
		return 0
	} else if tmpMin < 0 {
		tmpMin = 0
	}
	return tmpMin + rand.Float64()*(tmpMax-tmpMin)
}

// RandBytes return a random byte slice of the specified length, each byte has a value range of [0x00,0xff]
func RandBytes(length uint32) []byte {
	b := make([]byte, length)
	for i := range b {
		b[i] = byte(RandInt() % 256)
	}
	return b
}

// RandString return a random string of specified length
// The content of the string consists of uppercase and lowercase letters, numbers, and special characters
func RandString(length int, st ScopeType) string {
	if length <= 0 {
		return ""
	}

	var charset string
	switch st {
	case OnlyLowerCase:
		charset = lowerAlphaCharset
	case OnlyUpperCase:
		charset = upperAlphaCharset
	case OnlyNumber:
		charset = numberCharset
	case AlphaLetter:
		charset = lowerAlphaCharset + upperAlphaCharset
	case LowerCaseAndNumber:
		charset = lowerAlphaCharset + numberCharset
	case LowerCaseAndSpecial:
		charset = lowerAlphaCharset + specialCharset
	case UpperCaseAndNumber:
		charset = upperAlphaCharset + numberCharset
	case UpperCaseAndSpecial:
		charset = upperAlphaCharset + specialCharset
	case NumberAndSpecial:
		charset = numberCharset + specialCharset
	case AlphaAndNumber:
		charset = lowerAlphaCharset + upperAlphaCharset + numberCharset
	case AlphaAndSpecial:
		charset = lowerAlphaCharset + upperAlphaCharset + specialCharset
	case NoUpperCase:
		charset = lowerAlphaCharset + numberCharset + specialCharset
	case NoLowerCase:
		charset = upperAlphaCharset + numberCharset + specialCharset
	case AllLetter:
		charset = lowerAlphaCharset + upperAlphaCharset + numberCharset + specialCharset
	default:
		charset = lowerAlphaCharset + upperAlphaCharset
	}
	charsetLen := len(charset)

	b := make([]byte, length)
	for i := range b {
		b[i] = charset[RandIntN(charsetLen)]
	}
	return string(b)
}
