// 2017-11-21 13:37
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

func maxProfit(nums []int) int {
	if len(nums) == 0 {
		return 0
	}
	m := make([]int, len(nums))
	m[len(nums)-1] = nums[len(nums)-1]
	for i := len(nums) - 2; i >= 0; i-- {
		m[i] = max(m[i+1], nums[i])
	}

	ret := 0
	for i := 0; i < len(nums); i++ {
		ret = max(ret, m[i]-nums[i])
	}

	return ret
}

func maxProfitGeneric(nums []int) int {
	empty := make([]int, len(nums))
	full := make([]int, len(nums))
	empty[0] = 0
	full[0] = -nums[0]

	for i := 1; i < len(nums); i++ {
		empty[i] = max(empty[i-1], full[i-1]+nums[i])
		full[i] = max(full[i-1], -nums[i])
	}
	return empty[len(nums)-1]
}

func main() {
	fmt.Println(maxProfit([]int{1, 2, 3, 4, 5}))
	fmt.Println(maxProfitGeneric([]int{1, 2, 3, 4, 5}))
}
