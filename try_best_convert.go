package negoutils

import (
	"encoding/json"
	"math"
	"reflect"
	"strconv"
	"strings"
)

// 尽最大努力转换为float64
func TryBestToFloat(value interface{}) (float64, error) {
	val := reflect.ValueOf(value)
	basicKind, err := GetBasicKind(val)
	if err != nil {
		return 0, err
	}
	switch basicKind {
	case IntKind:
		return float64(val.Int()), nil
	case UintKind:
		return float64(val.Uint()), nil
	case StringKind:
		floatStr := getFloatStrFromLeft(val.String())
		if len(floatStr) <= 0 {
			return 0, nil
		}
		return strconv.ParseFloat(floatStr, 10)
	case FloatKind:
		return val.Float(), nil
	case BoolKind:
		if val.Bool() {
			return 1, nil
		}
		return 0, nil
	case PtrKind:
		if val.IsNil() {
			return 0, nil
		}
		return TryBestToFloat(val.Elem().Interface())
	default:
		return 0, ErrorInvalidInputType
	}
}

// 尽最大努力转为bool类型
func TryBestToBool(value interface{}) (bool, error) {
	val := reflect.ValueOf(value)
	basicKind, err := GetBasicKind(val)
	if err != nil {
		return false, err
	}
	switch basicKind {
	case FloatKind:
		return val.Float() != 0, nil
	case IntKind:
		return val.Int() != 0, nil
	case UintKind:
		return val.Uint() != 0, nil
	case StringKind:
		v := strings.TrimSpace(val.String())
		if len(v) > 0 {
			return true, nil
		}
		return false, nil
	case BoolKind:
		return val.Bool(), nil
	case PtrKind:
		if val.IsNil() {
			return false, nil
		}
		return TryBestToBool(val.Elem().Interface())
	case FuncKind:
		return !val.IsNil(), nil
	}

	//对于Array, Chan, Map, Slice长度>0即可
	switch val.Kind() {
	case reflect.Array, reflect.Chan, reflect.Slice, reflect.Map:
		return val.Len() != 0, nil
	}
	return false, ErrorInvalidInputType
}

// 尽最大努力转换为字符串
func TryBestToString(value interface{}) (string, error) {
	val := reflect.ValueOf(value)
	basicKind, err := GetBasicKind(val)
	if err != nil {
		return "", err
	}
	switch basicKind {
	case IntKind:
		return strconv.FormatInt(val.Int(), 10), nil
	case UintKind:
		return strconv.FormatUint(val.Uint(), 10), nil
	case StringKind:
		return val.String(), nil
	case FloatKind:
		return strconv.FormatFloat(val.Float(), 'f', -1, 64), nil
	case BoolKind:
		return strconv.FormatBool(val.Bool()), nil
	case PtrKind:
		if val.IsNil() {
			return "nil", nil
		}
		return TryBestToString(val.Elem().Interface())
	case ContainerKind:
		result, err := json.Marshal(value)
		if err != nil {
			return "", err
		}
		return string(result), err
	default:
		return val.String(), nil
	}
}

// 尽最大努力将给定的类型转换为int64
// 如"45.67"->45、"98.4abc3"->98、"34.87"->34
func TryBestToInt64(value interface{}) (int64, error) {
	ret, err := tryBestConvertAnyTypeToInt(value, false)
	if err != nil {
		return 0, err
	}
	val := reflect.ValueOf(ret)
	return val.Int(), nil
}

// 尽最大努力将给定的类型转换为uint64
// 如"45.67"->45、"98.4abc3"->98、"34.87"->34
func TryBestToUint64(value interface{}) (uint64, error) {
	ret, err := tryBestConvertAnyTypeToInt(value, true)
	if err != nil {
		return 0, err
	}
	val := reflect.ValueOf(ret)
	return val.Uint(), nil
}

// 尽最大努力将任意类型转为int64或uint64
func tryBestConvertAnyTypeToInt(value interface{}, isUnsigned bool) (interface{}, error) {
	val := reflect.ValueOf(value)
	basicKind, err := GetBasicKind(val)
	if err != nil {
		return int64(0), err
	}
	switch basicKind {
	case IntKind:
		v := val.Int()
		if isUnsigned {
			if v >= 0 {
				return uint64(v), nil
			}
			return uint64(0), ErrorLessThanZero
		}
		return int64(v), nil
	case UintKind:
		v := val.Uint()
		if isUnsigned {
			return uint64(v), nil
		}
		if v > math.MaxInt64 {
			return int64(0), ErrorOverflowMaxInt64
		}
		return int64(v), nil
	case StringKind: //取连续的最长的数字或小数点
		floatStr := getFloatStrFromLeft(val.String())
		if len(floatStr) <= 0 {
			if isUnsigned {
				return uint64(0), nil
			}
			return int64(0), nil
		}
		//先转成float，因为将"45.33"直接转为int/uint时会报错
		f, err := strconv.ParseFloat(floatStr, 10)
		if err != nil {
			return int64(0), err
		}
		return tryBestConvertAnyTypeToInt(f, isUnsigned)
		//float特殊处理，会有科学记数法表示形式
	case FloatKind:
		f := val.Float()
		if isUnsigned {
			if f > MaxUint64Float {
				return uint64(0), ErrorOverflowMaxUint64
			}
			if f < 0 {
				return uint64(0), ErrorLessThanZero
			}
			return uint64(f), nil
		}
		if f > MaxInt64Float {
			return int64(0), ErrorOverflowMaxInt64
		}
		if f < MinInt64Float {
			return int64(0), ErrorLessThanMinInt64
		}
		return int64(f), nil
	case BoolKind:
		b := val.Bool()
		tmp := 0
		if b {
			tmp = 1
		}
		if isUnsigned {
			return uint64(tmp), nil
		}
		return int64(tmp), nil
		//指针类型递归调用，直到取本值为止
	case PtrKind:
		if val.IsNil() {
			if isUnsigned {
				return uint64(0), nil
			}
			return int64(0), nil
		}
		return tryBestConvertAnyTypeToInt(val.Elem().Interface(), isUnsigned)
	default:
		return int64(0), ErrorInvalidInputType
	}
}
