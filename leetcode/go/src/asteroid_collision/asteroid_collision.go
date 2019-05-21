// 2018-01-26 18:25
package main

import (
	"fmt"
)

func asteroidCollision(asteroids []int) []int {
	stack := []int{}
	for i := 0; i < len(asteroids); i++ {
		stack = append(stack, asteroids[i])
		for len(stack) > 1 {
			a := stack[len(stack)-1]
			b := stack[len(stack)-2]
			if a > 0 || b < 0 {
				break
			}
			stack = stack[:len(stack)-2]
			if -a > b {
				stack = append(stack, a)
			} else if -a < b {
				stack = append(stack, b)
			}
		}
	}
	return stack
}
func main() {
	fmt.Println(asteroidCollision([]int{5, 5}))
}
