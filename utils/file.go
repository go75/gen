package utils

import (
	"os"
)

// 判断路径是否是文件夹
func IsDir(filePath string) bool {
	fileInfo, err := os.Stat(filePath)
	if err != nil {
		return false
	}

	return fileInfo.IsDir()
}

func MakeDir(dirPath string) error {
	err := os.Mkdir(dirPath, 0764)
	if err == nil {
		return nil
	}

	info, err := os.Stat(dirPath)
	if err != nil {
		return err
	}

	if info.IsDir() {
		return nil
	}
	return err
}