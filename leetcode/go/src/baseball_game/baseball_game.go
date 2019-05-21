// 2018-02-24 15:23
package main

import (
	"fmt"
	"strconv"
)

func sum(nums ...int) int {
	ret := 0
	for _, n := range nums {
		ret += n
	}
	return ret
}

func calPoints(ops []string) int {
	stack := []int{}
	for _, op := range ops {
		switch op {
		case "C":
			stack = stack[:len(stack)-1]
		case "D":
			top := stack[len(stack)-1]
			stack = append(stack, 2*top)
		case "+":
			stack = append(stack, stack[len(stack)-1]+stack[len(stack)-2])
		default:
			opInt, _ := strconv.Atoi(op)
			stack = append(stack, opInt)
		}
	}

	return sum(stack...)
}
func main() {
	fmt.Println(calPoints([]string{"5", "2", "C", "D", "+"}))
	fmt.Println(calPoints([]string{"5", "-2", "4", "C", "D", "9", "+", "+"}))
}
