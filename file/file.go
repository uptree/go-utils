package file

import (
	"bufio"
	"bytes"
	"io"
	"io/ioutil"
	"net/http"
	"os"
	"path"
	"path/filepath"
)

/**
os.O_WRONLY	只写
os.O_CREATE	创建文件
os.O_RDONLY	只读
os.O_RDWR	读写
os.O_TRUNC	清空
os.O_APPEND	追加
*/

// CreateFile 创建或清空文件
func CreateFile(filePath string) (*os.File, error) {
	if err := MkDir(path.Dir(filePath)); err != nil {
		return nil, err
	}
	return os.Create(filePath)
}

// ReadFileToBytes 读文件
func ReadFileToBytes(filePath string) ([]byte, error) {
	b, err := ioutil.ReadFile(filePath)
	if err != nil {
		return nil, err
	}
	return b, nil
}

// WriteBytesToFile 写文件 覆盖
func WriteBytesToFile(filePath string, b []byte) error {
	f, err := CreateFile(filePath)
	if err != nil {
		return err
	}
	defer f.Close()
	wt := bufio.NewWriter(f)
	_, err = io.Copy(wt, bytes.NewReader(b))
	if err != nil {
		return err
	}
	wt.Flush()
	return nil
}

// ReadHttpFileToBytes 读网络文件
func ReadHttpFileToBytes(fileUrl string) ([]byte, error) {
	resp, err := http.Get(fileUrl)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()
	return ioutil.ReadAll(resp.Body)
}

// Copy 复制文件
func Copy(sourcePath, targetPath string) error {
	b, err := ReadFileToBytes(sourcePath)
	if err != nil {
		return nil
	}
	return WriteBytesToFile(targetPath, b)
}

// Download 下载文件
func Download(sourceUrl, targetPath string) error {
	b, err := ReadHttpFileToBytes(sourceUrl)
	if err != nil {
		return nil
	}
	return WriteBytesToFile(targetPath, b)
}

// Name 获取文件名.
func Name(filePath string) string {
	return filepath.Base(filePath)
}
