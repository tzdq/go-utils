package crypt

import (
	"crypto/aes"
	"crypto/cipher"
	"crypto/rand"
	"errors"
	"io"
)

// Advanced Encryption Standard (AES) a symmetric block cipher that can process data blocks of 128 bits,
// using cipher keys with lengths of 128, 192, and 256 bits. Using different block cipher working mode may involve
// the selection of initialization vector (IV) and filling mode.
// Block cipher working mode: CBC(Cipher-block chaining)、 ECB(Electronic codebook)、CFB(Cipher feedback)、
// OFB(Output feedback)、PCBC(Propagating cipher block chaining)、 CTR(Counter)
// In all mode except ECB, IV is required to randomize the encryption result. The length of IV must be
// equal to blockSize, otherwise it will panic
// The key argument should be the AES key, either 16, 24, or 32 bytes to select AES-128, AES-192, or AES-256
// In ECB and CBC mode, If the original plaintext lengths are not a multiple of the block size,padding would have to be
// added when encrypting,here we use PKCS7Padding, blockSize is 16. For more padding algorithms, see padding.go

// AESEncrypt Encrypts data with AES algorithm in specify mode.
// Recommended in combination with base64,such as: Base64Encode(AESEncrypt(plaintext,key,am))
func AESEncrypt(plaintext, key []byte, am AesMode) ([]byte, error) {
	switch am {
	case AesModeCBC:
		return aesCBCEncrypt(plaintext, key)
	case AesModeECB:
		return aesECBEncrypt(plaintext, key)
	case AesModeCFB, AesModeCTR, AesModeOFB:
		return aesStreamEncrypt(plaintext, key, am)
	default:
		return nil, errors.New(sErrAesModeInvalid)
	}
}

// AESDecrypt Decrypts cipher text with AES algorithm in specify mode
// Recommended in combination with base64,such as: AESDecrypt(Base64Decode(ciphertext),key,am))
func AESDecrypt(ciphertext, key []byte, am AesMode) ([]byte, error) {
	switch am {
	case AesModeCBC:
		return aesCBCDecrypt(ciphertext, key)
	case AesModeECB:
		return aesECBDecrypt(ciphertext, key)
	case AesModeCFB, AesModeCTR, AesModeOFB:
		return aesStreamDecrypt(ciphertext, key, am)
	default:
		return nil, errors.New(sErrAesModeInvalid)
	}
}

// aesCBCEncrypt Encrypts data with AES algorithm in CBC mode
func aesCBCEncrypt(plaintext, key []byte) ([]byte, error) {
	// The length of the key has been judged here, only supports 16、24、32
	block, err := aes.NewCipher(key)
	if err != nil {
		return nil, err
	}

	blockSize := block.BlockSize()
	paddingData := PKCS7Padding(plaintext, blockSize)
	// if padding data not full blocks,it will panic
	if len(paddingData)%blockSize != 0 {
		return nil, errors.New(sErrBlockNotFull)
	}

	encrypted := make([]byte, blockSize+len(paddingData))
	iv := encrypted[:blockSize]
	if _, err := io.ReadFull(rand.Reader, iv); err != nil {
		return nil, err
	}

	blockMode := cipher.NewCBCEncrypter(block, iv)
	blockMode.CryptBlocks(encrypted[blockSize:], paddingData)
	return encrypted, nil
}

// aesCBCDecrypt Decrypts cipher text with AES algorithm in CBC mode
func aesCBCDecrypt(ciphertext, key []byte) ([]byte, error) {
	block, err := aes.NewCipher(key)
	if err != nil {
		return nil, err
	}

	blockSize := block.BlockSize()
	if len(ciphertext) < blockSize {
		return nil, errors.New(sErrDataInvalid)
	}

	iv := ciphertext[:blockSize]
	ciphertext = ciphertext[blockSize:]
	// if padding data not full blocks,it will panic
	if len(ciphertext)%blockSize != 0 {
		return nil, errors.New(sErrBlockNotFull)
	}

	blockMode := cipher.NewCBCDecrypter(block, iv)
	blockMode.CryptBlocks(ciphertext, ciphertext)

	ciphertext, err = PKCS7UnPadding(ciphertext)
	if err != nil {
		return nil, err
	}
	return ciphertext, err
}

// aesECBEncrypt Encrypts data with AES algorithm in ECB mode
func aesECBEncrypt(plaintext, key []byte) ([]byte, error) {
	block, err := aes.NewCipher(key)
	if err != nil {
		return nil, err
	}

	blockSize := block.BlockSize()
	paddingData := PKCS7Padding(plaintext, blockSize)
	paddingDataLen := len(paddingData)

	// if padding data not full blocks,it will panic
	if paddingDataLen%blockSize != 0 {
		return nil, errors.New(sErrBlockNotFull)
	}

	encrypted := make([]byte, paddingDataLen)
	src := paddingData
	dst := encrypted

	for len(src) > 0 {
		block.Encrypt(dst, src[:blockSize])
		src = src[blockSize:]
		dst = dst[blockSize:]
	}
	return encrypted, nil
}

// aesECBDecrypt Decrypts cipher text using AES algorithm in ECB mode
func aesECBDecrypt(ciphertext, key []byte) ([]byte, error) {
	dataLen := len(ciphertext)
	if dataLen == 0 {
		return nil, errors.New(sErrDataInvalid)
	}

	block, err := aes.NewCipher(key)
	if err != nil {
		return nil, err
	}

	// if padding data not full blocks,it will panic
	blockSize := block.BlockSize()
	if dataLen%blockSize != 0 {
		return nil, errors.New(sErrBlockNotFull)
	}

	origData := make([]byte, dataLen)
	src := ciphertext
	dst := origData

	for len(src) > 0 {
		block.Decrypt(dst, src[:blockSize])
		src = src[blockSize:]
		dst = dst[blockSize:]
	}
	origData, err = PKCS7UnPadding(origData)
	if err != nil {
		return nil, err
	}
	return origData, err
}

// aesStreamEncrypt Encrypts data with AES algorithm in stream mode,include CTR、OFB、CFB
func aesStreamEncrypt(plaintext, key []byte, am AesMode) ([]byte, error) {
	block, err := aes.NewCipher(key)
	if err != nil {
		return nil, err
	}

	blockSize := block.BlockSize()
	encrypted := make([]byte, blockSize+len(plaintext))
	iv := encrypted[:blockSize]
	if _, err := io.ReadFull(rand.Reader, iv); err != nil {
		return nil, err
	}

	var stream cipher.Stream
	switch am {
	case AesModeCTR:
		stream = cipher.NewCTR(block, iv)
	case AesModeOFB:
		stream = cipher.NewOFB(block, iv)
	case AesModeCFB:
		stream = cipher.NewCFBEncrypter(block, iv)
	default:
		return nil, errors.New(sErrAesModeInvalid)
	}

	stream.XORKeyStream(encrypted[blockSize:], plaintext)
	return encrypted, nil
}

// aesStreamDecrypt Decrypts cipher text using AES algorithm in stream mode,include CTR、OFB、CFB
func aesStreamDecrypt(ciphertext, key []byte, am AesMode) ([]byte, error) {
	block, err := aes.NewCipher(key)
	if err != nil {
		return nil, err
	}

	blockSize := block.BlockSize()
	if len(ciphertext) < blockSize {
		return nil, errors.New(sErrDataInvalid)
	}

	iv := ciphertext[:blockSize]
	ciphertext = ciphertext[blockSize:]

	var stream cipher.Stream
	switch am {
	case AesModeCTR:
		stream = cipher.NewCTR(block, iv)
	case AesModeOFB:
		stream = cipher.NewOFB(block, iv)
	case AesModeCFB:
		stream = cipher.NewCFBDecrypter(block, iv)
	default:
		return nil, errors.New(sErrAesModeInvalid)
	}
	stream.XORKeyStream(ciphertext, ciphertext)
	return ciphertext, nil
}
