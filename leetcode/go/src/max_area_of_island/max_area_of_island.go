// 2017-11-14 18:54
package main

import (
	"fmt"
	"math"
)

func max(nums ...int) int {
	ret := math.MinInt32
	for _, n := range nums {
		if ret < n {
			ret = n
		}
	}
	return ret
}

func dfs(grid [][]int, x, y int) int {
	ret := 1
	grid[x][y] = 0

	row := len(grid)
	column := len(grid[0])

	if x != 0 && grid[x-1][y] == 1 {
		ret += dfs(grid, x-1, y)
	}
	if x != row-1 && grid[x+1][y] == 1 {
		ret += dfs(grid, x+1, y)
	}
	if y != 0 && grid[x][y-1] == 1 {
		ret += dfs(grid, x, y-1)
	}
	if y != column-1 && grid[x][y+1] == 1 {
		ret += dfs(grid, x, y+1)
	}
	return ret
}

func maxAreaOfIsland(grid [][]int) int {
	ret := 0
	row := len(grid)
	column := len(grid[0])
	for i := 0; i < row; i++ {
		for j := 0; j < column; j++ {
			if grid[i][j] == 1 {
				ret = max(ret, dfs(grid, i, j))
			}
		}
	}

	return ret
}

func maxAreaOfIsland(grid [][]int) int {
	ret := 0
	row := len(grid)
	column := len(grid[0])
	for i := 0; i < row; i++ {
		for j := 0; j < column; j++ {
			if grid[i][j] == 1 {
				grid[i][j] = 0
				queue := [][]int{{i, j}}
				count := 0
				for len(queue) != 0 {
					head := queue[0]
					queue = queue[1:]
					count++

					x := head[0]
					y := head[1]

					if x != 0 && grid[x-1][y] == 1 {
						queue = append(queue, []int{x - 1, y})
						grid[x-1][y] = 0
					}
					if x != row-1 && grid[x+1][y] == 1 {
						queue = append(queue, []int{x + 1, y})
						grid[x+1][y] = 0
					}
					if y != 0 && grid[x][y-1] == 1 {
						queue = append(queue, []int{x, y - 1})
						grid[x][y-1] = 0
					}
					if y != column-1 && grid[x][y+1] == 1 {
						queue = append(queue, []int{x, y + 1})
						grid[x][y+1] = 0
					}
				}
				ret = max(ret, count)
			}
		}
	}

	return ret
}
func main() {
	// grid := make([][]int, 8)
	// grid[0] = []int{0, 0, 1, 0, 0, 0, 0, 1, 0, 0, 0, 0, 0}
	// grid[1] = []int{0, 0, 0, 0, 0, 0, 0, 1, 1, 1, 0, 0, 0}
	// grid[2] = []int{0, 1, 1, 0, 1, 0, 0, 0, 0, 0, 0, 0, 0}
	// grid[3] = []int{0, 1, 0, 0, 1, 1, 0, 0, 1, 0, 1, 0, 0}
	// grid[4] = []int{0, 1, 0, 0, 1, 1, 0, 0, 1, 1, 1, 0, 0}
	// grid[5] = []int{0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 1, 0, 0}
	// grid[6] = []int{0, 0, 0, 0, 0, 0, 0, 1, 1, 1, 0, 0, 0}
	// grid[7] = []int{0, 0, 0, 0, 0, 0, 0, 1, 1, 0, 0, 0, 0}

	grid := make([][]int, 2)
	grid[0] = []int{1, 1}
	grid[1] = []int{1, 1}
	fmt.Println(maxAreaOfIslandDFS(grid))
}
