// 2017-11-14 18:54
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

func dfs(matrix [][]int, dp [][]int, i, j int, visited [][]bool) int {
	row := len(matrix)
	column := len(matrix[0])

	if visited[i][j] == true {
		return dp[i][j]
	}

	visited[i][j] = true

	if i > 0 && matrix[i][j] < matrix[i-1][j] {
		dp[i][j] = max(dp[i][j], dfs(matrix, dp, i-1, j, visited))
	}
	if i < row-1 && matrix[i][j] < matrix[i+1][j] {
		dp[i][j] = max(dp[i][j], dfs(matrix, dp, i+1, j, visited))
	}
	if j > 0 && matrix[i][j] < matrix[i][j-1] {
		dp[i][j] = max(dp[i][j], dfs(matrix, dp, i, j-1, visited))
	}
	if j < column-1 && matrix[i][j] < matrix[i][j+1] {
		dp[i][j] = max(dp[i][j], dfs(matrix, dp, i, j+1, visited))
	}
	dp[i][j] += 1

	return dp[i][j]
}

func longestIncreasingPath(matrix [][]int) int {
	row := len(matrix)
	if row == 0 {
		return 0
	}
	column := len(matrix[0])
	if column == 0 {
		return 0
	}

	dp := make([][]int, row)
	for i := 0; i < row; i++ {
		dp[i] = make([]int, column)
	}

	visited := make([][]bool, row)
	for i := 0; i < row; i++ {
		visited[i] = make([]bool, column)
		for j := 0; j < column; j++ {
			visited[i][j] = false
		}
	}

	ret := 0
	for i := 0; i < row; i++ {
		for j := 0; j < column; j++ {
			ret = max(dfs(matrix, dp, i, j, visited), ret)
		}
	}

	return ret
}
func main() {
	matrix := [][]int{
		[]int{9, 9, 4},
		[]int{6, 6, 8},
		[]int{2, 1, 1},
	}
	fmt.Println(longestIncreasingPath(matrix))
}
