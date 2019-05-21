// 2017-11-24 11:05
package main

import "fmt"

func searchRange(nums []int, target int) []int {
	ret := make([]int, 2)

	left, right := 0, len(nums)-1
	found := false
	for left <= right {
		mid := left + (right-left)/2
		if nums[mid] == target {
			found = true
		}
		if nums[mid] <= target {
			left = mid + 1
		} else {
			right = mid - 1
		}
	}

	if !found {
		return []int{-1, -1}
	}
	ret[1] = right

	left, right = 0, len(nums)-1
	for left <= right {
		mid := left + (right-left)/2
		if nums[mid] < target {
			left = mid + 1
		} else {
			right = mid - 1
		}
	}
	ret[0] = left

	return ret
}
func main() {
	fmt.Println(searchRange([]int{5, 7, 7, 8, 8, 10}, 7))
}
