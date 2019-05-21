// 2017-11-22 11:04
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

func maxProfit(k int, nums []int) int {
	if len(nums) == 0 {
		return 0
	}

	if k >= len(nums)/2 {
		ret := 0
		for i := 0; i < len(nums)-1; i++ {
			if nums[i+1] > nums[i] {
				ret += nums[i+1] - nums[i]
			}
		}

		return ret
	}
	full := make([][]int, k+1)
	for i := 0; i < k+1; i++ {
		full[i] = make([]int, len(nums))
	}
	empty := make([][]int, k+1)
	for i := 0; i < k+1; i++ {
		empty[i] = make([]int, len(nums))
	}

	for i := 0; i < k+1; i++ {
		full[i][0] = -nums[0]
	}

	for i := 1; i < len(nums); i++ {
		for j := 1; j < k+1; j++ {
			empty[j][i] = max(empty[j][i-1], full[j][i-1]+nums[i])
			full[j][i] = max(full[j][i-1], empty[j-1][i-1]-nums[i])
		}
	}

	return empty[k][len(nums)-1]
}
func main() {
	fmt.Println(maxProfit(2, []int{3, 3, 4, 8}))
}
