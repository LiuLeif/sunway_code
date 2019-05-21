// 2017-11-14 18:54
package main

import "fmt"

func removeKdigits(num string, k int) string {
	stack := []byte{}
	for i := 0; i < len(num); i++ {
		if len(stack) == 0 || k == 0 {
			stack = append(stack, num[i])
		} else {
			for k != 0 && len(stack) != 0 {
				top := stack[len(stack)-1]
				if num[i] < top {
					stack = stack[:len(stack)-1]
					k--
				} else {
					break
				}
			}
			stack = append(stack, num[i])
		}
	}

	for k > 0 {
		stack = stack[:len(stack)-1]
		k--
	}

	zero := -1
	for i := 0; i < len(stack); i++ {
		if stack[i] == '0' {
			zero = i
		} else {
			break
		}
	}

	if zero != -1 {
		stack = stack[zero+1:]
	}
	if len(stack) == 0 {
		return "0"
	}
	return string(stack)
}
func main() {
	fmt.Println(removeKdigits("11", 1))
}

// 119
