// 2017-11-30 16:47
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

func minimumDeleteSum(s1 string, s2 string) int {
	dp := make([][]int, len(s1)+1)
	for i := 0; i < len(s1)+1; i++ {
		dp[i] = make([]int, len(s2)+1)
	}

	for i := 1; i < len(s1)+1; i++ {
		for j := 1; j < len(s2)+1; j++ {
			if s1[i-1] == s2[j-1] {
				dp[i][j] = dp[i-1][j-1] + int(s1[i-1])
			} else {
				dp[i][j] = max(dp[i-1][j], dp[i][j-1])
			}
		}
	}

	ret := 0
	for i := 0; i < len(s1); i++ {
		ret += int(s1[i])
	}
	for j := 0; j < len(s2); j++ {
		ret += int(s2[j])
	}

	ret -= 2 * dp[len(s1)][len(s2)]
	return ret
}
func main() {
	fmt.Println(minimumDeleteSum("delete", "leet"))
}
