// 2017-11-23 14:32
package main

import "fmt"

func findDiagonalOrder(matrix [][]int) []int {
	if len(matrix) == 0 {
		return []int{}
	}

	ret := []int{}

	row := len(matrix)
	column := len(matrix[0])

	count := make([][]int, row+column)
	for i := 0; i < row+column; i++ {
		count[i] = []int{}
	}
	for i := 0; i < row; i++ {
		for j := 0; j < column; j++ {
			if (i+j)&1 == 0 {
				count[i+j] = append([]int{matrix[i][j]}, count[i+j]...)
			} else {
				count[i+j] = append(count[i+j], matrix[i][j])
			}
		}
	}

	for i := 0; i < row+column-1; i++ {
		ret = append(ret, count[i]...)
	}

	return ret
}
func main() {
	matrix := [][]int{
		{1, 2, 3},
		{4, 5, 6},
		{7, 8, 9},
	}
	fmt.Println(findDiagonalOrder(matrix))
}
