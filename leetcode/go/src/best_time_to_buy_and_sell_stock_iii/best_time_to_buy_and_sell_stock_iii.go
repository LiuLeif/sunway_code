// 2017-11-21 16:24
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

func maxProfit(nums []int) int {
	if len(nums) == 0 {
		return 0
	}
	full2 := make([]int, len(nums))
	empty2 := make([]int, len(nums))
	full1 := make([]int, len(nums))
	empty1 := make([]int, len(nums))

	full2[0] = -nums[0]
	empty2[0] = 0
	full1[0] = -nums[0]
	empty1[0] = 0

	for i := 1; i < len(nums); i++ {
		empty2[i] = max(empty2[i-1], full2[i-1]+nums[i])
		full2[i] = max(full2[i-1], empty1[i-1]-nums[i])
		empty1[i] = max(empty1[i-1], full1[i-1]+nums[i])
		full1[i] = max(full1[i-1], -nums[i])
	}

	return empty2[len(nums)-1]
}
func main() {
	fmt.Println(maxProfit([]int{3, 3, 5, 0, 0, 4, 1, 3}))
}
