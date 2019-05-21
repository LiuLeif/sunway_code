// 2018-01-04 17:04
package main

import "fmt"

func checkValidString(s string) bool {
	// 0 -> (, 1 -> )
	stack := []int{}
	asterisk := 0
	for i := 0; i < len(s); i++ {
		if s[i] == '(' {
			stack = append(stack, 0)
		} else if s[i] == ')' {
			if len(stack) != 0 && stack[len(stack)-1] == 0 {
				stack = stack[:len(stack)-1]
			} else {
				if asterisk > 0 {
					asterisk--
				} else {
					return false
				}
			}
		} else {
			asterisk++
		}
	}

	stack = []int{}
	asterisk = 0
	for i := len(s) - 1; i >= 0; i-- {
		if s[i] == ')' {
			stack = append(stack, 0)
		} else if s[i] == '(' {
			if len(stack) != 0 && stack[len(stack)-1] == 0 {
				stack = stack[:len(stack)-1]
			} else {
				if asterisk > 0 {
					asterisk--
				} else {
					return false
				}
			}
		} else {
			asterisk++
		}
	}

	return true
}
func main() {
	fmt.Println(checkValidString("(())"))
	fmt.Println(checkValidString("(*))"))
	fmt.Println(checkValidString("(()**"))
	fmt.Println(checkValidString("((*)(*))((*"))
}
