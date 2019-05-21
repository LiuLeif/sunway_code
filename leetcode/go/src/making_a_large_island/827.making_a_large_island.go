// 2018-11-06 14:19
package main

import "fmt"

type Pair struct {
	a, b int
}

func largestIsland(grid [][]int) int {
	N := len(grid)
	color := 1
	var dfs func(x, y int) int
	dfs = func(x, y int) int {
		if grid[x][y] != 1 {
			return 0
		}
		grid[x][y] = color
		ret := 1
		for _, a := range []int{-1, 0, 1} {
			for _, b := range []int{-1, 0, 1} {
				if (a == 0 && b == 0) || a*b != 0 {
					continue
				}
				xx, yy := x+a, y+b
				if xx < 0 || xx >= N || yy < 0 || yy >= N {
					continue
				}
				ret += dfs(xx, yy)
			}
		}
		return ret
	}

	counter := map[int]int{}
	for i := 0; i < N; i++ {
		for j := 0; j < N; j++ {
			if grid[i][j] == 1 {
				color++
				counter[color] = dfs(i, j)
			}
		}
	}

	maxSum := 0
	for i := 0; i < N; i++ {
		for j := 0; j < N; j++ {
			if grid[i][j] != 0 {
				continue
			}
			sum := 1
			boundary := map[int]bool{}

			for _, a := range []int{-1, 0, 1} {
				for _, b := range []int{-1, 0, 1} {
					if (a == 0 && b == 0) || a*b != 0 {
						continue
					}
					x, y := i+a, j+b
					if x < 0 || x >= N || y < 0 || y >= N {
						continue
					}
					if grid[x][y] != 0 {
						boundary[grid[x][y]] = true
					}
				}
			}
			for k, _ := range boundary {
				sum += counter[k]
			}
			if sum > maxSum {
				maxSum = sum
			}
		}
	}
	if maxSum == 0 {
		return N * N
	}
	return maxSum
}

func main() {
	fmt.Println(largestIsland([][]int{
		{1, 1},
		{1, 1},
	}))
}
