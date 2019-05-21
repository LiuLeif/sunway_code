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

func findNumberOfLIS(nums []int) int {
	if len(nums) == 0 {
		return 0
	}
	if len(nums) == 1 {
		return 1
	}
	dp := make([]int, len(nums))
	count := make([]int, len(nums))

	for i := 0; i < len(nums); i++ {
		dp[i] = 1
		count[i] = 1
	}

	maxLength := 0
	for i := 1; i < len(nums); i++ {
		for j := 0; j < i; j++ {
			if nums[i] > nums[j] {
				if dp[j]+1 > dp[i] {
					dp[i] = dp[j] + 1
					count[i] = count[j]
				} else if dp[j]+1 == dp[i] {
					count[i] += count[j]
				}
			} else if nums[i] == nums[j] {
				if dp[j] > dp[i] {
					dp[i] = dp[j]
					count[i] = count[j]
				} else if dp[j] == dp[i] {
					count[i] = count[j]
				}
			}
			maxLength = max(maxLength, dp[i])
		}
	}

	ret := 0
	for i := 0; i < len(nums); i++ {
		if dp[i] == maxLength {
			ret += count[i]
		}
	}

	return ret
}
func main() {
	fmt.Println(findNumberOfLIS([]int{1, 2}))
}
