// 2017-11-14 18:54
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

func findMaxConsecutiveOnes(nums []int) int {
	ret := 0
	curr := 0
	for i := 0; i < len(nums); i++ {
		if nums[i] == 1 {
			curr++
		} else {
			ret = max(ret, curr)
			curr = 0
		}
	}
	ret = max(ret, curr)
	return ret
}
func main() {
	fmt.Println(findMaxConsecutiveOnes([]int{1, 1, 0, 1, 1, 1}))
}
