// 2017-11-29 10:16
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

func rob1(nums []int) (int, int) {
	rob, notrob := 0, 0
	for i := 0; i < len(nums); i++ {
		rob, notrob = notrob+nums[i], max(rob, notrob)
	}
	return rob, notrob
}

func rob(nums []int) int {
	if len(nums) == 0 {
		return 0
	}
	if len(nums) == 1 {
		return nums[0]
	}
	_, x := rob1(nums)
	y, _ := rob1(nums[1:])
	return max(x, y)
}
func main() {
	fmt.Println(rob([]int{2, 1}))
}
