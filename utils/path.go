package utils

import (
	// "os"
	"path/filepath"
)

func JoinPath(paths ...string) string {
	_filepath := filepath.Join(paths...)
	return _filepath
}

// 检查项目必须的编译路径
// func checkPath(paths []string) {
// 	for _, path := range paths {
// 		_, err := os.Stat(JoinPath(variable.BasePath, path))
// 		fmt.Println(err)
// 		if err != nil {
// 			log.Fatal(my_errors.ErrorsConfigYamlNotExists + err.Error())
// 		}
// 	}
// }
