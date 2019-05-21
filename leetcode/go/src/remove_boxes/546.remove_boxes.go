// 2018-11-07 17:41
package main

import (
	"fmt"
	"math"
)

func removeBoxes(boxes []int) int {
	max := func(nums ...int) int {
		ret := math.MinInt32
		for _, n := range nums {
			if ret < n {
				ret = n
			}
		}
		return ret
	}

	N := len(boxes)

	memo := make([][][]int, 100)
	for i := 0; i < 100; i++ {
		memo[i] = make([][]int, 100)
		for j := 0; j < 100; j++ {
			memo[i][j] = make([]int, 100)
		}
	}
	var dfs func(i, j, k int) int
	dfs = func(i, j, k int) int {
		if i > j {
			return 0
		}
		if memo[i][j][k] != 0 {
			return memo[i][j][k]
		}
		// 1 1 1 2 1 3
		//     i   m j
		for i < j && boxes[i] == boxes[i+1] {
			i++
			k++
		}
		memo[i][j][k] = (k+1)*(k+1) + dfs(i+1, j, 0)
		for m := i + 1; m <= j; m++ {
			if boxes[m] == boxes[i] {
				memo[i][j][k] = max(memo[i][j][k], dfs(i+1, m-1, 0)+dfs(m, j, k+1))
			}
		}
		return memo[i][j][k]
	}
	return dfs(0, N-1, 0)
}

func main() {
	fmt.Println(removeBoxes([]int{1, 3, 2, 2, 2, 3, 4, 3, 1}))
}
