// @author      Liu Yongshuai<liuyongshuai@hotmail.com>
// @file        conv_test.go
// @date        2019-08-20 14:15

package convertutils

import (
	"fmt"
	"github.com/kr/pretty"
	"regexp"
	"runtime"
	"testing"
)

// 获取调用者信息
func CallerName(skip int) (name, file string, line int, ok bool) {
	var (
		reInit    = regexp.MustCompile(`init·\d+$`) // main.init·1
		reClosure = regexp.MustCompile(`func·\d+$`) // main.func·001
	)
	for {
		var pc uintptr
		if pc, file, line, ok = runtime.Caller(skip + 1); !ok {
			return
		}
		name = runtime.FuncForPC(pc).Name()
		if reInit.MatchString(name) {
			name = reInit.ReplaceAllString(name, "init")
			return
		}
		if reClosure.MatchString(name) {
			skip++
			continue
		}
		return
	}
	return
}

func TestTryBestConvert(t *testing.T) {
	p1 := 45649065094658405684504232323223322334.555
	p2 := "45s89s"
	p3 := "wendao"
	p4 := &p2
	vals := []interface{}{
		"34343434",
		44.3222,
		989889,
		0.222,
		&p1,
		&p2,
		&p3,
		&p4,
		"",
		true,
		-22222,
	}
	for _, val := range vals {
		int64Val, int64Err := TryBestToInt64(val)
		uint64Val, uint64Err := TryBestToUint64(val)
		floatVal, floatErr := TryBestToFloat(val)
		strVal, strErr := TryBestToString(val)
		boolVal, boolErr := TryBestToBool(val)
		fmt.Printf("rawVal %# v \tint64[%v %v] uint64[%v %v] float[%v %v] str[%v %v] bool[%v %v]\n",
			pretty.Formatter(val),
			pretty.Formatter(int64Val), int64Err,
			pretty.Formatter(uint64Val), uint64Err,
			pretty.Formatter(floatVal), floatErr,
			pretty.Formatter(strVal), strErr,
			pretty.Formatter(boolVal), boolErr,
		)
	}

}

func TestByteToStr(t *testing.T) {
	var bt []byte
	a := ByteToStr(bt)
	fmt.Printf("%# v\n", pretty.Formatter(a))
}
