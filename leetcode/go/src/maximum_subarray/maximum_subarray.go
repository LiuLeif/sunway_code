// 2017-11-14 18:54
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

func maxSubArray(nums []int) int {
	m := make([]int, len(nums))
	for i := 0; i < len(nums); i++ {
		m[i] = math.MinInt32
	}

	m[0] = nums[0]
	ret := m[0]
	for i := 1; i < len(nums); i++ {
		m[i] = max(nums[i], m[i-1]+nums[i])
		ret = max(ret, m[i])
	}
	return ret
}
func main() {
	fmt.Println(maxSubArray([]int{-2, 1}))
}
