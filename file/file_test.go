package file

import (
	"testing"
)

func TestCreateFile(t *testing.T) {
	filePath := "testdata/file/create.txt"
	fs, err := CreateFile(filePath)
	defer fs.Close()
	t.Logf("CreateFile: %v", err)
}

func TestWriteBytesToFile(t *testing.T) {
	filePath := "testdata/file/write.txt"
	t.Logf("WriteBytesToFile: %v", WriteBytesToFile(filePath, []byte("hello")))
}

func TestReadFileToBytes(t *testing.T) {
	filePath := "testdata/file/write.txt"
	b, err := ReadFileToBytes(filePath)
	t.Logf("ReadFileToBytes: %v, %v", string(b), err)
}

func TestReadHttpFileToBytes(t *testing.T) {
	filePath := "https://game.gtimg.cn/images/yxzj/img201606/heroimg/109/109.jpg"
	b, err := ReadHttpFileToBytes(filePath)
	err = WriteBytesToFile("testdata/file/httpFile.jpeg", b)
	t.Logf("ReadHttpFileToBytes:  %v", err)
}

func TestCopy(t *testing.T) {
	sourcePath := "testdata/file/write.txt"
	targetPath := "testdata/file/target.txt"
	t.Logf("Copy: %v", Copy(sourcePath, targetPath))
}

func TestDownload(t *testing.T) {
	sourceUrl := "https://game.gtimg.cn/images/yxzj/img201606/heroimg/109/109.jpg"
	targetPath := "testdata/file/109-download.jpeg"
	t.Logf("Copy: %v", Download(sourceUrl, targetPath))
}

func TestName(t *testing.T) {
	t.Logf("Name: %v", Name("testdata/file/109-download.jpeg"))
}
