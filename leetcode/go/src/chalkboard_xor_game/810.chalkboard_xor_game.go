// 2018-11-08 10:33
package main

import (
	"fmt"
)

func xorGame(nums []int) bool {
	// 1 1 1 2
	sum := 0
	for _, v := range nums {
		sum ^= v
	}
	if sum == 0 {
		return true
	}
	return len(nums)%2 == 0
}

func main() {
	fmt.Println(xorGame([]int{1, 2}))
}
