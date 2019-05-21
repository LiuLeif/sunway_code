// 2017-12-22 15:50
package main

import "fmt"

func findDisappearedNumbers(nums []int) []int {
	ret := []int{}
	for i := 0; i < len(nums); i++ {
		k := nums[i]
		if k < 0 {
			k = -k
		}
		if nums[k-1] > 0 {
			nums[k-1] = -nums[k-1]
		}
	}
	for i := 0; i < len(nums); i++ {
		if nums[i] > 0 {
			ret = append(ret, i+1)
		}
	}
	return ret
}
func main() {
	fmt.Println(findDisappearedNumbers([]int{2, 2}))
}
