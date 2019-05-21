// 2018-01-29 19:39
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

func partitionLabels(S string) []int {
	ret := []int{}
	last := map[byte]int{}
	for i := len(S) - 1; i >= 0; i-- {
		if last[S[i]] == 0 {
			last[S[i]] = i
		}
	}
	left, right := -1, 0
	for i := 0; i < len(S); i++ {
		right = max(right, last[S[i]])
		if i == right {
			ret = append(ret, right-left)
			left = i
			right = 0
		}
	}
	return ret
}
func main() {
	fmt.Println(partitionLabels("ababcbacadefegdehijhklij"))
}
