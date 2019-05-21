// 2017-11-14 18:54
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

func PredictTheWinner(nums []int) bool {
	m := make([][]int, len(nums))
	for i := 0; i < len(nums); i++ {
		m[i] = make([]int, len(nums))
	}
	for i := 0; i < len(nums); i++ {
		m[i][i] = nums[i]
	}

	for k := 1; k < len(nums); k++ {
		for j := 0; j < len(nums)-k; j++ {
			m[j][j+k] = max(nums[j]-m[j+1][j+k], nums[j+k]-m[j][j+k-1])
		}
	}
	return m[0][len(nums)-1] >= 0
}
func main() {
	fmt.Println(PredictTheWinner([]int{1, 5, 2}))
}
