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

func lengthOfLIS(nums []int) int {
	if len(nums) == 0 {
		return 0
	}
	dp := make([]int, len(nums))
	for i := 0; i < len(nums); i++ {
		dp[i] = 1
	}

	ret := 1
	for i := 1; i < len(nums); i++ {
		for j := 0; j < i; j++ {
			if nums[i] > nums[j] {
				dp[i] = max(dp[i], dp[j]+1)
			} else if nums[i] == nums[j] {
				dp[i] = max(dp[i], dp[j])
			}
		}
		ret = max(ret, dp[i])
	}
	return ret
}
func main() {
	fmt.Println(lengthOfLIS([]int{1, 2, 3, 4, 1, 2, 3, 4, 5}))
}
