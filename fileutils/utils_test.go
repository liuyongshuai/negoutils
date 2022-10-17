// @author      Liu Yongshuai<liuyongshuai@hotmail.com>
// @file        utils_test.go
// @date        2022-10-17 下午9:28

package fileutils

import (
	"fmt"
	"testing"
)

func TestCurPath(t *testing.T) {
	fmt.Println(CurDir())
	fmt.Println(CurPath())
}

