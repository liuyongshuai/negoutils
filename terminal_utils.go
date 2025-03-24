// @author      Liu Yongshuai<liuyongshuai@hotmail.com>
// @file        utils.go
// @date        2020-01-25 16:43

package negoutils

import (
	"golang.org/x/crypto/ssh/terminal"
	"os"
)

var (
	ScreenWidth  int
	ScreenHeight int
)

func init() {
	var err error
	ScreenWidth, ScreenHeight, err = GetTerminalSize()
	if err != nil {
		ScreenWidth = 200
		ScreenHeight = 200
	}
}

// 获取当前终端的宽、高信息：字符数，非终端时（如IDE的执行环境）会报错
func GetTerminalSize() (width, height int, err error) {
	return terminal.GetSize(int(os.Stdout.Fd()))
}
