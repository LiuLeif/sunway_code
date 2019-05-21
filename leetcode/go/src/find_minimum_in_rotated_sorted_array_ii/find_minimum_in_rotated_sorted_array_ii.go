// 2017-12-08 10:08
package main

import "fmt"

func findMin(nums []int) int {
	lo, hi := 0, len(nums)-1
	for lo < hi {
		mid := lo + (hi-lo)/2
		if nums[hi] < nums[mid] {
			lo = mid + 1
		} else if nums[hi] > nums[mid] {
			hi = mid
		} else {
			hi--
		}
	}

	return nums[lo]
}
func main() {
	fmt.Println(findMin([]int{3, 4, 5, 1, 2}))
}
