// @author      Liu Yongshuai<liuyongshuai@hotmail.com>
// @file        strutils.go
// @date        2022-11-09 下午7:30

package comutils

//去重（不保证原顺序）
func UniqueStrSlice(slice []string) []string {
	tmp := map[string]bool{}
	for _, v := range slice {
		tmp[v] = true
	}
	var ret []string
	for t := range tmp {
		ret = append(ret, t)
	}
	return ret
}

//检查是否在slice里面
func InStrSlice(v string, sl []string) bool {
	for _, vv := range sl {
		if vv == v {
			return true
		}
	}
	return false
}
