// @author      Liu Yongshuai<liuyongshuai@hotmail.com>
// @file        utils.go
// @date        2022-10-17 下午9:14

package fileutils

import (
	"bufio"
	"bytes"
	"fmt"
	"io"
	"os"
	"path/filepath"
	"strings"
	"time"
)

//按行遍历文件的每一行
func FileLineIterator(f string, fn func(line string) error) error {
	if !FileExists(f) {
		return fmt.Errorf("file %s not exists", f)
	}
	fp, err := os.Open(f)
	if err != nil {
		return err
	}
	defer fp.Close()
	rd := bufio.NewReader(fp)
	var buf bytes.Buffer
	total := 0
	for {
		//开始按行读取文件
		line, isPrefix, err := rd.ReadLine()
		buf.Write(line)
		if !isPrefix && err == nil {
			//如果回调函数报错则直接返回
			if ferr := fn(buf.String()); ferr != nil {
				return ferr
			}
			buf.Reset()
			//校验数量
			total++
		}
		//读取文件结束
		if err == io.EOF {
			break
		}
		//出错了，退出
		if err != nil && err != io.EOF {
			return err
		}
	}
	return nil
}

//读取文件内容，按行返回
func GetFileContentLines(filePath string) (lines []string, err error) {
	err = FileLineIterator(filePath, func(line string) error {
		lines = append(lines, line)
		return nil
	})
	return
}

//读取文件内容，按行返回
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

//文件是否存在
func FileExists(filePath string) bool {
	if _, err := os.Stat(filePath); err != nil {
		if os.IsNotExist(err) {
			return false
		}
	}
	return true
}

//打开文件，得到一个句柄。如果不存在创建一个，如果存在再判断
func OpenNewFile(fileName, bakExt string, isBak bool) (fp *os.File, err error) {
	if isBak && len(bakExt) == 0 {
		bakExt = fmt.Sprintf("%s.bak", time.Now().Local().Format("20060102150405"))
	}
	if FileExists(fileName) {
		if isBak {
			err = os.Rename(fileName, fmt.Sprintf("%s.%s", fileName, bakExt))
		} else {
			err = os.Remove(fileName)
		}
		if err != nil {
			return
		}
	} else {
		//判断所在的目录是否存在
		dirName := filepath.Dir(fileName)
		if !FileExists(dirName) {
			err = os.MkdirAll(dirName, 0755)
			if err != nil {
				return
			}
		}
	}
	fp, err = os.Create(fileName)
	return
}

//文件最后修改时间
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
