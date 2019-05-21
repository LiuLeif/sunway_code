// 2018-04-08 11:18
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

func min(nums ...int) int {
	ret := math.MaxInt32
	for _, n := range nums {
		if ret > n {
			ret = n
		}
	}
	return ret
}

func maxIncreaseKeepingSkyline(grid [][]int) int {
	maxRow := make([]int, len(grid))
	maxCol := make([]int, len(grid[0]))
	for i := 0; i < len(grid); i++ {
		for j := 0; j < len(grid[0]); j++ {
			maxRow[i] = max(maxRow[i], grid[i][j])
			maxCol[j] = max(maxCol[j], grid[i][j])
		}
	}
	ret := 0
	for i := 0; i < len(grid); i++ {
		for j := 0; j < len(grid[0]); j++ {
			ret += min(maxRow[i], maxCol[j]) - grid[i][j]
		}
	}
	return ret
}
func main() {
	fmt.Println(maxIncreaseKeepingSkyline([][]int{
		{3, 0, 8, 4},
		{2, 4, 5, 7},
		{9, 2, 6, 3},
		{0, 3, 1, 0},
	}))
}
