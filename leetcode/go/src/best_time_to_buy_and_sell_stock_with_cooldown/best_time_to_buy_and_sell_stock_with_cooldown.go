// 2017-11-21 17:48
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
	empty := make([]int, len(nums))
	full := make([]int, len(nums))
	empty[0] = 0
	full[0] = -nums[0]

	for i := 1; i < len(nums); i++ {
		empty[i] = max(empty[i-1], full[i-1]+nums[i])
		if i >= 2 {
			full[i] = max(full[i-1], empty[i-2]-nums[i])
		} else {
			full[i] = max(full[i-1], empty[i-1]-nums[i])
		}
	}

	return empty[len(nums)-1]
}

func main() {
	fmt.Println(maxProfit([]int{1, 2, 3, 1, 0, 2}))
}
