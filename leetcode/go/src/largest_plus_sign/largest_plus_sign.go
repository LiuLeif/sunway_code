// 2018-10-18 16:06
package main

import (
	"fmt"
	"math"
)

func orderOfLargestPlusSign(N int, mines [][]int) int {
	grid := make([][]int, N)
	for i := 0; i < N; i++ {
		grid[i] = make([]int, N)
		for j := 0; j < N; j++ {
			grid[i][j] = 1
		}
	}
	for _, v := range mines {
		grid[v[0]][v[1]] = 0
	}
	up := make([][]int, N)
	for i := 0; i < N; i++ {
		up[i] = make([]int, N)
	}
	down := make([][]int, N)
	for i := 0; i < N; i++ {
		down[i] = make([]int, N)
	}
	left := make([][]int, N)
	for i := 0; i < N; i++ {
		left[i] = make([]int, N)
	}
	right := make([][]int, N)
	for i := 0; i < N; i++ {
		right[i] = make([]int, N)
	}
	for i := 0; i < N; i++ {
		for j := 0; j < N; j++ {
			if grid[i][j] == 1 {
				up[i][j] = 1
				if i > 0 {
					up[i][j] += up[i-1][j]
				}
				left[i][j] = 1
				if j > 0 {
					left[i][j] += left[i][j-1]
				}
			}
		}
	}
	for i := N - 1; i >= 0; i-- {
		for j := N - 1; j >= 0; j-- {
			if grid[i][j] == 1 {
				down[i][j] = 1
				if i < N-1 {
					down[i][j] += down[i+1][j]
				}
				right[i][j] = 1
				if j < N-1 {
					right[i][j] += right[i][j+1]
				}
			}
		}
	}
	min := func(nums ...int) int {
		ret := math.MaxInt32
		for _, n := range nums {
			if ret > n {
				ret = n
			}
		}
		return ret
	}
	max := func(nums ...int) int {
		ret := math.MinInt32
		for _, n := range nums {
			if ret < n {
				ret = n
			}
		}
		return ret
	}

	ret := 0
	for i := 0; i < N; i++ {
		for j := 0; j < N; j++ {
			ret = max(ret, min(up[i][j], down[i][j], left[i][j], right[i][j]))
		}
	}
	return ret
}
func main() {
	fmt.Println(orderOfLargestPlusSign(5, [][]int{
		{4, 2},
	}))
}
