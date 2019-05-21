// 2018-01-12 18:04
package main

import (
	"fmt"
	"math"
	"sort"
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

func maxEnvelopes(envelopes [][]int) int {
	if len(envelopes) == 0 {
		return 0
	}
	sort.Slice(envelopes, func(i, j int) bool {
		if envelopes[i][0] < envelopes[j][0] {
			return true
		}
		if envelopes[i][0] > envelopes[j][0] {
			return false
		}
		return envelopes[i][1] < envelopes[j][1]
	})
	dp := make([]int, len(envelopes))
	for i := 0; i < len(dp); i++ {
		dp[i] = 1
	}
	ret := 1
	for i := 1; i < len(envelopes); i++ {
		for j := 0; j < i; j++ {
			if envelopes[i][0] > envelopes[j][0] && envelopes[i][1] > envelopes[j][1] {
				dp[i] = max(dp[i], dp[j]+1)
			}
		}
		ret = max(dp[i], ret)
	}
	return ret
}
func main() {
	fmt.Println(maxEnvelopes([][]int{
		{5, 4},
		{6, 4},
		{6, 7},
		{2, 3},
	}))
}
