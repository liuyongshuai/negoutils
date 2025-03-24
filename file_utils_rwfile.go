// @author      Liu Yongshuai<liuyongshuai@hotmail.com>
// @file        rwfile.go
// @date        2022-10-17 下午9:30

package negoutils

import (
	"bufio"
	"bytes"
	"fmt"
	"io"
	"os"
	"path/filepath"
	"sort"
	"strings"
	"time"
)

// 按行遍历文件的每一行
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

// 打开文件，得到一个句柄。如果不存在创建一个，如果存在再判断
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

// 读取目录下的所有文件
func ReadDirFiles(dir string) (ret []string, err error) {
	if !FileExists(dir) {
		err = fmt.Errorf("dir %v not exists", dir)
		return ret, err
	}
	dir = strings.TrimRight(dir, "/")

	//开始读取目录
	var dp *os.File
	dp, err = os.Open(dir)
	if err != nil {
		return ret, err
	}
	if dp == nil {
		err = fmt.Errorf("open %v failed", dir)
		return ret, err
	}

	defer dp.Close()

	var dlist []os.FileInfo
	dlist, err = dp.Readdir(-1)
	if err != nil {
		return ret, err
	}

	//开始递归读取目录，忽略掉隐藏目录及文件
	for _, v := range dlist {
		if strings.HasPrefix(v.Name(), ".") {
			continue
		}
		f := dir + "/" + v.Name()
		if v.IsDir() {
			tret, terr := ReadDirFiles(f)
			if terr == nil && tret != nil && len(tret) > 0 {
				ret = append(ret, tret...)
			}
		} else {
			ret = append(ret, f)
		}
	}
	ret = UniqueStrSlice(ret)
	sort.Strings(ret)
	return ret, nil
}
