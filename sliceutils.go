// @author      Liu Yongshuai<liuyongshuai@hotmail.com>
// @file        sliceutils.go
// @date        2022-11-09 下午7:49

package negoutils

import (
	"math/rand"
	"reflect"
)

// 去重（不保证原顺序）
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

// 检查是否在slice里面
func InStrSlice(v string, sl []string) bool {
	for _, vv := range sl {
		if vv == v {
			return true
		}
	}
	return false
}

// 打乱一个字符串slice
func StrSliceShuffle(slice []string) []string {
	sl := len(slice)
	if sl <= 0 {
		return slice
	}
	for i := 0; i < sl; i++ {
		a := rand.Intn(sl)
		b := rand.Intn(sl)
		slice[a], slice[b] = slice[b], slice[a]
	}
	return slice
}

// 将其他类型转为interface的切片，支持slice/map/[u]int[8-64]/float../string
func ToSliceIface(in interface{}) []interface{} {
	vin := reflect.ValueOf(in)
	switch vin.Kind() {
	case reflect.Slice, reflect.Array: //in为slice类型
		vlen := vin.Len()
		ret := make([]interface{}, vlen)
		for i := 0; i < vlen; i++ {
			ret[i] = vin.Index(i).Interface()
		}
		return ret
	case reflect.Map: //in为map类型
		ks := vin.MapKeys()
		vlen := vin.Len()
		ret := make([]interface{}, vlen)
		for _, k := range ks {
			ret = append(ret, vin.MapIndex(k).Interface())
		}
		return ret
	case reflect.String: //字符串类型
		tmp := []byte(vin.String())
		var ret []interface{}
		for _, t := range tmp {
			ret = append(ret, t)
		}
		return ret
	default:
		return []interface{}{vin.Interface()}
	}
}

// 检查interface类型是否在slice里
func InSlice(val interface{}, sl []interface{}) bool {
	for _, sval := range sl {
		if sval == val {
			return true
		}
	}
	return false
}
