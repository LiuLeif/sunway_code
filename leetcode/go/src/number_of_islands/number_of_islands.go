// 2017-11-30 11:48
package main

import "fmt"

func dfs(matrix [][]byte, i, j int) {
	matrix[i][j] = '0'
	if i > 0 && matrix[i-1][j] == '1' {
		dfs(matrix, i-1, j)
	}
	if j > 0 && matrix[i][j-1] == '1' {
		dfs(matrix, i, j-1)
	}
	if i < len(matrix)-1 && matrix[i+1][j] == '1' {
		dfs(matrix, i+1, j)
	}
	if j < len(matrix[0])-1 && matrix[i][j+1] == '1' {
		dfs(matrix, i, j+1)
	}
}

func numIslands(matrix [][]byte) int {
	ret := 0
	for i := 0; i < len(matrix); i++ {
		for j := 0; j < len(matrix[0]); j++ {
			if matrix[i][j] == '1' {
				dfs(matrix, i, j)
				ret++
			}
		}
	}
	return ret
}
func main() {
	matrix := [][]byte{
		{'1'},
		{'1'},
	}

	fmt.Println(numIslands(matrix))
}
