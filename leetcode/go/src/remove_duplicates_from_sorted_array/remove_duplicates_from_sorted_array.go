// 2017-11-14 18:54
package main

import "fmt"

func removeDuplicates(nums []int) int {
	if len(nums) == 0 {
		return 0
	}
	start := 0
	for i := 0; i < len(nums); i++ {
		if nums[i] != nums[start] {
			start++
			nums[i], nums[start] = nums[start], nums[i]
		}
	}
	fmt.Println(nums)
	return start + 1
}
func main() {
	fmt.Println(removeDuplicates([]int{1, 1, 1, 2, 2, 2, 3, 3, 3, 4, 4, 5}))
}
