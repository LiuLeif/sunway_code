// 2017-11-17 18:01
package main

import (
	"fmt"
	"sort"
)

func fourSum(nums []int, target int) [][]int {
	ret := [][]int{}
	sort.Ints(nums)

	for i := 0; i < len(nums); i++ {
		if i != 0 && nums[i] == nums[i-1] {
			continue
		}
		for j := i + 1; j < len(nums); j++ {
			if j != i+1 && nums[j] == nums[j-1] {
				continue
			}
			left, right := j+1, len(nums)-1
			for left < right {
				if left != j+1 && nums[left] == nums[left-1] {
					left++
					continue
				}
				if nums[i]+nums[j]+nums[left]+nums[right] == target {
					ret = append(ret, []int{nums[i], nums[j], nums[left], nums[right]})
					left++
					continue
				}
				if nums[i]+nums[j]+nums[left]+nums[right] > target {
					right--
				} else {
					left++
				}
			}

		}
	}

	return ret
}
func main() {
	fmt.Println(fourSum([]int{-3, -2, -1, 0, 0, 1, 2, 3}, 0))
}
