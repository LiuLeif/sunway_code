// 2018-10-17 15:19
package main

import (
	"fmt"
	"math"
)

func largestOverlap(A [][]int, B [][]int) int {
	N := len(A)
	ret := 0
	max := func(nums ...int) int {
		ret := math.MinInt32
		for _, n := range nums {
			if ret < n {
				ret = n
			}
		}
		return ret
	}

	for i := 0; i < N; i++ {
		for j := 0; j < N; j++ {
			for _, k := range []int{i, -i} {
				for _, v := range []int{j, -j} {
					count := 0
					for x := 0; x < N; x++ {
						for y := 0; y < N; y++ {
							a, b := x+k, y+v
							if a < 0 || a > N-1 || b < 0 || b > N-1 {
								continue
							}
							if A[x][y] == 1 && B[a][b] == 1 {
								count++
							}
						}
					}
					ret = max(ret, count)
				}
			}
		}
	}
	return ret
}

func main() {
	fmt.Println(largestOverlap(
		[][]int{
			{0, 0, 0, 0, 0},
			{0, 1, 0, 0, 0},
			{0, 0, 1, 0, 0},
			{0, 0, 0, 0, 1},
			{0, 1, 0, 0, 1},
		},
		[][]int{
			{1, 0, 1, 1, 1},
			{1, 1, 1, 1, 1},
			{1, 1, 1, 1, 1},
			{0, 1, 1, 1, 1},
			{1, 0, 1, 1, 1},
		}))
}
