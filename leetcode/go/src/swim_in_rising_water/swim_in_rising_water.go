// 2018-04-13 13:50
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

func min(nums ...int) int {
	ret := math.MaxInt32
	for _, n := range nums {
		if ret > n {
			ret = n
		}
	}
	return ret
}

func swimInWaterBinarySearch(grid [][]int) int {
	maxValue := math.MinInt32
	for i := 0; i < len(grid); i++ {
		maxValue = max(maxValue, max(grid[i]...))
	}

	visited := make([][]bool, len(grid))

	var dfs func(x, y, time int)
	dfs = func(x, y, time int) {
		if x < 0 || y < 0 || x >= len(grid) || y >= len(grid[0]) || visited[x][y] || grid[x][y] > time {
			return
		}
		visited[x][y] = true
		dfs(x+1, y, time)
		dfs(x-1, y, time)
		dfs(x, y+1, time)
		dfs(x, y-1, time)
	}

	check := func(time int) bool {
		for i := 0; i < len(grid); i++ {
			visited[i] = make([]bool, len(grid[0]))
		}
		dfs(0, 0, time)
		return visited[len(grid)-1][len(grid[0])-1]
	}

	left := 0
	right := maxValue
	ret := 0
	for left <= right {
		mid := left + (right-left)/2
		if !check(mid) {
			left = mid + 1
		} else {
			ret = mid
			right = mid - 1
		}
	}

	return ret
}

func swimInWater(grid [][]int) int {
	visited := make([][]bool, len(grid))
	for i := 0; i < len(grid); i++ {
		visited[i] = make([]bool, len(grid[0]))
	}

	var dfs func(x, y, time int)
	dfs = func(x, y, time int) {
		if x < 0 || y < 0 || x >= len(grid) || y >= len(grid[0]) || visited[x][y] || grid[x][y] > time {
			return
		}
		visited[x][y] = true
		dfs(x+1, y, time)
		dfs(x-1, y, time)
		dfs(x, y+1, time)
		dfs(x, y-1, time)
	}
	time := 0
	for !visited[len(grid)-1][len(grid[0])-1] {
		for i := 0; i < len(grid); i++ {
			visited[i] = make([]bool, len(grid[0]))
		}
		time++
		dfs(0, 0, time)
	}
	return time
}

func swimInWaterBacktracking(grid [][]int) int {
	visited := make([][]bool, len(grid))
	for i := 0; i < len(grid); i++ {
		visited[i] = make([]bool, len(grid[0]))
	}

	var dfs func(x, y int) int
	dfs = func(x, y int) int {
		if (x == len(grid)-1) && (y == len(grid[0])-1) {
			return grid[x][y]
		}
		if x < 0 || y < 0 || x >= len(grid) || y >= len(grid[0]) {
			return math.MaxInt32
		}

		if visited[x][y] {
			return math.MaxInt32
		}
		visited[x][y] = true
		ret := math.MaxInt32
		ret = min(ret, dfs(x+1, y), dfs(x-1, y), dfs(x, y+1), dfs(x, y-1))
		ret = max(ret, grid[x][y])
		visited[x][y] = false
		return ret
	}
	ret := dfs(0, 0)
	return ret
}

func main() {
	fmt.Println(swimInWaterBinarySearch([][]int{
		{7, 34, 16, 12, 15, 0},
		{10, 26, 4, 30, 1, 20},
		{28, 27, 33, 35, 3, 8},
		{29, 9, 13, 14, 11, 32},
		{31, 21, 23, 24, 19, 18},
		{22, 6, 17, 5, 2, 25},
	}))

	fmt.Println(swimInWaterBinarySearch([][]int{
		{0, 2},
		{1, 3},
	}))

	fmt.Println(swimInWaterBinarySearch([][]int{
		{0, 1, 2, 3, 4},
		{24, 23, 22, 21, 5},
		{12, 13, 14, 15, 16},
		{11, 17, 18, 19, 20},
		{10, 9, 8, 7, 6},
	}))
}
