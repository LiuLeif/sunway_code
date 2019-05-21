// 2017-12-04 17:39
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

func findLength(A []int, B []int) int {
	dp := make([][]int, len(A)+1)
	for i := 0; i < len(A)+1; i++ {
		dp[i] = make([]int, len(B)+1)
	}
	ret := 0
	for i := 1; i < len(A)+1; i++ {
		for j := 1; j < len(B)+1; j++ {
			if A[i-1] == B[j-1] {
				dp[i][j] = dp[i-1][j-1] + 1
				ret = max(ret, dp[i][j])
			} else {
				dp[i][j] = 0
			}
		}
	}
	return ret
}
func main() {
	fmt.Println(findLength([]int{}, []int{3, 2, 1, 4, 7}))
}
