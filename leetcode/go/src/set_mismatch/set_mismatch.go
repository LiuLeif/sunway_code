// 2017-11-21 13:37
package main

import "fmt"

func findErrorNums(nums []int) []int {
	ret := []int{0, 0}
	for i := 0; i < len(nums); i++ {
		x := nums[i]
		if x < 0 {
			x = -x
		}
		if nums[x-1] < 0 {
			ret[0] = x
		}
		if nums[x-1] > 0 {
			nums[x-1] = -nums[x-1]
		}
	}
	for i := 0; i < len(nums); i++ {
		if nums[i] > 0 {
			ret[1] = i + 1
		}
	}
	return ret
}
func main() {
	fmt.Println(findErrorNums([]int{4, 4, 2, 1}))
}
