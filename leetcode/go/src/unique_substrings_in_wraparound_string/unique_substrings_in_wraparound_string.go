// 2018-01-03 14:30
package main

import (
	"fmt"
	"math"
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

func findSubstringInWraproundString(p string) int {
	if len(p) == 0 {
		return 0
	}
	increasing := make([]int, len(p))
	increasing[0] = 1
	for i := 1; i < len(p); i++ {
		if p[i] == p[i-1]+1 || (p[i] == 'a' && p[i-1] == 'z') {
			increasing[i] = increasing[i-1] + 1
		} else {
			increasing[i] = 1
		}
	}
	maxIncresing := map[byte]int{}
	for i := 0; i < len(p); i++ {
		maxIncresing[p[i]] = max(maxIncresing[p[i]], increasing[i])
	}
	ret := 0
	for _, v := range maxIncresing {
		ret += v
	}

	return ret
}
func main() {
	fmt.Println(findSubstringInWraproundString("zabz"))
}
