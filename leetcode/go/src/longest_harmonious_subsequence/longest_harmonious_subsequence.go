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

func findLHS(nums []int) int {
	m := make(map[int]int)
	for _, v := range nums {
		m[v]++
	}
	ret := 0
	for k, v := range m {
		if m[k+1] != 0 {
			ret = max(ret, v+m[k+1])
		}
	}

	return ret
}
func main() {
	fmt.Println(findLHS([]int{1, 3, 2, 1}))
}
