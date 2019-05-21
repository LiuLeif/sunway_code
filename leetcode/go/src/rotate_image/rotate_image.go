// 2018-01-30 14:46
package main

import "fmt"

func rotate2(matrix [][]int) {
	//            line reverse    symmetric swap
	// 1 2 3      7 8 9           7 4 1
	// 4 5 6  ->  4 5 6  ->       8 5 2
	// 7 8 9      1 2 3           9 6 3
	n := len(matrix)
	for i := 0; i < n/2; i++ {
		matrix[i], matrix[n-i-1] = matrix[n-i-1], matrix[i]
	}
	for i := 0; i < n; i++ {
		for j := 0; j < i; j++ {
			matrix[i][j], matrix[j][i] = matrix[j][i], matrix[i][j]
		}
	}
}

func rotate(matrix [][]int) {
	x1, y1, x2, y2 := 0, 0, len(matrix)-1, len(matrix)-1
	for x1 < x2 {
		n := x2 - x1
		for i := 0; i < n; i++ {
			matrix[x1][y1+i], matrix[x1+i][y2], matrix[x2][y2-i], matrix[x2-i][y1] = matrix[x2-i][y1], matrix[x1][y1+i], matrix[x1+i][y2], matrix[x2][y2-i]
		}
		x1, y1, x2, y2 = x1+1, y1+1, x2-1, y2-1
	}
}
func main() {
	matrix := [][]int{
		{5, 1, 9, 11},
		{2, 4, 8, 10},
		{13, 3, 6, 7},
		{15, 14, 12, 16},
	}
	rotate2(matrix)
	fmt.Println(matrix)
}
