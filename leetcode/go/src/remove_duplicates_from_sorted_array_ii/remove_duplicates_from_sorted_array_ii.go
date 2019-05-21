// 2017-12-29 14:47
package main

import "fmt"

func removeDuplicates(nums []int) int {
	k := 2
	if len(nums) <= k {
		return len(nums)
	}
	start := 0
	count := 0
	for i := 0; i < len(nums); i++ {
		if nums[i] != nums[start] {
			count = 0
			start++
			nums[i], nums[start] = nums[start], nums[i]
		} else {
			count++
			if count < k {
				start++
				nums[i], nums[start] = nums[start], nums[i]
			}
		}
	}
	fmt.Println(nums)
	return start + 1
}

func main() {
	fmt.Println(removeDuplicates([]int{1, 1, 1, 1, 2, 2, 2, 2, 3, 3, 5}))
}
