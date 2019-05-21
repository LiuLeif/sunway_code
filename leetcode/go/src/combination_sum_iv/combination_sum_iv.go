// 2017-11-15 15:08
package main

import "fmt"

func dfs(dp []int, nums []int, target int) int {
	if target == 0 {
		return 1
	}
	if dp[target] != -1 {
		return dp[target]
	}
	ret := 0
	for i := 0; i < len(nums); i++ {
		if target >= nums[i] {
			ret += dfs(dp, nums, target-nums[i])
		}
	}
	dp[target] = ret
	return ret
}

func combinationSum4(nums []int, target int) int {
	dp := make([]int, target+1)
	for i := 0; i < target+1; i++ {
		dp[i] = -1
	}

	dp[0] = 1
	return dfs(dp, nums, target)
}

func helper(dp []int, nums []int, target int) int {
	if dp[target] != -1 {
		return dp[target]
	}
	ret := 0
	for i := 0; i < len(nums); i++ {
		if nums[i] <= target {
			ret += helper(dp, nums, target-nums[i])
		}
	}
	dp[target] = ret
	return ret
}

func combinationSum4Recursive(nums []int, target int) int {
	dp := make([]int, target+1)
	for i := 0; i < target+1; i++ {
		dp[i] = -1
	}

	dp[0] = 1
	return helper(dp, nums, target)
}

func combinationSum4DP(nums []int, target int) int {
	dp := make([]int, target+1)
	dp[0] = 1
	for i := 1; i < target+1; i++ {
		for j := 0; j < len(nums); j++ {
			if i >= nums[j] {
				dp[i] += dp[i-nums[j]]
			}
		}
	}

	return dp[target]
}

func main() {
	fmt.Println(combinationSum4Recursive([]int{1, 2, 3}, 32))
}
