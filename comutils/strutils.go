// @author      Liu Yongshuai<liuyongshuai@hotmail.com>
// @file        strutils.go
// @date        2022-11-09 下午7:30

package comutils

import (
	"bytes"
	"fmt"
	"github.com/cespare/xxhash/v2"
	"github.com/liuyongshuai/negoutils/convertutils"
	"math/rand"
	"regexp"
	"strings"
	"unicode"
)

//从字符串里提取单字，只要中文汉字
func ExtractCNWord(str string) (ret []string) {
	ret = make([]string, 0)
	for _, r := range str {
		if unicode.Is(unicode.Scripts["Han"], r) {
			ret = append(ret, string(r))
		}
	}
	return
}

//是否全是汉字
func IsAllChinese(str string) bool {
	if len(str) <= 0 {
		return false
	}
	ret := true
	for _, r := range str {
		if !unicode.Is(unicode.Scripts["Han"], r) {
			ret = false
			break
		}
	}
	return ret
}

//是否为汉字、字母、数字
func IsNormalStr(str string) bool {
	if len(str) <= 0 {
		return false
	}
	reg := regexp.MustCompile("^[a-zA-Z0-9\u4e00-\u9fa5]+$")
	return reg.MatchString(str)
}

//字符串hash为uint64
func StrHashSum64(str string) uint64 {
	return xxhash.Sum64(convertutils.StrToByte(str))
}

//打乱一个字符串
func StrShuffle(str string) string {
	rs := []rune(str)
	sliceLen := len(rs)
	for i := 0; i < sliceLen; i++ {
		a := rand.Intn(sliceLen)
		b := rand.Intn(sliceLen)
		rs[a], rs[b] = rs[b], rs[a]
	}
	return string(rs)
}

//半角字符转全角字符【处理搜索的query用】
func ToDBC(str string) string {
	ret := ""
	for _, r := range str {
		if r == 32 {
			ret += string(12288)
		} else if r < 127 {
			ret += string(r + 65248)
		} else {
			ret += string(r)
		}
	}
	return ret
}

//全角字符转半角字符【处理搜索的query用】
func ToCBD(str string) string {
	ret := ""
	for _, r := range str {
		if r == 12288 {
			ret += string(r - 12256)
			continue
		}
		if r > 65280 && r < 65375 {
			ret += string(r - 65248)
		} else {
			ret += string(r)
		}
	}
	return ret
}

//判断字符串是否全为数字
func IsAllNumber(str string) bool {
	if len(str) <= 0 {
		return false
	}
	if len(strings.Trim(str, "0123456789")) > 0 {
		return false
	}
	return true
}

var (
	AlphaAll       = []byte(`0123456789ABCDEFGHIJKLMNOPQRSTUVWXYZabcdefghijklmnopqrstuvwxyz`)
	alphaNumLetter = []byte(`0123456789abcdefghijklmnopqrstuvwxyz`)
	AlphaAllLetter = []byte(`abcdefghijklmnopqrstuvwxyz`)
	AlphaAllNum    = []byte(`0123456789`)
)

//生成一堆随机数
func RandomStr(n int, alphabets ...byte) string {
	if len(alphabets) == 0 {
		alphabets = AlphaAll
	}
	var byteSlice = make([]byte, n)
	var randBy bool
	if num, err := rand.Read(byteSlice); num != n || err != nil {
		randBy = true
	}
	for i, b := range byteSlice {
		if randBy {
			byteSlice[i] = alphabets[rand.Intn(len(alphabets))]
		} else {
			byteSlice[i] = alphabets[b%byte(len(alphabets))]
		}
	}
	return string(byteSlice)
}

//高仿PHP的 preg_replace_callback
//pattern：正则表达式
//originStr：要处理的字符串
//fn：参数是字符串切片，0表示整个匹配的字符串，1表示正则里的第一个捕获项、2表示第二个、依次类推。。。。
func PregReplaceCallback(pattern, originStr string, fn func([]string) string) (string, error) {
	reg, err := regexp.Compile(pattern)
	if err != nil {
		return originStr, err
	}
	originLen := len(originStr)
	if originLen <= 0 {
		return originStr, err
	}
	buf := bytes.Buffer{}
	var endIndex, preIndex int
	//提取出所有的子字符串
	subList := reg.FindAllStringSubmatchIndex(originStr, -1)
	if len(subList) <= 0 {
		return originStr, nil
	}
	for _, subInfo := range subList {
		//subs的结构类似：[9,12,23,45]，表示匹配项的起止位置，必须是偶数个
		subLen := len(subInfo)
		//这玩意吧，啥玩意也没匹配上
		if subLen%2 != 0 || subLen < 2 {
			return originStr, fmt.Errorf("invalid sub match")
		}
		var matches []string
		//提取所有的匹配项
		for i := 0; i < subLen; i += 2 {
			si := subInfo[i]
			ei := subInfo[i+1]
			matches = append(matches, originStr[si:ei])
		}
		startIndex := subInfo[0]
		endIndex = subInfo[1]
		//上一次循环的结束位置：本次的开始位置的字符填充到结果里
		buf.WriteString(originStr[preIndex:startIndex])
		buf.WriteString(fn(matches))
		preIndex = endIndex
	}
	//如果没到字符串的末尾，全部填充进去即可
	if endIndex < originLen {
		buf.WriteString(originStr[endIndex:])
	}
	if buf.Len() <= 0 {
		return originStr, nil
	}
	return buf.String(), nil
}
