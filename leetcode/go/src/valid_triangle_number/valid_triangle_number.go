// 2018-01-22 11:24
package main

import (
	"fmt"
	"sort"
)

func triangleNumber(nums []int) int {
	ret := 0
	sort.Ints(nums)
	for i := len(nums) - 1; i >= 2; i-- {
		lo, hi := 0, i-1
		for lo < hi {
			if nums[lo]+nums[hi] > nums[i] {
				ret += hi - lo
				hi--
			} else {
				lo++
			}
		}
	}
	return ret
}
func main() {
	fmt.Println(triangleNumber([]int{2, 2, 3, 4}))
}
