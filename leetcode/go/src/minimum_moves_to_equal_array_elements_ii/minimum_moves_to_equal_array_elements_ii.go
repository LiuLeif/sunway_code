// 2017-12-22 15:50
package main

import (
	"fmt"
)

func partition(nums []int) int {
	p, f, t := 0, 1, len(nums)-1
	for f <= t {
		if nums[f] <= nums[p] {
			f++
		} else {
			nums[f], nums[p+1] = nums[p+1], nums[f]
			nums[p], nums[p+1] = nums[p+1], nums[p]
			p++
			f++
		}
	}
	return p
}

func findKLargest(nums []int, k int) int {
	p := partition(nums)
	if p == k {
		return nums[p]
	}
	if p > k {
		return findKLargest(nums[:p], k)
	} else {
		return findKLargest(nums[p+1:], k-p-1)
	}
}

func minMoves2(nums []int) int {
	median := findKLargest(nums, len(nums)/2)
	ret := 0
	for i := 0; i < len(nums); i++ {
		x := nums[i] - median
		if x < 0 {
			ret += -x
		} else {
			ret += x
		}
	}
	return ret
}

func main() {
	fmt.Println(minMoves2([]int{1, 1, 2, 2}))
}
