// 2017-12-07 19:09
package main

import "fmt"

func isValid(s string) bool {
	stack := []byte{}
	m := map[byte]byte{
		'(': ')',
		'[': ']',
		'{': '}',
	}
	for i := 0; i < len(s); i++ {
		if s[i] == '(' || s[i] == '[' || s[i] == '{' {
			stack = append(stack, s[i])
		} else {
			if len(stack) == 0 {
				return false
			}
			top := stack[len(stack)-1]
			if m[top] != s[i] {
				return false
			}
			stack = stack[:len(stack)-1]
		}
	}

	return len(stack) == 0
}
func main() {
	fmt.Println(isValid("()[]{()"))
}
