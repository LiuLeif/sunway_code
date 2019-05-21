// 2018-10-15 10:10
package main

import (
	"fmt"
	"math"
)

func smallestRangeI(A []int, K int) int {
	small := math.MaxInt32
	for _, v := range A {
		if v < small {
			small = v
		}
	}
	large := math.MinInt32
	for _, v := range A {
		if v > large {
			large = v
		}
	}
	ret := large - K - small - K
	if ret < 0 {
		ret = 0
	}
	return ret
}
func main() {
	fmt.Println(smallestRangeI([]int{10}, 2))
}
