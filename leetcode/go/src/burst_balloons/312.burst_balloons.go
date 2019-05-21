// 2018-11-27 13:08
package main

import (
	"fmt"
	"math"
)

func maxCoins(nums []int) int {
	//1 3 1 5 8 1
	//l   r
	nums = append(nums, 1)
	nums = append([]int{1}, nums...)
	memo := make([][]int, len(nums))
	for i := 0; i < len(nums); i++ {
		memo[i] = make([]int, len(nums))
	}

	max := func(nums ...int) int {
		ret := math.MinInt32
		for _, n := range nums {
			if ret < n {
				ret = n
			}
		}
		return ret
	}

	var burst func(left, right int) int

	burst = func(left, right int) int {
		if memo[left][right] != 0 {
			return memo[left][right]
		}
		for i := left + 1; i < right; i++ {
			memo[left][right] = max(memo[left][right], nums[left]*nums[right]*nums[i]+burst(left, i)+burst(i, right))
		}
		return memo[left][right]
	}

	return burst(0, len(nums)-1)
}

func main() {
	fmt.Println(maxCoins([]int{3, 1, 5, 8}))
}
