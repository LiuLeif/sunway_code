// 2018-01-08 10:16
package main

import (
	"fmt"
	"math"
	"sort"
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

func threeSumClosest(nums []int, target int) int {
	ret := math.MaxInt32
	minDelta := math.MaxInt32
	sort.Ints(nums)
	for i := 0; i < len(nums); i++ {
		left, right := 0, len(nums)-1
		for left < right {
			if left == i {
				left++
				continue
			}
			if right == i {
				right--
				continue
			}
			s := nums[left] + nums[right] + nums[i]
			if s == target {
				return target
			}
			delta := s - target
			if delta < 0 {
				delta = -delta
			}
			if delta < minDelta {
				minDelta = delta
				ret = s
			}

			if s > target {
				right--
			} else {
				left++
			}
		}
	}

	return ret
}
func main() {
	fmt.Println(threeSumClosest([]int{-1, 1, 0, 1}, 6))
}
