// 2017-12-14 16:21
package main

import "fmt"

type NumMatrix struct {
	matrix [][]int
}

func Constructor(matrix [][]int) NumMatrix {
	ret := NumMatrix{matrix}
	for i := 0; i < len(matrix); i++ {
		for j := 1; j < len(matrix[0]); j++ {
			ret.matrix[i][j] += ret.matrix[i][j-1]
		}
	}
	for i := 1; i < len(matrix); i++ {
		for j := 0; j < len(matrix[0]); j++ {
			ret.matrix[i][j] += ret.matrix[i-1][j]
		}
	}
	return ret
}

func (this *NumMatrix) SumRegion(row1 int, col1 int, row2 int, col2 int) int {
	up := 0
	if row1 > 0 {
		up = this.matrix[row1-1][col2]
	}
	left := 0
	if col1 > 0 {
		left = this.matrix[row2][col1-1]
	}
	upLeft := 0
	if row1 > 0 && col1 > 0 {
		upLeft = this.matrix[row1-1][col1-1]
	}
	return this.matrix[row2][col2] - up - left + upLeft
}

func main() {
	matrix := Constructor([][]int{
		{3, 0, 1, 4, 2},
		{5, 6, 3, 2, 1},
		{1, 2, 0, 1, 5},
		{4, 1, 0, 1, 7},
		{1, 0, 3, 0, 5},
	})
	fmt.Println(matrix.SumRegion(2, 1, 4, 3))
	fmt.Println(matrix.SumRegion(0, 0, 0, 3))
}
