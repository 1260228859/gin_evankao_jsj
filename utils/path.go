package utils

import (
	"os"
	"path/filepath"
)

func JoinPath(paths ...string) string {
	_filepath := filepath.Join(paths...)
	return _filepath
}


// 检查文件/路径是否存在
func CheckPathIsExits(path string) bool {
	_, err := os.Stat(path)
	return err != nil
}
