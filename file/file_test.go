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
