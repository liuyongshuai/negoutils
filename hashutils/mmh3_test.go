// @author      Liu Yongshuai<liuyongshuai@hotmail.com>
// @file        mmh3_test.go
// @date        2022-10-18 下午1:16

package hashutils

import (
	. "github.com/smartystreets/goconvey/convey"
	"testing"
)

func Test_ranklib(t *testing.T) {
	s := "2000000000059236943"
	hv := MurmurHash3([]byte(s))
	Convey("Hash Value", t, func() {
		So(hv, ShouldEqual, 1498636499)
	})
}
