// 2018-01-17 17:00
package main

import "fmt"

func dfs(grid [][]int, x, y int, from, to int) {
	if x < 0 || y < 0 || x >= len(grid) || y >= len(grid[0]) || grid[x][y] != from {
		return
	}
	grid[x][y] = to
	dfs(grid, x-1, y, from, to)
	dfs(grid, x+1, y, from, to)
	dfs(grid, x, y-1, from, to)
	dfs(grid, x, y+1, from, to)
}

func floodFill(image [][]int, sr int, sc int, newColor int) [][]int {
	if image[sr][sc] == newColor {
		return image
	}
	dfs(image, sr, sc, image[sr][sc], newColor)
	return image
}
func main() {
	fmt.Println(floodFill(
		[][]int{
			{0, 0, 0},
			{0, 1, 1},
		},
		1, 1, 1))
}
