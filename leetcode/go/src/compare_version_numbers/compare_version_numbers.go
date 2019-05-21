// 2018-02-07 17:27
package main

import (
	"fmt"
	"math"
	"strings"
)

func max(nums ...int) int {
	ret := math.MinInt32
	for _, n := range nums {
		if ret < n {
			ret = n
		}
	}
	return ret
}

func compareString(v1 string, v2 string) int {
	// 02 = 2
	// 10 > 2
	l1, l2 := len(v1), len(v2)
	if l1 > l2 {
		v2 = strings.Repeat("0", l1-l2) + v2
	}
	if l1 < l2 {
		v1 = strings.Repeat("0", l2-l1) + v1
	}
	if v1 > v2 {
		return 1
	} else if v1 < v2 {
		return -1
	} else {
		return 0
	}
}

func compareVersion(version1 string, version2 string) int {
	v1 := strings.Split(version1, ".")
	v2 := strings.Split(version2, ".")
	for i := 0; i < max(len(v1), len(v2)); i++ {
		x := "0"
		y := "0"
		if i < len(v1) {
			x = v1[i]
		}
		if i < len(v2) {
			y = v2[i]
		}
		// 00 11
		t := compareString(x, y)
		if t != 0 {
			return t
		}
	}
	return 0
}
func main() {
	fmt.Println(compareVersion("1.0", "1.1"))
}
