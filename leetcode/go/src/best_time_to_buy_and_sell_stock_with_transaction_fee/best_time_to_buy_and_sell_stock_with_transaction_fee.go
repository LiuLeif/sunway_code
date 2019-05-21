// 2017-11-21 16:24
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

func maxProfit(nums []int, fee int) int {
	full := make([]int, len(nums))
	empty := make([]int, len(nums))
	full[0] = -nums[0]
	empty[0] = 0
	for i := 1; i < len(nums); i++ {
		full[i] = max(full[i-1], empty[i-1]-nums[i])
		empty[i] = max(empty[i-1], full[i-1]+nums[i]-fee)
	}

	return empty[len(nums)-1]
}
func main() {
	fmt.Println(maxProfit([]int{1, 3, 2, 8, 4, 9}, 2))
}
