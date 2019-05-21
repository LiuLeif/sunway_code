// 2018-01-31 14:07
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

func minimumTotal(triangle [][]int) int {
	if len(triangle) == 0 {
		return 0
	}

	dp1 := make([]int, len(triangle[len(triangle)-1]))
	dp2 := make([]int, len(triangle[len(triangle)-1]))

	get := func(i, j int) int {
		if j < 0 || j > len(triangle[i])-1 {
			return math.MaxInt32
		}
		return dp1[j]
	}

	dp1[0] = triangle[0][0]

	for i := 1; i < len(triangle); i++ {
		for j := 0; j < i+1; j++ {
			dp2[j] = min(get(i-1, j), get(i-1, j-1)) + triangle[i][j]
		}
		dp1 = append([]int{}, dp2...)
	}

	return min(dp1...)
}
func main() {
	fmt.Println(minimumTotal([][]int{
		{2},
		{3, 4},
		{6, 5, 7},
		{4, 1, 8, 3},
	}))
}
