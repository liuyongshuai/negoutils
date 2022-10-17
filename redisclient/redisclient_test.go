// @author      Liu Yongshuai<liuyongshuai@hotmail.com>
// @file        redisclient_test.go
// @date        2022-06-24 10:53

package redisclient

import (
	"fmt"
	"testing"
	"time"
)

func TestRedisClient(t *testing.T) {
	r, e := InitRedisClient("100.69.239.173:3000", 60*time.Second)
	if e != nil {
		fmt.Println(e)
		return
	}
	v, e := r.Get("vector_20220620:raw_1080125649252777986")
	fmt.Println(v, e)
}
