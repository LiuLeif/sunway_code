// 2018-02-09 10:51
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

func max(nums ...int) int {
	ret := math.MinInt32
	for _, n := range nums {
		if ret < n {
			ret = n
		}
	}
	return ret
}

func helper(sums []int, f, t, m int, mem map[[3]int]int) int {
	if m == 1 {
		return sums[t-1] - sums[f]
	}
	if x, ok := mem[[3]int{f, t, m}]; ok {
		return x
	}
	ret := math.MaxInt32
	for i := f; i < t-1; i++ {
		ret = min(max(sums[i]-sums[f], helper(sums, i, t, m-1, mem)), ret)
	}
	mem[[3]int{f, t, m}] = ret
	return ret
}

func splitArray(nums []int, m int) int {
	sums := make([]int, len(nums)+1)
	for i := 0; i < len(nums); i++ {
		sums[i+1] = sums[i] + nums[i]
	}
	mem := map[[3]int]int{}
	return helper(sums, 0, len(sums), m, mem)
}
func main() {
	fmt.Println(splitArray([]int{7, 2, 5, 10, 8}, 3))
}
