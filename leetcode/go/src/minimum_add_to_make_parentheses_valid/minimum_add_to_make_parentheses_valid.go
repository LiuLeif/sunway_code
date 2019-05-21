// 2018-10-17 10:55
package main

import "fmt"

func minAddToMakeValid(S string) int {
	stack := []rune{}
	for _, c := range S {
		if c == '(' || len(stack) == 0 {
			stack = append(stack, c)
		} else {
			if stack[len(stack)-1] == '(' {
				stack = stack[:len(stack)-1]
			} else {
				stack = append(stack, c)
			}
		}
	}
	return len(stack)
}
func main() {
	fmt.Println(minAddToMakeValid("()))(("))
}
