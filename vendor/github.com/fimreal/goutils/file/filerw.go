package file

import (
	"bufio"
	"io"
	"io/ioutil"
	"os"
)

func GetFileContentByte(fileName string) ([]byte, error) {
	return ioutil.ReadFile(fileName)
}

func WriteToFile(fileName string, b []byte) (err error) {
	return ioutil.WriteFile(fileName, b, 0644)
}

/* 简单复制文件，如果目标文件不存在，则创建，权限 0644
返回复制字节数，和 error
*/
func CopyFile(srcFilePath string, dstFilePath string) (int64, error) {

	srcFile, _ := os.OpenFile(srcFilePath, os.O_RDONLY, 0644)
	dstFile, _ := os.OpenFile(dstFilePath, os.O_RDWR|os.O_CREATE, 0644)
	defer func() {
		srcFile.Close()
		dstFile.Close()
	}()
	return io.Copy(dstFile, srcFile)
}

func CopyFileBufio(srcFilePath string, dstFilePath string) error {
	srcFile, _ := os.OpenFile(srcFilePath, os.O_RDONLY, 0644)
	dstFile, _ := os.OpenFile(dstFilePath, os.O_RDWR|os.O_CREATE, 0644)
	defer func() {
		srcFile.Close()
		dstFile.Close()
	}()

	reader := bufio.NewReader(srcFile)
	writer := bufio.NewWriter(dstFile)
	buffer := make([]byte, 4096)

	for {
		_, err := reader.Read(buffer)
		if err != nil {
			if err == io.EOF {
				// fmt.Println("文件读取完毕")
				break
			}
			return err
		}
		_, err = writer.Write(buffer)
		if err != nil {
			return err
		}
		writer.Flush()
	}
	return nil
}
