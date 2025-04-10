package fileutil

import (
	"errors"
	"os"
	"path"
	"path/filepath"
)

// SelfPath gets compiled executable file absolute path.
func SelfPath() string {
	selfPath, _ := filepath.Abs(os.Args[0])
	return selfPath
}

// SelfDir gets compiled executable file directory.
func SelfDir() string {
	return filepath.Dir(SelfPath())
}

// MkDir 创建目录.
func MkDir(dirPath string) error {
	if IsExist(dirPath) {
		return nil
	}
	return os.MkdirAll(dirPath, os.ModePerm)
}

// IsExist 判断文件或目录是否存在.
func IsExist(filePath string) bool {
	_, err := os.Stat(filePath)
	return err == nil || os.IsExist(err)
}

// IsEmpty 判断目录是否为空.
func IsEmpty(dirname string) bool {
	dir, _ := os.ReadDir(dirname)
	if len(dir) == 0 {
		return true
	} else {
		return false
	}
}

// IsFile 判断文件是否存在.
func IsFile(filePath string) bool {
	f, e := os.Stat(filePath)
	if e != nil {
		return false
	}
	return !f.IsDir()
}

// IsDir 判断目录是否存在.
func IsDir(filePath string) bool {
	f, e := os.Stat(filePath)
	if e != nil {
		return false
	}
	return f.IsDir()
}

// ListIndex 目录下文件名和子目录名列表.
func ListIndex(dirPath string) ([]string, []string, error) {
	if !IsDir(dirPath) {
		return nil, nil, errors.New(dirPath + " not a directory")
	}
	fs, err := os.ReadDir(dirPath)
	var files, dirs []string
	for _, fi := range fs {
		if !fi.IsDir() {
			files = append(files, fi.Name())
		} else {
			dirs = append(dirs, fi.Name())
		}
	}
	return files, dirs, err
}

// ClearDir 清空目录下所有文件不包括子目录.
func ClearDir(dirPath string) error {
	dir, e := os.ReadDir(dirPath)
	for _, d := range dir {
		if d.IsDir() {
			continue
		}
		_ = os.Remove(path.Join([]string{dirPath, d.Name()}...))
	}
	return e
}

// ClearDirF 清空目录下所有文件和目录.
func ClearDirF(dirPath string) error {
	dir, e := os.ReadDir(dirPath)
	for _, d := range dir {
		_ = os.RemoveAll(path.Join([]string{dirPath, d.Name()}...))
	}
	return e
}

// RemoveDir 删除空目录.
func RemoveDir(dirPath string) error {
	if !IsDir(dirPath) {
		return errors.New(dirPath + " not a directory")
	}
	return os.Remove(dirPath)
}

// GetAllFiles 获取指定目录下的所有文件路径
func GetAllFiles(dirPath string) ([]string, error) {
	var files []string
	// 使用filepath.Walk遍历目录
	if err := filepath.Walk(dirPath, func(path string, info os.FileInfo, err error) error {
		if err != nil {
			return err
		}
		// 如果是文件而非目录，则添加到结果中
		if !info.IsDir() {
			files = append(files, path)
		}
		return nil
	}); err != nil {
		return nil, err
	}
	return files, nil
}
