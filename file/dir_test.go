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
	path2 := "dir1/dir2/dir3"
	t.Logf("Mkdir: %v", MkDir(path2))
}

func TestIsEmpty(t *testing.T) {
	path2 := "dir1/dir2"
	t.Logf("IsEmpty: %v", IsEmpty(path2))
}

func TestIsDir(t *testing.T) {
	path2 := "dir1/dir2"
	t.Logf("IsDir: %v", IsDir(path2))
}

func TestIsFile(t *testing.T) {
	path2 := "dir1/dir2"
	t.Logf("IsFile: %v", IsFile(path2))
}

func TestListIndex(t *testing.T) {
	path2 := "dir1"
	files, dirs, err := ListIndex(path2)
	t.Logf("ListFiles: %v, %v, %v", files, dirs, err)
}

func TestClearDir(t *testing.T) {
	path2 := "dir1/dir2/dir3"
	t.Logf("ClearDir: %v", ClearDir(path2))
}

func TestClearDirF(t *testing.T) {
	path2 := "dir1"
	t.Logf("ClearDirF: %v", ClearDirF(path2))
}

func TestRemoveDir(t *testing.T) {
	path2 := "dir1"
	t.Logf("RemoveDir: %v", RemoveDir(path2))
}

func TestName(t *testing.T) {
	t.Logf("Name: %v", Name("dir1/dir2/1.txt"))
}
