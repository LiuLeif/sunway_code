// 2017-11-14 18:54
package main

import (
	"fmt"
	"math"
)

func findPeakElement(nums []int) int {
	f := 0
	t := len(nums) - 1
	for f < t {
		m := (f + t) / 2
		left := math.MinInt64
		if m != 0 {
			left = nums[m-1]
		}
		right := math.MinInt64
		if m != len(nums)-1 {
			right = nums[m+1]
		}
		if nums[m] > left && nums[m] > right {
			return m
		}
		if nums[m] <= left {
			t = m - 1
		} else {
			f = m + 1
		}
	}
	return f
}
func main() {
	fmt.Println(findPeakElement([]int{-2147483648, -2147483647}))
}
