// 2017-12-27 17:05
package main

import (
	"fmt"
	"math"
)

func min(nums ...int) int {
	ret := math.MaxInt32
	for _, n := range nums {
		if ret > n {
			ret = n
		}
	}
	return ret
}

func minCostClimbingStairs(cost []int) int {
	cost = append([]int{0}, cost...)
	cost = append(cost, 0)
	dp := make([]int, len(cost))
	dp[0] = 0
	dp[1] = 0
	for i := 2; i < len(cost); i++ {
		dp[i] = min(dp[i-1]+cost[i-1], dp[i-2]+cost[i-2])
	}
	return dp[len(dp)-1]
}
func main() {
	fmt.Println(minCostClimbingStairs([]int{1}))
}
