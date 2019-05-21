// 2018-01-23 17:29
package main

import (
	"fmt"
	"math"
)

func getRightSmaller(nums []int) []int {
	ret := make([]int, len(nums))
	for i := 0; i < len(nums); i++ {
		ret[i] = len(nums)
	}
	stack := []int{}
	for i := 0; i < len(nums); i++ {
		for len(stack) != 0 && nums[stack[len(stack)-1]] > nums[i] {
			ret[stack[len(stack)-1]] = i
			stack = stack[:len(stack)-1]
		}
		stack = append(stack, i)
	}
	return ret
}

func getLeftSmaller(nums []int) []int {
	ret := make([]int, len(nums))
	for i := 0; i < len(nums); i++ {
		ret[i] = -1
	}

	stack := []int{}
	for i := len(nums) - 1; i >= 0; i-- {
		for len(stack) != 0 && nums[stack[len(stack)-1]] > nums[i] {
			ret[stack[len(stack)-1]] = i
			stack = stack[:len(stack)-1]
		}
		stack = append(stack, i)
	}
	return ret
}

func max(nums ...int) int {
	ret := math.MinInt32
	for _, n := range nums {
		if ret < n {
			ret = n
		}
	}
	return ret
}

func largestRectangleArea(heights []int) int {
	right := getRightSmaller(heights)
	left := getLeftSmaller(heights)
	ret := 0
	for i := 0; i < len(heights); i++ {
		ret = max(ret, (right[i]-left[i]-1)*heights[i])
	}
	return ret
}
func main() {
	fmt.Println(largestRectangleArea([]int{2, 1, 5, 6, 2, 3}))
}
