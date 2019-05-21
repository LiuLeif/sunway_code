// 2018-01-18 11:43
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

func maximalRectangle(matrix [][]byte) int {
	if len(matrix) == 0 {
		return 0
	}
	row, col := len(matrix), len(matrix[0])
	dpHeight := make([][]int, row+1)
	dpWidth := make([][]int, row+1)
	for i := 0; i < row+1; i++ {
		dpHeight[i] = make([]int, col+1)
		dpWidth[i] = make([]int, col+1)
	}
	for i := 1; i < row+1; i++ {
		for j := 1; j < col+1; j++ {
			if matrix[i-1][j-1] == '0' {
				dpHeight[i][j] = 0
				dpWidth[i][j] = 0
			} else {
				dpHeight[i][j] = dpHeight[i-1][j] + 1
				dpWidth[i][j] = dpWidth[i][j-1] + 1
			}
		}
	}
	ret := 0
	for i := 1; i < row+1; i++ {
		for j := 1; j < col+1; j++ {
			minHeight := math.MaxInt32
			currCol := j
			for ; currCol >= 1; currCol-- {
				if matrix[i-1][currCol-1] == '0' {
					break
				}
				minHeight = min(minHeight, dpHeight[i][currCol])
				ret = max(ret, minHeight*(j-currCol+1))
			}
		}
	}
	return ret
}
func main() {
	fmt.Println(maximalRectangle([][]byte{
		{'1', '1', '1', '0', '0'},
		{'1', '1', '0', '1', '1'},
		{'1', '1', '1', '1', '1'},
		{'1', '1', '0', '1', '0'},
	}))
}
