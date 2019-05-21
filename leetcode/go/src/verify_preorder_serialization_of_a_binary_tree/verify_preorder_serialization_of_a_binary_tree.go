// 2018-01-31 15:40
package main

import (
	"fmt"
	"strings"
)

func isValidSerialization(preorder string) bool {
	order := strings.Split(preorder, ",")
	stack := []string{}
	for i := 0; i < len(order); i++ {
		stack = append(stack, order[i])
		for len(stack) >= 3 && stack[len(stack)-1] == "#" && stack[len(stack)-2] == "#" && stack[len(stack)-3] != "#" {
			stack = stack[:len(stack)-3]
			stack = append(stack, "#")
		}
	}
	return len(stack) == 1 && stack[len(stack)-1] == "#"
}
func main() {
	fmt.Println(isValidSerialization("9,3,4,#,#,1,#,#,2,#,6,#,#"))
	fmt.Println(isValidSerialization(""))
}
