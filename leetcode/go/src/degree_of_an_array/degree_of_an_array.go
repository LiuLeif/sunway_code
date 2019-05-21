// 2017-12-04 17:39
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

func findShortestSubArray(nums []int) int {
	counter := make(map[int]int)
	first := make(map[int]int)
	last := make(map[int]int)

	maxValue := math.MinInt32
	for k, v := range nums {
		if first[v] == 0 {
			first[v] = k + 1
		}
		last[v] = k + 1
		counter[v]++
		maxValue = max(maxValue, counter[v])
	}

	ret := math.MaxInt32
	for k, v := range counter {
		if maxValue == v {
			ret = min(last[k]-first[k]+1, ret)
		}
	}

	return ret
}
func main() {
	fmt.Println(findShortestSubArray([]int{1, 2, 1}))
}
