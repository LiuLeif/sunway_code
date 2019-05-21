// 2018-01-27 16:14
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

func maxCount(m int, n int, ops [][]int) int {
	if len(ops) == 0 {
		return m * n
	}
	row, col := math.MaxInt32, math.MaxInt32
	for _, op := range ops {
		row = min(row, op[0])
		col = min(col, op[1])
	}

	return row * col
}
func main() {
	fmt.Println(maxCount(3, 3, [][]int{
		{1, 1},
	}))
}
