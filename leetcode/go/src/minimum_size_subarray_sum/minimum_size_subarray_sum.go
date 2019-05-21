// 2017-11-14 18:54
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

func minSubArrayLen(s int, nums []int) int {
	if len(nums) == 0 {
		return 0
	}
	right, sum := 0, 0
	ret := math.MaxInt32

	for left := 0; left < len(nums); left++ {
		for sum < s && right < len(nums) {
			sum += nums[right]
			right++
		}

		if sum >= s {
			ret = min(ret, right-left)
		}
		sum -= nums[left]
	}

	if ret == math.MaxInt32 {
		return 0
	}
	return ret
}
func main() {
	fmt.Println(minSubArrayLen(7, []int{1, 6, 8}))
}
