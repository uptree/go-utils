package file

import (
	"bufio"
	"bytes"
	"io"
	"io/ioutil"
	"net/http"
	"os"
	"path"
)

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
	_ = MkDir(path.Dir(filePath))
	out, _ := os.Create(filePath)
	defer out.Close()
	wt := bufio.NewWriter(out)
	_, err := io.Copy(wt, bytes.NewReader(b))
	wt.Flush()
	return err
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
