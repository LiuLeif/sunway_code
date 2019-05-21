// 2018-01-05 16:11
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

func knapsack(items []Item, maxWeight int) int {
	dp := make([]int, maxWeight+1)
	// 若必须装满
	// for i := 1; i < maxWeight+1; i++ {
	// 	dp[i] = math.MinInt32
	// }
	for w := 1; w < maxWeight+1; w++ {
		for i := 0; i < len(items); i++ {
			if w >= items[i].weight {
				dp[w] = max(dp[w], dp[w-items[i].weight]+items[i].value)
			}
		}
	}
	return dp[maxWeight]
}

func main() {
	items := []Item{
		{4, 9},
		{3, 10},
	}
	fmt.Println(knapsack(items, 8))
}
