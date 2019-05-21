// 2017-11-14 18:54
package main

import "fmt"

func dfs(nums []int, sum int, start int) bool {
	fmt.Println("dfs", sum, start)
	if sum == 0 {
		return true
	}
	for i := start; i < len(nums); i++ {
		if sum-nums[i] < 0 {
			continue
		}
		ret := dfs(nums, sum-nums[i], i+1)
		if ret {
			return true
		}
	}
	return false
}

func canPartitionDFS(nums []int) bool {
	// did pass for test case:
	//
	//{1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1,
	//1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1,
	//1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1,
	//1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1,
	//1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 100}

	fmt.Println(len(nums))
	sum := 0
	for _, v := range nums {
		sum += v
	}
	if sum%2 != 0 {
		return false
	}
	return dfs(nums, sum/2, 0)
}

func canPartition(nums []int) bool {
	sum := 0
	for _, v := range nums {
		sum += v
	}
	if sum%2 != 0 {
		return false
	}
	sum /= 2

	dp := make([][]bool, len(nums)+1)
	for i := 0; i < len(nums)+1; i++ {
		dp[i] = make([]bool, sum+1)
	}

	dp[0][0] = true
	for i := 1; i < len(nums)+1; i++ {
		dp[i][0] = true
	}

	for i := 1; i < len(nums)+1; i++ {
		for j := 1; j < sum+1; j++ {
			v := nums[i-1]
			if j >= v {
				dp[i][j] = dp[i-1][j] || dp[i-1][j-v]
			} else {
				dp[i][j] = dp[i-1][j]
			}
		}
	}
	return dp[len(nums)][sum]
}

func main() {
	fmt.Println(canPartition([]int{1, 2, 1}))
	fmt.Println(canPartition(
		[]int{1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1,
			1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1,
			1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1,
			1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1,
			1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 100}))
}
