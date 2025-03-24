// @author      Liu Yongshuai<liuyongshuai@hotmail.com>
// @file        utils_test.go
// @date        2022-10-19 下午4:59

package negoutils

import (
	"fmt"
	"testing"
)

func TestMemoryGetUsage(t *testing.T) {
	fmt.Println(MemoryGetUsage())
}
