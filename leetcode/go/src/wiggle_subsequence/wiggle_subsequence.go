// 2018-02-26 16:15
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

func wiggleMaxLength(nums []int) int {
	if len(nums) == 0 {
		return 0
	}
	up := 1
	down := 1
	for i := 1; i < len(nums); i++ {
		if nums[i] > nums[i-1] {
			up = down + 1
		} else if nums[i] < nums[i-1] {
			down = up + 1
		}
	}
	return max(up, down)
}
func main() {
	fmt.Println(wiggleMaxLength([]int{1, 17, 5, 10, 13, 15, 10, 5, 16, 8}))
	fmt.Println(wiggleMaxLength([]int{1, 2, 3, 4, 5, 6, 7, 8, 9}))
	fmt.Println(wiggleMaxLength([]int{0, 0}))
}
