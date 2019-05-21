// 2017-12-27 17:05
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

func maxArea(height []int) int {
	ret, left, right := 0, 0, len(height)-1
	for left < right {
		area := min(height[left], height[right]) * (right - left)
		ret = max(ret, area)
		if height[left] < height[right] {
			left += 1
		} else {
			right -= 1
		}
	}
	return ret
}
func main() {
	height := []int{1, 3, 2, 5, 25, 24, 5}
	fmt.Println(maxArea(height))
}
