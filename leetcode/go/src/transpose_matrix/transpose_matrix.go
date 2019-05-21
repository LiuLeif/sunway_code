// 2018-10-15 09:59
package main

import "fmt"

func transpose(A [][]int) [][]int {
	if len(A) == 0 {
		return A
	}
	row, col := len(A), len(A[0])
	ret := make([][]int, col)
	for i := 0; i < col; i++ {
		ret[i] = make([]int, row)
	}
	for i := 0; i < col; i++ {
		for j := 0; j < row; j++ {
			ret[i][j] = A[j][i]
		}
	}
	return ret
}
func main() {
	fmt.Println(transpose([][]int{{1, 2, 3}, {4, 5, 6}}))
}
