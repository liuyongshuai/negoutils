// @author      Liu Yongshuai<liuyongshuai@hotmail.com>
// @file        logiface.go
// @date        2022-10-17 下午9:12

package logutils

import (
	"io"
	"strings"
)

// Level 表示一个已知的日志级别。
type Level int

// 所有支持的日志级别。
const (
	PANIC Level = iota
	FATAL
	ERROR
	WARN
	INFO
	DEBUG
	printLevel
)

// String 返回 level 对应的字符串值。
func (l Level) String() string {
	switch l {
	case PANIC:
		return "PANIC"
	case FATAL:
		return "FATAL"
	case ERROR:
		return "ERROR"
	case WARN:
		return "WARN"
	case INFO:
		return "INFO"
	case DEBUG:
		return "DEBUG"
	case printLevel:
		return ""
	default:
		return "<UNKNOWN>"
	}
}

// ParseLevel 将字符串形式的 l 解析成 Level，l 不区分大小写。
// 如果 l 并没有被定义，ok 返回 false。
func ParseLevel(l string) (level Level, ok bool) {
	switch strings.ToUpper(l) {
	case "PANIC":
		return PANIC, true
	case "FATAL":
		return FATAL, true
	case "ERROR":
		return ERROR, true
	case "WARN":
		return WARN, true
	case "INFO":
		return INFO, true
	case "DEBUG", "":
		return DEBUG, true
	}
	return DEBUG, false
}

// Logger 是一个通用的日志输出接口，
type Logger interface {
	// 需要有关闭日志的能力。
	io.Closer

	// 一系列带格式的日志函数。
	Printf(format string, args ...interface{}) // 无视日志级别，始终都能输出日志，内容里不包含任何额外信息。
	Debugf(format string, args ...interface{})
	Infof(format string, args ...interface{})
	Warnf(format string, args ...interface{})
	Errorf(format string, args ...interface{})
	Fatalf(format string, args ...interface{}) // 应该调用 os.Exit 终止程序。
	Panicf(format string, args ...interface{}) // 应该调用 panic 终止程序。

	// 一系列不带格式的日志函数。
	Print(args ...interface{}) // 无视日志级别，始终都能输出日志，内容里不包含任何额外信息。
	Debug(args ...interface{})
	Info(args ...interface{})
	Warn(args ...interface{})
	Error(args ...interface{})
	Fatal(args ...interface{}) // 应该调用 os.Exit 终止程序。
	Panic(args ...interface{}) // 应该调用 panic 终止程序。
}
