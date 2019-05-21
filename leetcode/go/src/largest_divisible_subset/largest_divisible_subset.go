// 2018-01-08 14:51
package main

import (
	"fmt"
	"sort"
)

func largestDivisibleSubset(nums []int) []int {
	ret := []int{}
	if len(nums) == 0 {
		return ret
	}
	sort.Ints(nums)
	dp := make([][]int, len(nums))
	maxLen := 0
	for i := 0; i < len(nums); i++ {
		dp[i] = []int{nums[i]}
		if len(dp[i]) > maxLen {
			ret = dp[i]
			maxLen = len(dp[i])
		}
		for j := 0; j < i; j++ {
			if nums[i]%nums[j] == 0 {
				dp[i] = append([]int{nums[i]}, dp[j]...)
				if len(dp[i]) > maxLen {
					ret = dp[i]
					maxLen = len(dp[i])
				}
			}
		}
	}
	return ret
}

func main() {
	fmt.Println(largestDivisibleSubset([]int{1, 2, 4, 8, 16}))
	fmt.Println(largestDivisibleSubset([]int{3, 4, 8, 16}))
}
