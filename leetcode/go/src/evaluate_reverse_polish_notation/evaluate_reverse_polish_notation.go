// 2018-01-12 17:08
package main

import (
	"fmt"
	"strconv"
)

func pop(stack *[]int) int {
	ret := (*stack)[len(*stack)-1]
	*stack = (*stack)[:len(*stack)-1]
	return ret
}

func push(stack *[]int, x int) {
	*stack = append(*stack, x)
}

func evalRPN(tokens []string) int {
	stack := []int{}
	m := map[string]func(a, b int) int{}
	m["+"] = func(a, b int) int {
		return a + b
	}
	m["-"] = func(a, b int) int {
		return a - b
	}
	m["*"] = func(a, b int) int {
		return a * b
	}
	m["/"] = func(a, b int) int {
		return a / b
	}
	for i := 0; i < len(tokens); i++ {
		switch tokens[i] {
		case "+", "-", "*", "/":
			b, a := pop(&stack), pop(&stack)
			push(&stack, m[tokens[i]](a, b))
		default:
			v, _ := strconv.Atoi(tokens[i])
			push(&stack, v)
		}
	}
	return stack[0]
}
func main() {
	fmt.Println(evalRPN([]string{"2", "1", "+", "3", "*"}))
	fmt.Println(evalRPN([]string{"4", "13", "+"}))
}
