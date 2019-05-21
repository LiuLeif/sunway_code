// 2018-01-23 13:44
package main

import (
	"fmt"
	"math"
)

func min(nums ...int) int {
	ret := math.MaxInt32
	for _, n := range nums {
		if ret > n {
			ret = n
		}
	}
	return ret
}
func sum(nums []int) int {
	ret := 0
	for _, v := range nums {
		ret += v
	}
	return ret
}
func minMoves(nums []int) int {
	return sum(nums) - len(nums)*min(nums...)
}
func main() {
	fmt.Println(minMoves([]int{1, 2, 3}))
}
