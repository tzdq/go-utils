package gzip

import (
	"bytes"
	"compress/gzip"
	"io/ioutil"
)

// ZipCompress gzip压缩
func ZipCompress(data []byte) ([]byte, error) {
	var buf bytes.Buffer

	writer := gzip.NewWriter(&buf)
	_, err := writer.Write(data)
	if err != nil {
		_ = writer.Close() //不能在defer中使用
		return nil, err
	}
	err = writer.Flush()
	if err != nil {
		_ = writer.Close()
		return nil, err
	}
	err = writer.Close()
	if err != nil {
		return nil, err
	}
	return buf.Bytes(), nil
}

// ZipDecompress gzip解压缩
func ZipDecompress(data []byte) ([]byte, error) {
	reader, err := gzip.NewReader(bytes.NewReader(data))
	if err != nil {
		return nil, err
	}
	defer func() {
		_ = reader.Close()
	}()
	return ioutil.ReadAll(reader)
}
