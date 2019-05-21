// 2018-01-23 10:36
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

func min(nums ...int) int {
	ret := math.MaxInt32
	for _, n := range nums {
		if ret > n {
			ret = n
		}
	}
	return ret
}

func leftLargest(nums []int) []int {
	ret := make([]int, len(nums))
	largest := 0
	for i := 0; i < len(nums); i++ {
		largest = max(largest, nums[i])
		ret[i] = largest
	}
	return ret
}

func rightLargest(nums []int) []int {
	ret := make([]int, len(nums))
	largest := 0
	for i := len(nums) - 1; i >= 0; i-- {
		largest = max(largest, nums[i])
		ret[i] = largest
	}
	return ret
}

func trap(height []int) int {
	left := leftLargest(height)
	right := rightLargest(height)
	ret := 0
	for i := 0; i < len(height); i++ {
		ret += min(left[i], right[i]) - height[i]
	}
	return ret
}

func main() {
	fmt.Println(trap([]int{0, 1, 0, 2, 1, 0, 1, 3, 2, 1, 2, 1}))
	fmt.Println(trap([]int{4, 1, 4}))
}
