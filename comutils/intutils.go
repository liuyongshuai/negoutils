// @author      Liu Yongshuai<liuyongshuai@hotmail.com>
// @file        intutils.go
// @date        2022-11-09 下午7:49

package comutils

func MaxInt(is ...int) int {
	if len(is) == 0 {
		return 0
	}
	m := is[0]
	for _, v := range is {
		if v > m {
			m = v
		}
	}
	return m
}

func MinInt(is ...int) int {
	if len(is) == 0 {
		return 0
	}
	m := is[0]
	for _, v := range is {
		if v < m {
			m = v
		}
	}
	return m
}

func MaxInt64(is ...int64) int64 {
	if len(is) == 0 {
		return 0
	}
	m := is[0]
	for _, v := range is {
		if v > m {
			m = v
		}
	}
	return m
}

func MinInt64(is ...int64) int64 {
	if len(is) == 0 {
		return 0
	}
	m := is[0]
	for _, v := range is {
		if v < m {
			m = v
		}
	}
	return m
}
