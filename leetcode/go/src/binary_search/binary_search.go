// 2018-10-15 14:57
package main

import "fmt"

func search(nums []int, target int) int {
	lo, hi := 0, len(nums)-1
	for lo <= hi {
		mid := (lo + hi) / 2
		if nums[mid] == target {
			return mid
		} else if nums[mid] > target {
			hi = mid - 1
		} else {
			lo = mid + 1
		}
	}
	return -1
}
func main() {
	fmt.Println(search([]int{-1, 0, 3, 5, 9, 12}, 10))
}
