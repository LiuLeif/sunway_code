// 2017-12-07 14:07
package main

import "fmt"

func findMin(nums []int) int {
	lo, hi := 0, len(nums)-1
	for lo < hi {
		if nums[lo] < nums[hi] {
			return nums[lo]
		}
		mid := lo + (hi-lo)/2
		if nums[lo] <= nums[mid] {
			lo = mid + 1
		} else {
			hi = mid
		}
	}

	return nums[lo]
}

func main() {
	fmt.Println(findMin([]int{3, 4, 5, 1, 2}))
}
