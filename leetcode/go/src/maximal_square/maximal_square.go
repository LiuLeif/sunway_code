// 2017-12-25 17:35
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

func max(nums ...int) int {
	ret := math.MinInt32
	for _, n := range nums {
		if ret < n {
			ret = n
		}
	}
	return ret
}

func maximalSquare(matrix [][]byte) int {
	if len(matrix) == 0 {
		return 0
	}
	dp := make([][]int, len(matrix)+1)
	for i := 0; i < len(matrix)+1; i++ {
		dp[i] = make([]int, len(matrix[0])+1)
	}
	ret := 0
	for i := 1; i < len(matrix)+1; i++ {
		for j := 1; j < len(matrix[0])+1; j++ {
			if matrix[i-1][j-1] == '1' {
				dp[i][j] = min(dp[i-1][j-1], dp[i-1][j], dp[i][j-1]) + 1
				ret = max(ret, dp[i][j])
			}
		}
	}
	return ret * ret
}
func main() {
	matrix := [][]byte{
		{'1', '0', '1', '0', '0'},
		{'1', '0', '1', '1', '1'},
		{'1', '1', '1', '1', '1'},
		{'1', '0', '0', '1', '0'},
	}
	fmt.Println(maximalSquare(matrix))
}
