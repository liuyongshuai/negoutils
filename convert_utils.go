package negoutils

import (
	"bytes"
	"encoding/binary"
	"errors"
	"math"
	"reflect"
	"strings"
	"unsafe"
)

type Basickind int

const (
	//转换时最大值
	MaxInt64Float  = float64(math.MaxInt64)
	MinInt64Float  = float64(math.MinInt64)
	MaxUint64Float = float64(math.MaxUint64)
	//基本类型归纳，类型转换时用得着
	InvalidKind Basickind = iota
	BoolKind
	ComplexKind
	IntKind
	FloatKind
	StringKind
	UintKind
	PtrKind
	ContainerKind
	FuncKind
)

var (
	//一些错误信息
	ErrorOverflowMaxInt64  = errors.New("this value overflow math.MaxInt64")
	ErrorOverflowMaxUint64 = errors.New("this value overflow math.MaxUint64")
	ErrorLessThanMinInt64  = errors.New("this value less than math.MinInt64")
	ErrorLessThanZero      = errors.New("this value less than zero")
	ErrorBadComparisonType = errors.New("invalid type for comparison")
	ErrorBadComparison     = errors.New("incompatible types for comparison")
	ErrorNoComparison      = errors.New("missing argument for comparison")
	ErrorInvalidInputType  = errors.New("invalid input type")
)

// 转换成特定类型，便于判断
func GetBasicKind(v reflect.Value) (Basickind, error) {
	switch v.Kind() {
	case reflect.Bool:
		return BoolKind, nil
	case reflect.Int, reflect.Int8, reflect.Int16, reflect.Int32, reflect.Int64:
		return IntKind, nil
	case reflect.Uint, reflect.Uint8, reflect.Uint16, reflect.Uint32, reflect.Uint64, reflect.Uintptr:
		return UintKind, nil
	case reflect.Float32, reflect.Float64:
		return FloatKind, nil
	case reflect.Complex64, reflect.Complex128:
		return ComplexKind, nil
	case reflect.String:
		return StringKind, nil
	case reflect.Ptr:
		return PtrKind, nil
	case reflect.Struct, reflect.Map, reflect.Slice, reflect.Array:
		return ContainerKind, nil
	case reflect.Func:
		return FuncKind, nil
	}
	return InvalidKind, ErrorInvalidInputType
}

// int64转byte
func Int64ToBytes(i int64) []byte {
	var buf = make([]byte, 8)
	binary.BigEndian.PutUint64(buf, uint64(i))
	return buf
}

// bytes转int64
func BytesToInt64(buf []byte) int64 {
	return int64(binary.BigEndian.Uint64(buf))
}

// float64转byte
func Float64ToByte(float float64) []byte {
	bits := math.Float64bits(float)
	bs := make([]byte, 8)
	binary.BigEndian.PutUint64(bs, bits)
	return bs
}

// bytes转float
func ByteToFloat64(bytes []byte) float64 {
	bits := binary.BigEndian.Uint64(bytes)
	return math.Float64frombits(bits)
}

// 字符串转为字节切片
func StrToByte(s string) []byte {
	x := (*[2]uintptr)(unsafe.Pointer(&s))
	h := [3]uintptr{x[0], x[1], x[1]}
	return *(*[]byte)(unsafe.Pointer(&h))
}

// 字节切片转为字符串
func ByteToStr(b []byte) string {
	return *(*string)(unsafe.Pointer(&b))
}

// []byte -> []int64
func BytesToInt64Slice(req *[]byte) (ret []int64) {
	if req == nil {
		return
	}
	rlen := len(*req)
	if rlen <= 0 || rlen%8 != 0 {
		return
	}
	bytSegNum := rlen / 8
	for i := 0; i < bytSegNum; i++ {
		s := i * 8
		ret = append(ret, BytesToInt64((*req)[s:s+8]))
	}
	return
}

// []int64 -> []byte
func Int64SliceToBytes(req *[]int64) (ret []byte) {
	if req == nil {
		return
	}
	buf := bytes.Buffer{}
	for _, r := range *req {
		buf.Write(Int64ToBytes(r))
	}
	ret = buf.Bytes()
	return
}

// 从左边开始提取数据及小数点
func getFloatStrFromLeft(val string) string {
	val = strings.TrimSpace(val)
	valBytes := StrToByte(val)
	buf := bytes.Buffer{}
	for _, b := range valBytes {
		if b >= 48 && b <= 57 || b == 46 {
			buf.WriteByte(b)
			continue
		}
		break
	}
	return buf.String()
}
