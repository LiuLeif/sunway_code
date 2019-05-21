// 2018-11-07 16:44
package main

import (
	"fmt"
	"math"
)

func strangePrinter(s string) int {
	if len(s) == 0 {
		return 0
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

	N := len(s)
	dp := make([][]int, N)
	for i := 0; i < N; i++ {
		dp[i] = make([]int, N)
		for k := 0; k < N; k++ {
			dp[i][k] = 1
		}
	}
	for i := 1; i < N; i++ {
		for j := 0; j < N-i; j++ {
			// j,i+j
			dp[j][i+j] = i + 1
			for k := j + 1; k <= i+j; k++ {
				tmp := dp[j][k-1] + dp[k][i+j]
				if s[j] == s[k] {
					tmp--
				}
				dp[j][i+j] = min(dp[j][i+j], tmp)
			}
		}
	}
	// fmt.Println(dp)
	return dp[0][N-1]
}

func main() {
	fmt.Println(strangePrinter("tbgtgb"))
	fmt.Println(strangePrinter("ccdcadbddbaddcbccdcdabcbcddbccdcbad"))
}
