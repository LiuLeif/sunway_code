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

func jump(nums []int) int {
	// 1 1 3 4
	start, end := 0, 0
	ret := 0
	for end < len(nums)-1 {
		nextEnd := end
		for i := start; i <= end; i++ {
			nextEnd = max(nextEnd, nums[i]+i)
		}
		ret++
		start = end + 1
		end = nextEnd
	}
	return ret
}

func main() {
	fmt.Println(jump([]int{3, 1, 1, 2, 2, 1}))
}
