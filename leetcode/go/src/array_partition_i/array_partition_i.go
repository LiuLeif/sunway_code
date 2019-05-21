// 2018-01-23 16:30
package main

import (
	"fmt"
	"sort"
)

func arrayPairSum(nums []int) int {
	sort.Ints(nums)
	ret := 0
	for i := 0; i < len(nums); i += 2 {
		ret += nums[i]
	}
	return ret
}
func main() {
	// 1 2 3 4 5 6
	fmt.Println(arrayPairSum([]int{1, 4, 3, 2}))
}
