// 2018-10-18 17:58
package main

import (
	"fmt"
	"math"
)

func lenLongestFibSubseq(A []int) int {
	dp := make([][]int, len(A))
	for i := 0; i < len(A); i++ {
		dp[i] = make([]int, len(A))
		for j := i + 1; j < len(A); j++ {
			dp[i][j] = 2
		}
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
	for i := 2; i < len(A); i++ {
		a, b := 0, i-1
		for a < b {
			if A[a]+A[b] < A[i] {
				a++
			} else if A[a]+A[b] > A[i] {
				b--
			} else {
				dp[b][i] = max(dp[b][i], dp[a][b]+1)
				ret = max(ret, dp[b][i])
				a++
				b--
			}
		}
	}
	return ret
}

func main() {
	fmt.Println(lenLongestFibSubseq([]int{1, 3, 7, 11, 12, 14, 18}))
	fmt.Println(lenLongestFibSubseq([]int{1, 2, 3, 4, 5, 6, 7, 8}))
}
