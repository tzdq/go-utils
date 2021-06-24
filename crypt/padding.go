package crypt

import (
	"bytes"
	"errors"
)

// X923Padding padding the data through ANSI x9.23 padding Algorithm
// ANSI X9.23 : The block is padded with 0x0 and the last byte of the block is set to the number of bytes added.
// When data and blockSize are abnormal, it will panic, expect blockSize to be greater than 1, and data is not empty.
//
// example:
//  data := []byte{0x99, 0x98, 0x97, 0x96, 0x95, 0x94, 0x93, 0x92, 0x91, 0x90}
//  blockSize = 8
//
// The data length is 10,the block length is 8, and 6(8-10%8=6) bytes need to be filled, so the last byte is padded
// with 0x06, and the rest is padded with 0x00
// The final data is:{0x99, 0x98, 0x97, 0x96, 0x95, 0x94, 0x93, 0x92, 0x91, 0x90, 0x00, 0x00, 0x00, 0x00, 0x00, 0x06}
func X923Padding(data []byte, blockSize int) []byte {
	paddingCnt := blockSize - len(data)%blockSize
	padData := bytes.Repeat([]byte{0}, paddingCnt)
	padData[paddingCnt-1] = byte(paddingCnt)
	return append(data, padData...)
}

// X923UnPadding for the data padded by the ANSI X9.23 padding algorithm, get the original data
// example:
//  data := []byte{0x99, 0x98, 0x97, 0x96, 0x95, 0x94, 0x93, 0x92, 0x91, 0x90, 0x00, 0x00, 0x00, 0x00, 0x00, 0x06}
//
// According to the last byte, the number of padding bytes is 6, and the original byte array can be obtained by
// removing the last 6 bytes
// The final original data is:{0x99, 0x98, 0x97, 0x96, 0x95, 0x94, 0x93, 0x92, 0x91, 0x90}
func X923UnPadding(data []byte) ([]byte, error) {
	dataLen := len(data)
	if dataLen == 0 {
		return nil, errors.New(sErrDataEmpty)
	}
	paddingCnt := uint(data[dataLen-1])
	if paddingCnt > uint(dataLen) {
		return nil, errors.New(sErrDataLenInvalid)
	}
	for i := paddingCnt; i > 1; i-- {
		if data[uint(dataLen)-i] != 0 {
			return nil, errors.New(sErrDataInvalid)
		}
	}
	return data[:(uint(dataLen) - paddingCnt)], nil
}

// ISO10126Padding padding the data through ISO 10126 padding algorithm
// ISO 10126 : The block is padded with random bytes and the last byte of the block is set to the number of bytes added.
// When data and blockSize are abnormal, it will panic, expect blockSize to be greater than 1, and data is not empty.
//
// example:
//  data := []byte{0x99, 0x98, 0x97, 0x96, 0x95, 0x94, 0x93, 0x92, 0x91, 0x90}
//  blockSize = 8
//
// The data length is 10,the block length is 8, and 6(8-10%8=6) bytes need to be filled, so the last byte is padded
// with 0x06, and the rest is padded with random bytes
// The final data is:{0x99, 0x98, 0x97, 0x96, 0x95, 0x94, 0x93, 0x92, 0x91, 0x90, 0x34, 0x8A, 0xEF, 0xF3, 0x00, 0x06}
func ISO10126Padding(data []byte, blockSize int) []byte {
	paddingCnt := blockSize - len(data)%blockSize
	padData := RandBytes(uint32(paddingCnt))
	padData[paddingCnt-1] = byte(paddingCnt)
	return append(data, padData...)
}

// ISO10126UnPadding for the data padded by the ISO 10126 padding algorithm, get the original data
// example:
//  data := []byte{0x99, 0x98, 0x97, 0x96, 0x95, 0x94, 0x93, 0x92, 0x91, 0x90, 0x34, 0x8A, 0xEF, 0xF3, 0x00, 0x06}
//
// According to the last byte, the number of padding bytes is 6, and the original byte array can be obtained by
// removing the last 6 bytes
// The final original data is:{0x99, 0x98, 0x97, 0x96, 0x95, 0x94, 0x93, 0x92, 0x91, 0x90}
func ISO10126UnPadding(data []byte) ([]byte, error) {
	dataLen := len(data)
	if dataLen == 0 {
		return nil, errors.New(sErrDataEmpty)
	}
	paddingCnt := uint(data[dataLen-1])
	if paddingCnt > uint(dataLen) {
		return nil, errors.New(sErrDataLenInvalid)
	}
	return data[:(uint(dataLen) - paddingCnt)], nil
}

// ISO7816Padding padding the data through ISO 7816-4 padding Algorithm
// ISO 7816-4 : The first byte is padded with 0x80, and the remaining bytes are padded with 0.
// When data and blockSize are abnormal, it will panic, expect blockSize to be greater than 1, and data is not empty.
//
// example:
//  data := []byte{0x99, 0x98, 0x97, 0x96, 0x95, 0x94, 0x93, 0x92, 0x91, 0x90}
//  blockSize = 8
//
// The data length is 10,the block length is 8, and 6(8-10%8=6) bytes need to be filled, so the first byte is padded
// with 0x80, and the rest is padded with 0x0
// The final data is:{0x99, 0x98, 0x97, 0x96, 0x95, 0x94, 0x93, 0x92, 0x91, 0x90, 0x80, 0x00, 0x00, 0x00, 0x00, 0x00}
func ISO7816Padding(data []byte, blockSize int) []byte {
	paddingCnt := blockSize - len(data)%blockSize
	padData := bytes.Repeat([]byte{0}, paddingCnt)
	padData[0] = 0x80
	return append(data, padData...)
}

// ISO7816UnPadding for the data padded by the ISO 7816-4 padding algorithm, get the original data
// example:
//  data := []byte{0x99, 0x98, 0x97, 0x96, 0x95, 0x94, 0x93, 0x92, 0x91, 0x90, 0x80, 0x00, 0x00, 0x00, 0x00, 0x00}
//
// Traverse from back to front, search for subscripts with consecutive 0x00 values, If the previous digit of this
// subscript is 0x80, then remove the data starting from 0x80 to get the original data.
// The final original data is:{0x99, 0x98, 0x97, 0x96, 0x95, 0x94, 0x93, 0x92, 0x91, 0x90}
func ISO7816UnPadding(data []byte) ([]byte, error) {
	dataLen := len(data)
	if dataLen == 0 {
		return nil, errors.New(sErrDataEmpty)
	}
	paddingCnt := dataLen - 1
	for data[paddingCnt] == 0 {
		paddingCnt = paddingCnt - 1
	}
	if data[paddingCnt] != 0x80 {
		return nil, errors.New(sErrDataInvalid)
	}
	return data[:paddingCnt], nil
}

// PKCS7Padding padding the data through PKCS7 padding Algorithm
// PKCS7 : The value of each added byte is the number of bytes that are added
// When data and blockSize are abnormal, it will panic, expect blockSize to be greater than 1, and data is not empty.
//
// example:
//  data := []byte{0x99, 0x98, 0x97, 0x96, 0x95, 0x94, 0x93, 0x92, 0x91, 0x90}
//  blockSize = 8
//
// The data length is 10,the block length is 8, and 6(8-10%8=6) bytes need to be filled, so padded all the stuffing
// bytes with 0x06
// The final data is:{0x99, 0x98, 0x97, 0x96, 0x95, 0x94, 0x93, 0x92, 0x91, 0x90, 0x06, 0x06, 0x06, 0x06, 0x06, 0x06}
func PKCS7Padding(data []byte, blockSize int) []byte {
	paddingCnt := blockSize - len(data)%blockSize
	padData := bytes.Repeat([]byte{byte(paddingCnt)}, paddingCnt)
	return append(data, padData...)
}

// PKCS7UnPadding for the data padded by the PKCS7 padding algorithm, get the original data
// example:
//  data := []byte{0x99, 0x98, 0x97, 0x96, 0x95, 0x94, 0x93, 0x92, 0x91, 0x90, 0x06, 0x06, 0x06, 0x06, 0x06, 0x06}
//
// According to the last byte, the number of padding bytes is 6, and the original byte array can be obtained by
// removing the last 6 bytes
// The final original data is:{0x99, 0x98, 0x97, 0x96, 0x95, 0x94, 0x93, 0x92, 0x91, 0x90}
func PKCS7UnPadding(data []byte) ([]byte, error) {
	dataLen := len(data)
	if dataLen == 0 {
		return nil, errors.New(sErrDataEmpty)
	}
	paddingCnt := uint(data[dataLen-1])
	if paddingCnt > uint(dataLen) {
		return nil, errors.New(sErrDataLenInvalid)
	}

	for i := 1; i <= int(paddingCnt); i++ {
		if data[dataLen-i] != byte(paddingCnt) {
			return nil, errors.New(sErrDataInvalid)
		}
	}
	return data[:(uint(dataLen) - paddingCnt)], nil
}

// ZeroPadding padding the data through Zero padding Algorithm
// Zero : The block is padded with 0x0. If there is data ending with 0x0 in the original data, there will be problems
// during unpading. This method is not recommended
// When data and blockSize are abnormal, it will panic, expect blockSize to be greater than 1, and data is not empty.
//
// example:
//  data := []byte{0x99, 0x98, 0x97, 0x96, 0x95, 0x94, 0x93, 0x92, 0x91, 0x90}
//  blockSize = 8
//
// The data length is 10,the block length is 8, and 6(8-10%8=6) bytes need to be filled, sp padded all the stuffing
// bytes with 0x0
// The final data is:{0x99, 0x98, 0x97, 0x96, 0x95, 0x94, 0x93, 0x92, 0x91, 0x90, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00}
func ZeroPadding(data []byte, blockSize int) []byte {
	paddingCnt := blockSize - len(data)%blockSize
	padData := bytes.Repeat([]byte{0}, paddingCnt)
	return append(data, padData...)
}

// ZeroUnPadding for the data padded by the Zero padding algorithm, get the original data
// example:
//  data := []byte{0x99, 0x98, 0x97, 0x96, 0x95, 0x94, 0x93, 0x92, 0x91, 0x90, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00}
//
// Traverse from back to front, search for subscripts with a continuous value of 0x00, and remove the data starting
// from this subscript to get the original data. Note that if the original data contains data ending in 0x00, there
// may be problems with the original data obtained at this time. This method is not recommended
// The final original data is:{0x99, 0x98, 0x97, 0x96, 0x95, 0x94, 0x93, 0x92, 0x91, 0x90}
func ZeroUnPadding(data []byte) ([]byte, error) {
	dataLen := len(data)
	if dataLen == 0 {
		return nil, errors.New(sErrDataEmpty)
	}
	for data[dataLen-1] == 0 {
		dataLen = dataLen - 1
	}
	return data[:dataLen], nil
}
