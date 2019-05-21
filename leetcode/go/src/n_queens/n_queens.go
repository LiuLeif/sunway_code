// 2017-12-20 17:40
package main

import "fmt"

func isValid(matrix [][]byte, x, y int) bool {
	n := len(matrix)
	// col
	for i := 0; i < n; i++ {
		if matrix[i][y] == 'Q' {
			return false
		}
	}
	// left-up
	for i := 1; i < n; i++ {
		x1, y1 := x-i, y-i
		if x1 < 0 || y1 < 0 {
			break
		}
		if matrix[x1][y1] == 'Q' {
			return false
		}
	}
	// left-down
	for i := 1; i < n; i++ {
		x1, y1 := x+i, y-i
		if x1 >= n || y1 < 0 {
			break
		}
		if matrix[x1][y1] == 'Q' {
			return false
		}
	}
	// right-up
	for i := 1; i < n; i++ {
		x1, y1 := x-i, y+i
		if x1 < 0 || y1 >= n {
			break
		}
		if matrix[x1][y1] == 'Q' {
			return false
		}
	}
	// right-down
	for i := 1; i < n; i++ {
		x1, y1 := x+i, y+i
		if x1 >= n || y1 >= n {
			break
		}
		if matrix[x1][y1] == 'Q' {
			return false
		}
	}
	return true
}
func dfs(ret *[][]string, matrix [][]byte, n int, current int) {
	if current == n {
		tmp := make([]string, n)
		for i := 0; i < n; i++ {
			tmp[i] = string(matrix[i])
		}
		*ret = append(*ret, tmp)
		return
	}
	for i := 0; i < n; i++ {
		if isValid(matrix, current, i) {
			matrix[current][i] = 'Q'
			dfs(ret, matrix, n, current+1)
			matrix[current][i] = '.'
		}
	}
}
func solveNQueens(n int) [][]string {
	ret := [][]string{}
	matrix := make([][]byte, n)
	for i := 0; i < n; i++ {
		matrix[i] = make([]byte, n)
		for j := 0; j < n; j++ {
			matrix[i][j] = '.'
		}
	}
	dfs(&ret, matrix, n, 0)
	return ret
}
func main() {
	fmt.Println(solveNQueens(4))
}
