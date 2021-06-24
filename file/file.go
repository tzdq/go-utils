package file

import (
	"bufio"
	"errors"
	"fmt"
	"io"
	"io/ioutil"
	"os"
	"path/filepath"
)

// 内核中的缓冲：无论进程是否提供缓冲，内核都是提供缓冲的，系统对磁盘的读写都会提供一个缓冲（内核高速缓冲），将数据写入到块缓冲进行排队，
//             当块缓冲达到一定的量时，才把数据写入磁盘。
// 进程中的缓冲：是指对输入输出流进行了改进，提供了一个流缓冲，当调用一个函数向磁盘写数据时，先把数据写入缓冲区，当达到某个条件，如流缓冲满了，
//             或刷新流缓冲，这时候才会把数据一次送往内核提供的块缓冲中，再经块化重写入磁盘。
// 这里的io和bufio是进程中的缓冲。

// CheckAndCreate 检查目录是否存在，如果不存在，创建目录
func CheckAndCreate(dir string) error {
	_, err := os.Stat(dir)
	if err == nil {
		return nil
	}
	if !os.IsNotExist(err) {
		return fmt.Errorf("can't stat %q:%s", dir, err)
	}
	if err := os.MkdirAll(dir, 0755); err != nil {
		return fmt.Errorf("can't mkdir %q:%s", dir, err)
	}
	return nil
}

// IsExist 检查文件是否存在
func IsExist(fileName string) bool {
	if _, err := os.Stat(fileName); os.IsNotExist(err) {
		return false
	}
	return true
}

// IsFile 判断是目录还是文件
func IsFile(fileName string) bool {
	fi, err := os.Stat(fileName)
	if err != nil {
		return false
	}
	return !fi.IsDir()
}

// CreateFile 创建文件,如果文件所处的目录不存在，会创建目录
func CreateFile(pathName string) (*os.File, error) {
	path, fileName := filepath.Split(pathName)
	if len(fileName) == 0 {
		return nil, errors.New("file name is empty")
	}
	if len(path) != 0 && path != "./" {
		err := os.MkdirAll(path, os.ModePerm)
		if err != nil {
			return nil, err
		}
	}

	return os.Create(pathName)
}

// ---------------------------------------------------------------
// 从benchmark的结果对比,具体结果保存在testdata/bench.txt
// 读取速度 ReadFile > ReadByBytes > ReadAll
// 内存分配 ReadFile > ReadAll > ReadByBytes
// 对象分配 ReadFile > ReadByBytes = ReadAll

// ReadAll 从文件中读取全部内容
func ReadAll(fileName string) ([]byte, error) {
	fp, err := os.Open(fileName)
	if err != nil {
		return nil, err
	}
	defer func() {
		_ = fp.Close()
	}()

	return ioutil.ReadAll(fp)
}

// ReadFile 从文件中读取全部内容,推荐使用该函数
func ReadFile(fileName string) ([]byte, error) {
	return ioutil.ReadFile(fileName)
}

// ReadByBytes 按照字节读取文件，如果bits为0，默认按照1024字节
func ReadByBytes(fileName string, bits int) ([]byte, error) {
	if bits <= 0 {
		bits = 1024
	}
	fp, err := os.Open(fileName)
	if err != nil {
		return nil, err
	}
	defer func() {
		_ = fp.Close()
	}()

	reader := bufio.NewReader(fp)
	chunks := make([]byte, 0)
	buf := make([]byte, bits) //一次性读取bits个字节
	for {
		n, err := reader.Read(buf)
		if err != nil && err != io.EOF {
			return nil, err
		}
		if 0 == n {
			break
		}
		chunks = append(chunks, buf[:n]...)
	}
	return chunks, nil
}

// -----------------------------------------------------------------------------------------
// 从benchmark的结果对比,具体结果保存在testdata/bench.txt
// 写入速度 WriteFile4BufIO = WriteFile4OS > WriteFile4IOUtil > WriteFile
// 内存分配 WriteFile4IOUtil > WriteFile4OS = WriteFile > WriteFile4BufIO
// 对象分配 WriteFile4IOUtil > WriteFile4OS = WriteFile = WriteFile4BufIO

// WriteFile4IOUtil 写入文件，默认清空文件内容后追加，如果文件不存在会创建文件
func WriteFile4IOUtil(fileName string, data []byte) error {
	return ioutil.WriteFile(fileName, data, 0644)
}

// WriteFile4OS 写入文件，默认清空文件内容后追加，如果文件不存在会创建文件
func WriteFile4OS(fileName string, data string) error {
	var fp *os.File
	var err error
	if IsExist(fileName) {
		fp, err = os.OpenFile(fileName, os.O_APPEND, 0666)
	} else {
		fp, err = os.Create(fileName)
	}
	if err != nil {
		return err
	}
	defer func() {
		_ = fp.Close()
	}()

	_, err = io.WriteString(fp, data)
	return err
}

// WriteFile 写入文件，默认清空文件内容后追加，如果文件不存在会创建文件
func WriteFile(fileName string, data string) error {
	var fp *os.File
	var err error
	if IsExist(fileName) {
		fp, err = os.OpenFile(fileName, os.O_APPEND, 0666)
	} else {
		fp, err = os.Create(fileName)
	}
	if err != nil {
		return err
	}
	defer func() {
		_ = fp.Close()
	}()

	_, err = fp.WriteString(data)
	if err != nil {
		return err
	}
	return fp.Sync()
}

// WriteFile4BufIO 写入文件，默认清空文件内容后追加，如果文件不存在会创建文件
func WriteFile4BufIO(fileName string, data string) error {
	var fp *os.File
	var err error
	if IsExist(fileName) {
		fp, err = os.OpenFile(fileName, os.O_APPEND, 0666)
	} else {
		fp, err = os.Create(fileName)
	}
	if err != nil {
		return err
	}
	defer func() {
		_ = fp.Close()
	}()

	writer := bufio.NewWriter(fp)
	_, err = writer.WriteString(data)
	if err != nil {
		return err
	}
	return writer.Flush()
}
