// 2018-11-26 13:55
package main

import "fmt"

func validateStackSequences(pushed []int, popped []int) bool {
	// 1 2 3 4
	// 3 2 1 4
	stack := []int{}
	index := 0
	for _, n := range pushed {
		stack = append(stack, n)
		for len(stack) != 0 && stack[len(stack)-1] == popped[index] {
			index++
			stack = stack[:len(stack)-1]
		}
	}
	return index == len(popped)
}

func main() {
	fmt.Println(validateStackSequences([]int{1, 2, 3, 4, 5}, []int{4, 5, 3, 2, 1}))
	fmt.Println(validateStackSequences([]int{1, 0}, []int{1, 1}))
}
