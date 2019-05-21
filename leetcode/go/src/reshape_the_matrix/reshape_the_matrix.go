// 2017-12-14 16:21
package main

import "fmt"

func matrixReshape(matrix [][]int, r int, c int) [][]int {
	if len(matrix) == 0 {
		return matrix
	}
	if r*c != len(matrix)*len(matrix[0]) {
		return matrix
	}
	ret := make([][]int, r)
	for i := 0; i < r; i++ {
		ret[i] = make([]int, c)
	}
	x, y := 0, 0
	for i := 0; i < len(matrix); i++ {
		for j := 0; j < len(matrix[0]); j++ {
			ret[x][y] = matrix[i][j]
			y++
			if y >= c {
				x++
				y = 0
			}
		}
	}

	return ret
}
func main() {
	matrix := [][]int{
		{1, 2, 3},
	}
	fmt.Println(matrixReshape(matrix, 3, 1))
}
