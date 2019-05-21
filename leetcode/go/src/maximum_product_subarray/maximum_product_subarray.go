// 2017-12-06 15:10
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

func min(nums ...int) int {
	ret := math.MaxInt32
	for _, n := range nums {
		if ret > n {
			ret = n
		}
	}
	return ret
}

func maxProduct(nums []int) int {
	imax, imin := nums[0], nums[0]
	ret := imax
	for i := 1; i < len(nums); i++ {
		imax, imin = max(nums[i], nums[i]*imax, nums[i]*imin), min(nums[i], nums[i]*imax, nums[i]*imin)
		ret = max(ret, imax)
	}

	return ret
}
func main() {
	fmt.Println(maxProduct([]int{-1, -2, -9, -6}))
}
