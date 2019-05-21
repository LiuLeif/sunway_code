// 2017-12-12 16:01
package main

import (
	"fmt"
	"math"
)

func max(nums ...int) int {
	ret := math.MinInt32
	for _, n := range nums {
		if ret < n {
			ret = n
		}
	}
	return ret
}

func longestValidParentheses(s string) int {
	stack := []byte{}
	count := []int{}
	for i := 0; i < len(s); i++ {
		if s[i] == '(' || len(stack) == 0 {
			stack = append(stack, s[i])
			count = append(count, i)
		} else {
			top := stack[len(stack)-1]
			if top == '(' {
				stack = stack[:len(stack)-1]
				count = count[:len(count)-1]
			} else {
				stack = append(stack, ')')
				count = append(count, i)
			}
		}
	}

	ret := 0
	count = append([]int{-1}, count...)
	count = append(count, len(s))
	for i := 1; i < len(count); i++ {
		ret = max(ret, count[i]-count[i-1]-1)
	}

	return ret
}

func main() {
	fmt.Println(longestValidParentheses(")))()((()))"))
}
