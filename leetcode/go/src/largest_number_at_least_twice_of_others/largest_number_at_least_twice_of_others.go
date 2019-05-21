// 2018-02-26 13:15
package main

import "fmt"

func dominantIndex(nums []int) int {
	if len(nums) == 0 {
		return -1
	}
	if len(nums) == 1 {
		return 0
	}
	first, second := -1, -1
	for i := 0; i < len(nums); i++ {
		if first == -1 {
			first = i
			continue
		}
		if nums[i] > nums[first] {
			first, second = i, first
		} else {
			if second == -1 {
				second = i
			} else {
				if nums[i] > nums[second] {
					second = i
				}
			}
		}
	}
	fmt.Println(first, second)
	if nums[first] >= 2*nums[second] {
		return first
	}
	return -1
}
func main() {
	fmt.Println(dominantIndex([]int{1, 0, 3}))
}
