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

func rob(nums []int) int {
	rob, notrob := 0, 0
	for i := 0; i < len(nums); i++ {
		rob, notrob = notrob+nums[i], max(rob, notrob)
	}
	return max(rob, notrob)
}

func main() {
	fmt.Println(rob([]int{104, 209, 137, 52, 158, 67, 213, 86, 141, 110, 151, 127, 238, 147, 169, 138, 240, 185, 246, 225, 147, 203, 83, 83, 131, 227, 54, 78, 165, 180, 214, 151, 111, 161, 233, 147, 124, 143}))
	fmt.Println(rob([]int{1, 2, 3, 4, 5}))
}
