// @author      Liu Yongshuai<liuyongshuai@hotmail.com>
// @file        utils.go
// @date        2022-10-17 下午9:14

package negoutils

import (
	"os"
	"path/filepath"
	"strings"
)

// 读取文件内容，按行返回
func GetFileContentLines(filePath string) (lines []string, err error) {
	err = FileLineIterator(filePath, func(line string) error {
		lines = append(lines, line)
		return nil
	})
	return
}

// 读取文件内容，按行返回
func GetFileContentStr(filePath string) (str string, err error) {
	var lines []string
	err = FileLineIterator(filePath, func(line string) error {
		lines = append(lines, line)
		return nil
	})
	str = strings.Join(lines, "\n")
	str = strings.TrimSpace(str)
	return
}

// 文件是否存在
func FileExists(filePath string) bool {
	if _, err := os.Stat(filePath); err != nil {
		if os.IsNotExist(err) {
			return false
		}
	}
	return true
}

// 文件最后修改时间
func FileModTime(filename string) (int64, error) {
	fd, err := os.Open(filename)
	if err != nil {
		return 0, err
	}
	defer fd.Close()
	fileinfo, err := fd.Stat()
	if err != nil {
		return 0, err
	}
	return fileinfo.ModTime().Unix(), nil
}

// 当前路径
func CurPath() string {
	path, _ := filepath.Abs(os.Args[0])
	return path
}

// 当前目录
func CurDir() string {
	return filepath.Dir(CurPath())
}
