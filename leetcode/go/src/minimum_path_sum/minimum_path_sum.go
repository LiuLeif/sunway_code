// 2017-12-22 15:50
package main

import (
	"fmt"
	"math"
)

func min(nums ...int) int {
	ret := math.MaxInt32
	for _, n := range nums {
		if ret > n {
			ret = n
		}
	}
	return ret
}

type Pair struct {
	x, y int
}

func helper(m map[Pair]int, grid [][]int, p Pair) int {
	if p.x == len(grid)-1 && p.y == len(grid[0])-1 {
		return grid[p.x][p.y]
	}
	if m[p] != 0 {
		return m[p]
	}
	down := math.MaxInt32
	if p.x < len(grid)-1 {
		down = helper(m, grid, Pair{p.x + 1, p.y}) + grid[p.x][p.y]
	}
	right := math.MaxInt32
	if p.y < len(grid[0])-1 {
		right = helper(m, grid, Pair{p.x, p.y + 1}) + +grid[p.x][p.y]
	}
	m[p] = min(right, down)
	return m[p]
}

func minPathSum(grid [][]int) int {
	if len(grid) == 0 {
		return 0
	}
	m := make(map[Pair]int)
	return helper(m, grid, Pair{0, 0})
}
func main() {
	grid := [][]int{
		{1, 3, 1},
		{1, 2, 5},
		{4, 2, 1},
	}
	fmt.Println(minPathSum(grid))
}
