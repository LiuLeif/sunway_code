// 2017-12-25 17:35
package main

import "fmt"

func pop(stack *[]int) int {
	ret := (*stack)[len(*stack)-1]
	*stack = (*stack)[:len(*stack)-1]
	return ret
}

func push(stack *[]int, x int) {
	*stack = append(*stack, x)
}

func calculate(s string) int {
	digit := 0
	sign := 1
	result := 0

	prevResult := []int{}
	prevSign := []int{}

	for i := 0; i < len(s); i++ {
		if s[i] >= '0' && s[i] <= '9' {
			digit = digit*10 + int(s[i]) - '0'
		} else {
			result = result + sign*digit
			digit = 0
			if s[i] == '+' {
				sign = 1
			} else if s[i] == '-' {
				sign = -1
			} else if s[i] == '(' {
				push(&prevResult, result)
				push(&prevSign, sign)
				result = 0
				sign = 1
			} else if s[i] == ')' {
				digit = result
				result = pop(&prevResult)
				sign = pop(&prevSign)
			}
		}
	}
	result = result + sign*digit
	return result
}
func main() {
	fmt.Println(calculate("1+21-(3+1)"))
}
