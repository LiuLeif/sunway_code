// 2018-11-08 18:06
package main

import (
	"fmt"
	"math"
)

func getMoneyAmount(n int) int {
	// 1..10
	max := func(nums ...int) int {
		ret := math.MinInt32
		for _, n := range nums {
			if ret < n {
				ret = n
			}
		}
		return ret
	}

	min := func(nums ...int) int {
		ret := math.MaxInt32
		for _, n := range nums {
			if ret > n {
				ret = n
			}
		}
		return ret
	}

	memo := map[[2]int]int{}
	var dfs func(i, j int) int
	dfs = func(i, j int) int {
		if memo[[2]int{i, j}] != 0 {
			return memo[[2]int{i, j}]
		}
		if i >= j {
			return 0
		}
		if i+1 == j {
			return i
		}
		// 1 2 3 4
		// i     j
		memo[[2]int{i, j}] = math.MaxInt32
		for m := i + 1; m < j; m++ {
			memo[[2]int{i, j}] = min(memo[[2]int{i, j}], m+max(dfs(i, m-1), dfs(m+1, j)))
		}
		return memo[[2]int{i, j}]
	}
	return dfs(1, n)
}

func main() {
	// 1 2 3 4 5 6 7
	fmt.Println(getMoneyAmount(2))
}
