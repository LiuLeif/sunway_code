// 2017-11-14 18:54
package main

import "fmt"

func findTargetSumWays(nums []int, S int) int {
	if len(nums) == 0 {
		if S == 0 {
			return 1
		} else {
			return 0
		}
	}
	x := nums[0]
	nums = nums[1:]
	return findTargetSumWays(nums, S-x) + findTargetSumWays(nums, S+x)
}

func dfs(nums []int, S int, start int) int {
	if start == len(nums) {
		if S == 0 {
			return 1
		} else {
			return 0
		}
	}
	return dfs(nums, S+nums[start], start+1) + dfs(nums, S-nums[start], start+1)
}

func findTargetSumWaysDFS(nums []int, S int) int {
	return dfs(nums, S, 0)
}

func findTargetSumWaysDP(nums []int, S int) int {
	total := 0
	for _, v := range nums {
		total += v
	}

	dp := make([]map[int]int, len(nums))
	for i := 0; i < len(nums); i++ {
		dp[i] = make(map[int]int)
	}

	dp[0][nums[0]] = 1
	dp[0][-nums[0]] += 1

	for i := 1; i < len(nums); i++ {
		for j := -total; j < total+1; j++ {
			dp[i][j] = dp[i-1][j+nums[i]] + dp[i-1][j-nums[i]]
		}
	}

	return dp[len(nums)-1][S]
}

func main() {
	fmt.Println(findTargetSumWaysDP([]int{0, 0, 0, 0, 0, 0, 0, 0, 1}, 1))
}
