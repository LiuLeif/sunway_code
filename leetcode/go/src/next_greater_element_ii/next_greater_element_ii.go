// 2017-11-17 14:21
package main

import "fmt"

func nextGreaterElements(nums []int) []int {
	nums = append(nums, nums...)
	stack := []int{}
	ret := make([]int, len(nums))
	for i := 0; i < len(nums); i++ {
		ret[i] = -1
	}

	// stack 用来保存当前 `未决` 的 index, 且 stack 中 nums[index] 是
	// 递减的
	for i := 0; i < len(nums)*2; i++ {
		index := i
		if index >= len(nums) {
			index -= len(nums)
		}

		for len(stack) != 0 && nums[stack[len(stack)-1]] < nums[index] {
			t := stack[len(stack)-1]
			stack = stack[:len(stack)-1]
			ret[t] = nums[index]
		}
		stack = append(stack, index)
	}

	return ret[:len(nums)/2]
}

func main() {
	fmt.Println(nextGreaterElements([]int{1, 2}))
}
