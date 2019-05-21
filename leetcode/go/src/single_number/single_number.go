// 2018-01-03 18:42
package main

import "fmt"

func singleNumber(nums []int) int {
	ret := 0
	for _, v := range nums {
		ret ^= v
	}
	return ret
}
func main() {
	fmt.Println(singleNumber([]int{1, 1, 2, 2, 3}))
}
