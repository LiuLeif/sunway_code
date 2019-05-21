// 2017-12-08 14:35
package main

import "fmt"

func singleNumber(nums []int) []int {
	ret := [2]int{0, 0}
	diff := 0
	for _, n := range nums {
		diff ^= n
	}
	diff &= -diff
	for _, n := range nums {
		if n&diff == 0 {
			ret[0] ^= n
		} else {
			ret[1] ^= n
		}
	}
	return ret[:]
}

//  11
// 101
// 110
func main() {
	fmt.Println(singleNumber([]int{1, 2, 1, 3, 3, 4, 2, 5}))
}
