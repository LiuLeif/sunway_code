// 2018-10-17 14:17
package main

import "fmt"

func scoreOfParentheses(S string) int {
	stackP := []rune{}
	stackN := []int{}
	result := 0
	for _, c := range S {
		if c == '(' {
			stackP = append(stackP, c)
			stackN = append(stackN, result)
			result = 0
		} else {
			stackP = stackP[:len(stackP)-1]
			if result == 0 {
				result = 1
			} else {
				result *= 2
			}
			result += stackN[len(stackN)-1]
			stackN = stackN[:len(stackN)-1]
		}
	}
	return result
}
func main() {
	fmt.Println(scoreOfParentheses("(()()())"))
}
