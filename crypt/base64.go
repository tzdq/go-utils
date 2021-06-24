package crypt

import (
	"encoding/base64"
)

// Base64Encode base64 encode
func Base64Encode(data []byte) string {
	return base64.StdEncoding.EncodeToString(data)
}

// Base64Decode base64 decode
func Base64Decode(str string) ([]byte, error) {
	return base64.StdEncoding.DecodeString(str)
}

// Base64UrlEncode url-safe base64 encode
// The special characters "+", "/" that are transcoded by the URL will not be generated
func Base64UrlEncode(data []byte) string {
	return base64.URLEncoding.EncodeToString(data)
}

// Base64UrlDecode url-safe base64 decode
func Base64UrlDecode(str string) ([]byte, error) {
	return base64.URLEncoding.DecodeString(str)
}
