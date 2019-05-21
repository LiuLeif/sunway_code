// 2018-10-23 17:24
package main

import (
	"fmt"
	"math"
	"sort"
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

func backtrack(A []int, K int, minVal, maxVal int) int {
	if len(A) == 0 {
		return maxVal - minVal
	}
	return min(backtrack(A[1:], K, min(minVal, A[0]-K), max(maxVal, A[0]-K)), backtrack(A[1:], K, min(minVal, A[0]+K), max(maxVal, A[0]+K)))

}

func smallestRangeII2(A []int, K int) int {
	// 1 2 | 3 6
	return backtrack(A, K, math.MaxInt32, math.MinInt32)
}

func smallestRangeII(A []int, K int) int {
	// 1 2 | 3 6
	min := func(nums ...int) int {
		ret := math.MaxInt32
		for _, n := range nums {
			if ret > n {
				ret = n
			}
		}
		return ret
	}

	max := func(nums ...int) int {
		ret := math.MinInt32
		for _, n := range nums {
			if ret < n {
				ret = n
			}
		}
		return ret
	}
	// 1 2 3 4
	sort.Ints(A)
	ret := A[len(A)-1] - A[0]
	for i := 0; i < len(A)-1; i++ {
		la, lb := A[0]+K, A[i]+K
		ra, rb := A[i+1]-K, A[len(A)-1]-K
		ret = min(ret, max(rb-ra, lb-la, rb-la, lb-ra))
	}
	return ret
}

func main() {
	fmt.Println(smallestRangeII([]int{0, 10}, 3))
}
