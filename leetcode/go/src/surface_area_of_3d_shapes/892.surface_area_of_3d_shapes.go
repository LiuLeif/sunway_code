// 2018-11-06 15:24
package main

import "fmt"

func surfaceArea(grid [][]int) int {
	ret := 0
	N := len(grid)
	for i := 0; i < N; i++ {
		for j := 0; j < N; j++ {
			if grid[i][j] == 0 {
				continue
			}
			// up-down
			ret += 2
			// left-right
			if j == 0 {
				ret += grid[i][j]
			} else if grid[i][j-1] < grid[i][j] {
				ret += grid[i][j] - grid[i][j-1]
			}
			if j == N-1 {
				ret += grid[i][j]
			} else if grid[i][j+1] < grid[i][j] {
				ret += grid[i][j] - grid[i][j+1]
			}
			// front-end
			if i == 0 {
				ret += grid[i][j]
			} else if grid[i-1][j] < grid[i][j] {
				ret += grid[i][j] - grid[i-1][j]
			}
			if i == N-1 {
				ret += grid[i][j]
			} else if grid[i+1][j] < grid[i][j] {
				ret += grid[i][j] - grid[i+1][j]
			}
		}
	}
	return ret
}

func main() {
	fmt.Println(surfaceArea([][]int{
		{1, 1, 1},
		{1, 0, 1},
		{1, 1, 1},
	}))
}
