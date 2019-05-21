// 2018-01-05 15:34
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

type Item struct {
	weight, value int
}

func knapsack2(items []Item, weight int) int {
	nItems := len(items)
	dp1 := make([]int, weight+1)
	dp2 := make([]int, weight+1)

	// 若必须装满
	// for i := 1; i < weight+1; i++ {
	// 	dp1[i] = math.MinInt32
	// }

	for i := 1; i < nItems+1; i++ {
		for j := 1; j < weight+1; j++ {
			if j >= items[i-1].weight {
				dp2[j] = max(dp1[j], dp1[j-items[i-1].weight]+items[i-1].value)
			} else {
				dp2[j] = dp1[j]
			}
		}
		dp2, dp1 = dp1, dp2
	}
	return dp1[weight]
}

func knapsack(items []Item, weight int) int {
	nItems := len(items)
	dp := make([][]int, nItems+1)
	for i := 0; i < nItems+1; i++ {
		dp[i] = make([]int, weight+1)
		// 若必须装满:
		// for j := 1; j < weight+1; j++ {
		// 	dp[i][j] = math.MinInt32
		// }
	}
	for i := 1; i < nItems+1; i++ {
		for j := 1; j < weight+1; j++ {
			switch {
			case j >= items[i-1].weight:
				dp[i][j] = max(dp[i-1][j], dp[i-1][j-items[i-1].weight]+items[i-1].value)
			default:
				dp[i][j] = dp[i-1][j]
			}
		}
	}
	return dp[nItems][weight]
}

func main() {
	items := []Item{
		{1, 2},
		{2, 3},
		{4, 6},
	}
	fmt.Println(knapsack(items, 10))
}
