package file

import (
	"bufio"
	"bytes"
	"io"
	"io/ioutil"
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

// WriteBytesToFile 写文件
func WriteBytesToFile(filePath string, b []byte) error {
	_ = MkDir(path.Dir(filePath))
	out, _ := os.Create(filePath)
	defer out.Close()
	wt := bufio.NewWriter(out)
	_, err := io.Copy(wt, bytes.NewReader(b))
	wt.Flush()
	return err
}
