// 2017-12-21 15:30
package main

import "fmt"

type Pair struct {
	x, y int
}

func helper(matrix [][]int, m map[Pair]int, x, y int) int {
	lx := len(matrix) - 1
	ly := len(matrix[0]) - 1
	if matrix[x][y] == 1 {
		return 0
	}
	if x == lx && y == ly {
		return 1
	}
	cached, found := m[Pair{x, y}]
	if found {
		return cached
	}
	right, down := 0, 0
	if y < ly && matrix[x][y+1] != 1 {
		right = helper(matrix, m, x, y+1)
	}
	if x < lx && matrix[x+1][y] != 1 {
		down = helper(matrix, m, x+1, y)
	}
	m[Pair{x, y}] = right + down
	return right + down
}

func uniquePathsWithObstacles(matrix [][]int) int {
	m := make(map[Pair]int)
	return helper(matrix, m, 0, 0)
}
func main() {
	matrix := [][]int{
		{0, 1},
	}
	fmt.Println(uniquePathsWithObstacles(matrix))
}
