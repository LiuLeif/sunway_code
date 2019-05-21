// 2018-10-18 09:36
package main

import (
	"fmt"
	"math"
)

func stoneGame(piles []int) bool {
	sumVal := make([]int, len(piles)+1)
	for i := 0; i < len(piles); i++ {
		sumVal[i+1] = sumVal[i] + piles[i]
	}

	dp := make([][]int, len(piles))
	for i := 0; i < len(piles); i++ {
		dp[i] = make([]int, len(piles))
		dp[i][i] = piles[i]
	}

	sum := func(i, j int) int {
		return sumVal[j+1] - sumVal[i]
	}

	max := func(nums ...int) int {
		ret := math.MinInt32
		for _, n := range nums {
			if ret < n {
				ret = n
			}
		}
		return ret
	}

	for step := 1; step < len(piles); step++ {
		for i := 0; i < len(piles)-step; i++ {
			j := i + step
			dp[i][j] = max(sum(i, j-1)-dp[i][j-1]+piles[j], sum(i+1, j)-dp[i+1][j]+piles[i])
		}
	}

	if dp[0][len(piles)-1] > sum(0, len(piles)-1)-dp[0][len(piles)-1] {
		return true
	}
	return false
}
func main() {
	fmt.Println(stoneGame([]int{5, 1, 6, 1}))
}
