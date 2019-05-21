// 2017-12-26 16:09
package main

import (
	"fmt"
)

func pop(stack *[]int) int {
	ret := (*stack)[len(*stack)-1]
	*stack = (*stack)[:len(*stack)-1]
	return ret
}

func push(stack *[]int, x int) {
	*stack = append(*stack, x)
}

func calculate(s string) int {
	s = s + "+"
	resultA, resultB, sign := 0, 0, int('+')
	digit := 0

	prevResultA := []int{}
	prevResultB := []int{}
	prevSign := []int{}

	for i := 0; i < len(s); i++ {
		if s[i] == ' ' {
			continue
		}
		if s[i] >= '0' && s[i] <= '9' {
			digit = digit*10 + int(s[i]) - '0'
		} else {
			switch sign {
			case '+':
				resultA = resultA + digit
				resultB = digit
			case '-':
				resultA = resultA - digit
				resultB = -digit
			case '*':
				resultA = resultA - resultB + resultB*digit
				resultB = resultB * digit
			case '/':
				resultA = resultA - resultB + resultB/digit
				resultB = resultB / digit
			}
			switch s[i] {
			case '+', '-', '*', '/':
				digit = 0
				sign = int(s[i])
			case '(':
				push(&prevResultA, resultA)
				push(&prevResultB, resultB)
				push(&prevSign, sign)
				resultA, resultB, sign = 0, 0, int('+')
			case ')':
				digit = resultA
				resultA = pop(&prevResultA)
				resultB = pop(&prevResultB)
				sign = pop(&prevSign)
			}
		}
	}

	return resultA
}
func main() {
	fmt.Println(calculate("(1+2*2)*3"))
}
