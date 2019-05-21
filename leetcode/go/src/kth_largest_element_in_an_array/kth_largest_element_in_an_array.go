// 2017-11-14 18:54
package main

import "fmt"

func partition(nums []int) int {
	s := 0
	f := 1
	t := len(nums) - 1
	for f <= t {
		if nums[f] > nums[s] {
			nums[f], nums[s+1] = nums[s+1], nums[f]
			nums[s], nums[s+1] = nums[s+1], nums[s]
			s++
		}
		f++
	}
	return s
}

func findKthLargest(nums []int, k int) int {
	s := partition(nums)
	if s == k-1 {
		return nums[s]
	}
	if s > k-1 {
		return findKthLargest(nums[:s], k)
	} else {
		return findKthLargest(nums[s+1:], k-s-1)
	}
}
func main() {
	fmt.Println(findKthLargest([]int{5, 1}, 2))
}
