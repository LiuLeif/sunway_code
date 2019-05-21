// 2017-11-28 11:45
package main

import "fmt"

func singleNonDuplicateSimple(nums []int) int {
	lo, hi := 0, len(nums)-1
	for lo < hi {
		mid := lo + (hi-lo)/2
		if nums[mid] == nums[mid^1] {
			lo = mid + 1
		} else {
			hi = mid
		}
	}
	return nums[lo]
}

func singleNonDuplicate(nums []int) int {
	lo, hi := 0, len(nums)-1
	for lo < hi {
		mid := lo + (hi-lo)/2
		x := false
		if mid > 0 {
			x = (nums[mid-1] == nums[mid])
		}
		y := false
		if mid < len(nums)-1 {
			y = (nums[mid+1] == nums[mid])
		}
		if !x && !y {
			return nums[mid]
		}
		if x {
			if (mid-1)%2 == 0 {
				lo = mid + 1
			} else {
				hi = mid - 1
			}
		}
		if y {
			if (mid+1)%2 == 0 {
				hi = mid - 1
			} else {
				lo = mid + 1
			}
		}
	}
	return nums[lo]
}
func main() {
	fmt.Println(singleNonDuplicateSimple([]int{1, 1, 2, 3, 3}))
	// fmt.Println(singleNonDuplicateSimple([]int{0, 1, 1, 2, 2, 3, 3, 4, 4, 8, 8}))
	// fmt.Println(singleNonDuplicateSimple([]int{1}))
}
