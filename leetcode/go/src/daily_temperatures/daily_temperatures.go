// 2018-01-03 18:42
package main

import "fmt"

func dailyTemperatures(nums []int) []int {
	ret := make([]int, len(nums))
	stack := []int{0}
	for i := 1; i < len(nums); i++ {
		for len(stack) != 0 && nums[i] > nums[stack[len(stack)-1]] {
			x := len(stack) - 1
			ret[stack[x]] = i - stack[x]
			stack = stack[:x]
		}
		stack = append(stack, i)
	}
	return ret
}
func main() {
	fmt.Println(dailyTemperatures([]int{73, 74, 71, 76}))
}
