// 2017-12-08 14:35
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

func findLengthOfLCIS(nums []int) int {
	if len(nums) == 0 {
		return 0
	}
	ret := 1
	left := 0
	for i := 1; i < len(nums); i++ {
		if nums[i] <= nums[i-1] {
			ret = max(ret, i-left)
			left = i
		}
	}

	return max(ret, len(nums)-left)
}
func main() {
	fmt.Println(findLengthOfLCIS([]int{1}))
}
