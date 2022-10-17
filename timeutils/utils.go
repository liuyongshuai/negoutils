// @author      Liu Yongshuai<liuyongshuai@hotmail.com>
// @file        utils.go
// @date        2022-10-17 下午9:17

package timeutils

import (
	"time"
)

func CurTimeMillis() int64 {
	return time.Now().UnixNano() / 1000000
}
