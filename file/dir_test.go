package file

import (
	"testing"
)

func TestSelfPath(t *testing.T) {
	t.Logf("SelfPath: %v", SelfPath())
}

func TestSelfDir(t *testing.T) {
	t.Logf("SelfDir: %v", SelfDir())
}

func TestMkDir(t *testing.T) {
	dirPath := "testdata/dir/dir1"
	t.Logf("Mkdir: %v", MkDir(dirPath))
}

func TestIsEmpty(t *testing.T) {
	dirPath := "testdata/dir"
	t.Logf("IsEmpty: %v", IsEmpty(dirPath))
}

func TestIsDir(t *testing.T) {
	dirPath := "testdata/dir"
	t.Logf("IsDir: %v", IsDir(dirPath))
}

func TestIsFile(t *testing.T) {
	dirPath := "testdata/dir"
	t.Logf("IsFile: %v", IsFile(dirPath))
}

func TestListIndex(t *testing.T) {
	dirPath := "testdata"
	files, dirs, err := ListIndex(dirPath)
	t.Logf("ListFiles: %v, %v, %v", files, dirs, err)
}

func TestClearDir(t *testing.T) {
	dirPath := "testdata/dir/dir1"
	t.Logf("ClearDir: %v", ClearDir(dirPath))
}

func TestClearDirF(t *testing.T) {
	dirPath := "testdata"
	t.Logf("ClearDirF: %v", ClearDirF(dirPath))
}

func TestRemoveDir(t *testing.T) {
	dirPath := "testdata"
	t.Logf("RemoveDir: %v", RemoveDir(dirPath))
}
