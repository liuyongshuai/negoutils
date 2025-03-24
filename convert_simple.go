// @author      Liu Yongshuai<liuyongshuai@hotmail.com>
// @file        simple.go
// @date        2019-07-04 11:47

package negoutils

import (
	"fmt"
	"strconv"
)

func ToString(a interface{}) string {
	return fmt.Sprint(a)
}

func ToFloat64(a interface{}) (s float64) {
	switch v := a.(type) {
	case string:
		s, _ = strconv.ParseFloat(v, 64)
	case float64:
		s = v
	case float32:
		s = float64(v)
	case int:
		s = float64(v)
	case int16:
		s = float64(v)
	case uint16:
		s = float64(v)
	case int32:
		s = float64(v)
	case uint32:
		s = float64(v)
	case int64:
		s = float64(v)
	case uint64:
		s = float64(v)
	default:
		s = 0
	}
	return
}
