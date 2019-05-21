// 2018-01-17 15:42
package main

import "fmt"

func islandPerimeter(grid [][]int) int {
	ret := 0
	for i := 0; i < len(grid); i++ {
		for j := 0; j < len(grid[0]); j++ {
			if grid[i][j] == 0 {
				continue
			}
			if i == 0 || grid[i-1][j] == 0 {
				ret++
			}
			if i == len(grid)-1 || grid[i+1][j] == 0 {
				ret++
			}
			if j == 0 || grid[i][j-1] == 0 {
				ret++
			}
			if j == len(grid[0])-1 || grid[i][j+1] == 0 {
				ret++
			}
		}
	}

	return ret
}

func main() {
	fmt.Println(islandPerimeter([][]int{
		{1, 1},
		{1, 1},
	}))
}
