// @author      Liu Yongshuai<liuyongshuai@hotmail.com>
// @file        utils.go
// @date        2022-10-17 下午9:17

package timeutils

import (
	"time"
)

var (
	TimeLayoutDate     = "2006-01-02"
	TimeLayoutDateTime = "2006-01-02 15:04:05"
)

func CurTimeMillis() int64 {
	return time.Now().UnixNano() / 1000000
}

func Timestamp2Str(sec, nsec int64, format string) string {
	dt := time.Unix(sec, nsec)
	if format == "" {
		format = TimeLayoutDateTime
	}
	return dt.Format(format)
}

func Strptime(dtStr, format string) (dt time.Time, err error) {
	var TimeFormat string
	if format == "" {
		format = TimeLayoutDateTime
	}
	dt, err = time.ParseInLocation(TimeFormat, dtStr, time.Local)
	return
}

func MonthStart() time.Time {
	y, m, _ := time.Now().Date()
	return time.Date(y, m, 1, 0, 0, 0, 0, time.Local)
}

func TodayStart() time.Time {
	y, m, d := time.Now().Date()
	return time.Date(y, m, d, 0, 0, 0, 0, time.Local)
}

func TodayEnd() time.Time {
	y, m, d := time.Now().Date()
	return time.Date(y, m, d, 23, 59, 59, 1e9-1, time.Local)
}
