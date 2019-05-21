// 2017-11-14 18:54
package main

import (
	"fmt"
)

func findDuplicates(nums []int) []int {
	ret := []int{}
	for i := 0; i < len(nums); i++ {
		abs := nums[i]
		if abs < 0 {
			abs = -nums[i]
		}
		index := abs - 1
		if nums[index] < 0 {
			ret = append(ret, abs)
		}
		nums[index] = -nums[index]
	}
	return ret
}
func main() {
	fmt.Println(findDuplicates([]int{4, 3, 1, 3}))
}
