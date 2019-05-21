// 2018-10-22 17:53
package main

import (
	"fmt"
	"math"
)

func maxSubarraySumCircular(A []int) int {
	// 5 -3 5 5 -3
	max := func(nums ...int) int {
		ret := math.MinInt32
		for _, n := range nums {
			if ret < n {
				ret = n
			}
		}
		return ret
	}

	min := func(nums ...int) int {
		ret := math.MaxInt32
		for _, n := range nums {
			if ret > n {
				ret = n
			}
		}
		return ret
	}
	minVal, maxVal := math.MaxInt32, math.MinInt32
	dpMin, dpMax := 0, 0
	total := 0
	for i := 0; i < len(A); i++ {
		dpMin = min(dpMin+A[i], A[i])
		dpMax = max(dpMax+A[i], A[i])
		minVal = min(minVal, dpMin)
		maxVal = max(maxVal, dpMax)
		total += A[i]
	}
	if total == minVal {
		return maxVal
	}
	return max(maxVal, total-minVal)
}

func main() {
	fmt.Println(maxSubarraySumCircular([]int{-2, -3, -1}))
}
