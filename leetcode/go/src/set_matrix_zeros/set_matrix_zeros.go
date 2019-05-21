// 2017-11-14 18:54
package main

import "fmt"

func setZeroes(matrix [][]int) {
	row := len(matrix)
	column := len(matrix[0])
	markRow, markColumn := -1, -1

	for i := 0; i < row; i++ {
		for j := 0; j < column; j++ {
			if matrix[i][j] == 0 {
				if markRow == -1 {
					markRow, markColumn = i, j
				} else {
					matrix[markRow][j] = 0
					matrix[i][markColumn] = 0
				}
			}
		}
	}

	if markRow == -1 {
		return
	}
	for i := 0; i < row; i++ {
		if i == markRow {
			continue
		}
		if matrix[i][markColumn] == 0 {
			for j := 0; j < column; j++ {
				matrix[i][j] = 0
			}
		}
	}
	for j := 0; j < column; j++ {
		if j == markColumn {
			continue
		}
		if matrix[markRow][j] == 0 {
			for i := 0; i < row; i++ {
				matrix[i][j] = 0
			}
		}
	}
	for i := 0; i < row; i++ {
		matrix[i][markColumn] = 0
	}

	for i := 0; i < column; i++ {
		matrix[markRow][i] = 0
	}

}

func main() {
	matrix := make([][]int, 1)
	matrix[0] = []int{}

	setZeroes(matrix)
	fmt.Println(matrix)
}
