// 2018-10-29 12:45
package main

import (
	"fmt"
	"math"
)

func minFallingPathSum(A [][]int) int {
	row, col := len(A), len(A[0])
	dp := make([]int, col)
	dpTmp := make([]int, col)
	min := func(nums ...int) int {
		ret := math.MaxInt32
		for _, n := range nums {
			if ret > n {
				ret = n
			}
		}
		return ret
	}

	get := func(x int) int {
		if x < 0 || x >= col {
			return math.MaxInt32
		}
		return dp[x]
	}

	for i := 0; i < row; i++ {
		for j := 0; j < col; j++ {
			if i == 0 {
				dpTmp[j] = A[i][j]
				continue
			}
			dpTmp[j] = min(get(j-1), get(j), get(j+1)) + A[i][j]
		}
		copy(dp, dpTmp)
	}
	return min(dp...)
}

func main() {
	fmt.Println(minFallingPathSum([][]int{
		{-19, 57},
		{-40, -5},
	}))
	fmt.Println(minFallingPathSum([][]int{
		{1, 2, 3},
		{4, 5, 6},
		{7, 8, 9},
	}))
}
