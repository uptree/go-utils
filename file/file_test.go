package file

import (
	"testing"
)

func TestWriteBytesToFile(t *testing.T) {
	path2 := "dir1/dir2/1.txt"
	t.Logf("WriteBytesToFile: %v", WriteBytesToFile(path2, []byte("hello")))
}

func TestReadFileToBytes(t *testing.T) {
	path2 := "dir1/dir2/1.txt"
	b, err := ReadFileToBytes(path2)
	t.Logf("ReadFileToBytes: %v, %v", string(b), err)
}

func TestReadHttpFileToBytes(t *testing.T) {
	path2 := "https://game.gtimg.cn/images/yxzj/img201606/heroimg/109/109.jpg"
	b, err := ReadHttpFileToBytes(path2)
	WriteBytesToFile("dir1/dir2/109.jpeg", b)
	t.Logf("ReadHttpFileToBytes: %v, %v", string(b), err)
}

func TestCopy(t *testing.T) {
	sourcePath := "dir1/dir2/1.txt"
	targetPath := "dir1/dir2/2.txt"
	t.Logf("Copy: %v", Copy(sourcePath, targetPath))
}

func TestDownload(t *testing.T) {
	sourceUrl := "https://game.gtimg.cn/images/yxzj/img201606/heroimg/109/109.jpg"
	targetPath := "dir1/dir2/109-download.jpeg"
	t.Logf("Copy: %v", Download(sourceUrl, targetPath))
}
