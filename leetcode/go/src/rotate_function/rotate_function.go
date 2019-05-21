// 2018-01-09 13:26
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

func maxRotateFunction(nums []int) int {
	if len(nums) == 0 {
		return 0
	}
	sum := 0
	for i := 0; i < len(nums); i++ {
		sum += nums[i]
	}

	next := 0
	for i := 0; i < len(nums); i++ {
		next += i * nums[i]
	}

	ret := math.MinInt32
	for i := 0; i < len(nums); i++ {
		next = next - sum + nums[i]*len(nums)
		ret = max(ret, next)
	}

	return ret
}
func main() {
	fmt.Println(maxRotateFunction([]int{4, 3}))
}
