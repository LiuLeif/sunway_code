// 2017-11-14 18:54
package main

import (
	"fmt"
	"sort"
)

func findUnsortedSubarray(nums []int) int {
	nums2 := make([]int, len(nums))
	copy(nums2, nums)
	sort.Ints(nums2)
	// 12345
	// 13542
	left, right := -1, -1
	for i := 0; i < len(nums); i++ {
		if nums[i] != nums2[i] {
			left = i
			break
		}
	}
	if left == -1 {
		return 0
	}
	for i := len(nums) - 1; i >= 0; i-- {
		if nums[i] != nums2[i] {
			right = i
			break
		}
	}
	return right - left + 1
}
func main() {
	fmt.Println(findUnsortedSubarray([]int{1, 3, 5, 4, 2}))
}
