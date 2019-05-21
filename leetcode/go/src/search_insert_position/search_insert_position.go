// 2017-12-04 11:17
package main

import "fmt"

func searchInsert(nums []int, target int) int {
	lo, hi := 0, len(nums)-1
	for lo <= hi {
		mid := lo + (hi-lo)/2
		if nums[mid] >= target {
			hi = mid - 1
		} else {
			lo = mid + 1
		}
	}
	return lo
}
func main() {
	fmt.Println(searchInsert([]int{1, 3, 5, 6}, 7))
}
