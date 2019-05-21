// 2018-01-03 18:42
package main

import "fmt"

func singleNumber(nums []int) int {
	ones, twos, threes := 0, 0, 0
	for _, v := range nums {
		twos |= ones & v
		ones ^= v
		threes = twos & ones
		twos &= ^threes
		ones &= ^threes
	}
	return ones
}
func main() {
	fmt.Println(singleNumber([]int{1, 1, 1, 2, 2, 2, 3}))
}
