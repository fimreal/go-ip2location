package file

import "os"

// 判断所给路径文件/文件夹是否存在
func PathExists(path string) bool {
	_, err := os.Stat(path) //os.Stat获取文件信息
	if err != nil {
		if os.IsExist(err) {
			return true
		}
		return false
	}
	return true
}

// 判断所给路径文件/文件夹是否存在, 简化版本
// func PathExists(path string) bool {
//     if _, err := os.Stat(path); err != nil {
//         return !os.IsNotExist(err)
//     }
//     return true
// }

// 判断所给路径是否为文件夹
func PathIsDir(path string) bool {
	s, err := os.Stat(path) // s 是 fs.FileInfo
	if err != nil {
		return false
	}
	return s.IsDir() // 内置函数，和该函数名无关
}

// 判断所给路径是否为文件
func PathIsFile(path string) bool {
	return !PathIsDir(path)
}
