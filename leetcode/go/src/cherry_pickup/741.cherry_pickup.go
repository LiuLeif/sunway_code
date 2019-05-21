// 2018-12-13 18:43
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

func cherryPickup(grid [][]int) int {
	N := len(grid)
	cache := map[[4]int]int{}
	var dfs func(x1, y1, x2, y2 int) int
	dfs = func(x1, y1, x2, y2 int) int {
		if x1 == N-1 && y1 == N-1 && x2 == N-1 && y2 == N-1 {
			return grid[x1][y1]
		}
		if c, ok := cache[[4]int{x1, y1, x2, y2}]; ok {
			return c
		}
		ret := grid[x1][y1] + grid[x2][y2]
		if x1 == x2 && y1 == y2 {
			ret -= grid[x1][y1]
		}
		maxVal := -1
		for _, d1 := range [][]int{{1, 0}, {0, 1}} {
			x11, y11 := x1+d1[0], y1+d1[1]
			if x11 >= N || y11 >= N {
				continue
			}
			if grid[x11][y11] == -1 {
				continue
			}
			for _, d2 := range [][]int{{1, 0}, {0, 1}} {
				x22, y22 := x2+d2[0], y2+d2[1]
				if x22 >= N || y22 >= N {
					continue
				}
				if grid[x22][y22] == -1 {
					continue
				}
				tmp := dfs(x11, y11, x22, y22)
				maxVal = max(maxVal, tmp)
			}
		}
		if maxVal == -1 {
			ret = -1
		} else {
			ret += maxVal
		}
		cache[[4]int{x1, y1, x2, y2}] = ret
		return ret
	}

	ret := dfs(0, 0, 0, 0)
	if ret == -1 {
		return 0
	}
	return ret
}

func main() {
	fmt.Println(cherryPickup([][]int{
		{0, 1, 1, 0, 0},
		{1, 1, 1, 1, 0},
		{-1, 1, 1, 1, -1},
		{0, 1, 1, 1, 0},
		{1, 0, -1, 0, 0},
	}))
	fmt.Println(cherryPickup([][]int{
		{0, 1, -1},
		{1, 0, -1},
		{1, 1, 1},
	}))
}
