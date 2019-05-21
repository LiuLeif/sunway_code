// 2017-11-14 18:54
package main

import (
	"fmt"
	"sort"
)

func findPairs(nums []int, k int) int {
	// 1 1 3 4 5
	sort.Ints(nums)

	ret, left, right := 0, 0, 1

	for right < len(nums) {
		if nums[right]-nums[left] == k {
			ret += 1
			left += 1
			for left < len(nums) && nums[left] == nums[left-1] {
				left++
			}
			right = left + 1
		} else if nums[right]-nums[left] > k {
			left += 1
			for left < len(nums) && nums[left] == nums[left-1] {
				left++
			}
			right = left + 1
		} else {
			right += 1
			for right < len(nums) && nums[right] == nums[right-1] {
				right++
			}
		}
	}

	return ret
}
func main() {
	fmt.Println(findPairs([]int{3, 1, 4, 1, 5}, 2))
}
