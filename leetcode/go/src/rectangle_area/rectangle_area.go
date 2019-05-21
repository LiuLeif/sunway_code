// 2018-01-27 20:32
package main

import (
	"fmt"
	"math"
)

func min(nums ...int) int {
	ret := math.MaxInt32
	for _, n := range nums {
		if ret > n {
			ret = n
		}
	}
	return ret
}

func max(nums ...int) int {
	ret := math.MinInt32
	for _, n := range nums {
		if ret < n {
			ret = n
		}
	}
	return ret
}

func computeArea(A int, B int, C int, D int, E int, F int, G int, H int) int {
	left := max(A, E)
	right := max(min(C, G), left)

	bottom := max(B, F)
	top := max(min(D, H), bottom)

	return (C-A)*(D-B) + (G-E)*(H-F) - (right-left)*(top-bottom)
}
func main() {
	fmt.Println(computeArea(-3, 0, 3, 4, 0, -1, 9, 2))
}
