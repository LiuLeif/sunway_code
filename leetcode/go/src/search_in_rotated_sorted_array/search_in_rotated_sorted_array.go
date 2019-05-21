// 2018-01-09 14:55
package main

import "fmt"

func search(nums []int, target int) int {
	f, t := 0, len(nums)-1
	for f <= t {
		mid := f + (t-f)/2
		if nums[mid] == target {
			return mid
		}
		if nums[f] <= nums[mid] {
			if target >= nums[f] && target < nums[mid] {
				t = mid - 1
			} else {
				f = mid + 1
			}
		} else {
			if target > nums[mid] && target <= nums[t] {
				f = mid + 1
			} else {
				t = mid - 1
			}
		}
	}
	return -1
}
func main() {
	fmt.Println(search([]int{1, 3}, 3))
}
