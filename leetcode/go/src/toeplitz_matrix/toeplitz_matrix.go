// 2018-02-12 16:51
package main

import "fmt"

func isToeplitzMatrix(matrix [][]int) bool {
	if len(matrix) == 0 {
		return true
	}
	row := len(matrix)
	col := len(matrix[0])

	for i := 0; i < row-1; i++ {
		for j := 0; j < col-1; j++ {
			if matrix[i][j] != matrix[i+1][j+1] {
				return false
			}
		}
	}
	return true
}
func main() {
	fmt.Println(isToeplitzMatrix(
		[][]int{
			{1, 2, 3, 4},
			{5, 1, 2, 3},
			{9, 5, 1, 2},
		},
	))
}
