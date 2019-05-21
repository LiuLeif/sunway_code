// 2018-01-29 14:53
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

func deleteAndEarn(nums []int) int {
	// 0,0,2,3,4...
	N := make([]int, 10003)
	for i := 0; i < len(nums); i++ {
		N[nums[i]]++
	}
	dp := make([]int, 10003)
	for i := 2; i < 10003; i++ {
		dp[i] = max(N[i-2]*(i-2)+dp[i-2], dp[i-1])
	}
	return dp[10002]
}
func main() {
	fmt.Println(deleteAndEarn([]int{1, 1, 1, 1, 2, 3, 4}))
	fmt.Println(deleteAndEarn([]int{2, 2, 3, 3, 3, 4}))
}
