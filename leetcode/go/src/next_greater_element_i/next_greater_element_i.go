// 2017-12-25 17:35
package main

import "fmt"

func nextGreaterElement(findNums []int, nums []int) []int {
	ret := []int{}
	stack := []int{}
	m := map[int]int{}
	for i := 0; i < len(nums); i++ {
		for len(stack) != 0 && nums[i] > nums[stack[len(stack)-1]] {
			m[nums[stack[len(stack)-1]]] = nums[i]
			stack = stack[:len(stack)-1]
		}
		stack = append(stack, i)
	}
	for i := 0; i < len(findNums); i++ {
		x, ok := m[findNums[i]]
		if !ok {
			x = -1
		}
		ret = append(ret, x)
	}
	return ret
}
func main() {
	fmt.Println(nextGreaterElement([]int{4, 1, 2}, []int{1, 4, 2}))
}
