// 2018-01-29 11:16
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

func count(x string) (int, int) {
	zero, one := 0, 0
	for _, v := range x {
		if v == '0' {
			zero++
		} else {
			one++
		}
	}
	return zero, one
}

func findMaxForm(strs []string, m int, n int) int {
	dp := make([][][]int, len(strs)+1)
	for i := 0; i < len(strs)+1; i++ {
		dp[i] = make([][]int, m+1)
		for j := 0; j < m+1; j++ {
			dp[i][j] = make([]int, n+1)
		}
	}
	for i := 0; i < len(strs)+1; i++ {
		for j := 0; j < m+1; j++ {
			for k := 0; k < n+1; k++ {
				if i == 0 {
					dp[i][j][k] = 0
					continue
				}
				a, b := count(strs[i-1])
				if j >= a && k >= b {
					dp[i][j][k] = max(dp[i-1][j][k], dp[i-1][j-a][k-b]+1)
				} else {
					dp[i][j][k] = dp[i-1][j][k]
				}
			}
		}
	}
	return dp[len(strs)][m][n]
}
func main() {
	fmt.Println(findMaxForm([]string{"10", "0001", "111001", "1", "0"}, 5, 3))
	fmt.Println(findMaxForm([]string{"10", "0", "1"}, 1, 1))
}
