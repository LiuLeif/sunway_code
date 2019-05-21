// 2017-12-21 15:30
package main

import "fmt"

type Pair struct {
	x, y int
}

func helper(matrix [][]int, target int, lo, hi Pair) bool {
	if lo.x > hi.x || lo.y > hi.y {
		return false
	}

	mid := Pair{(lo.x + hi.x) / 2, (lo.y + hi.y) / 2}

	if matrix[mid.x][mid.y] == target {
		return true
	}

	if matrix[mid.x][mid.y] > target {
		return helper(matrix, target, lo, Pair{mid.x - 1, hi.y}) || helper(matrix, target, lo, Pair{hi.x, mid.y - 1})
	} else {
		return helper(matrix, target, Pair{lo.x, mid.y + 1}, hi) || helper(matrix, target, Pair{mid.x + 1, lo.y}, hi)
	}
}

func searchMatrix(matrix [][]int, target int) bool {
	if len(matrix) == 0 {
		return false
	}
	return helper(matrix, target, Pair{0, 0}, Pair{len(matrix) - 1, len(matrix[0]) - 1})
}
func main() {
	matrix := [][]int{
		{1, 4, 7, 11, 15},
		{2, 5, 8, 12, 19},
		{3, 6, 9, 16, 22},
		{10, 13, 14, 17, 24},
		{18, 21, 23, 26, 30},
	}

	fmt.Println(searchMatrix(matrix, 0))
}
