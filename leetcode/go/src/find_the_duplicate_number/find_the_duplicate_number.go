// 2017-11-14 18:54
package main

import "fmt"

func findDuplicate(nums []int) int {
	n := len(nums) - 1
	left := 1
	right := n
	for left < right {
		mid := left + (right-left)/2
		count := 0
		for i := 0; i < len(nums); i++ {
			if nums[i] <= mid {
				count++
			}
		}
		if count > mid {
			right = mid
		} else {
			left = mid + 1
		}
	}
	return left
}
func main() {
	fmt.Println(findDuplicate([]int{1, 2, 3, 4, 5, 5}))
}
