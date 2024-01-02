package Utils

import (
	"fmt"
	"os"
	"path/filepath"
)

func CreateDir(folderName string) (string, error) {
	currentPath, err := os.Getwd()
	if err != nil {
		return "", err
	}
	targetFolder := filepath.Join(currentPath, folderName)
	_, err = os.Stat(targetFolder)
	if os.IsNotExist(err) {
		// 如果文件夹不存在，则创建
		err := os.Mkdir(targetFolder, 0755)
		if err != nil {
			return "", err
		}
		fmt.Printf("Folder '%s' created.\n", folderName)
	} else if err == nil {
		// 如果文件夹已经存在
		fmt.Printf("Folder '%s' already exists.\n", folderName)
	} else {
		// 如果发生其他错误
		return "", err
	}
	return targetFolder, nil
}
