// 2017-12-20 17:40
package main

import "fmt"

func ok(nums []int) bool {
	fmt.Println("OK", nums)
	for i := 1; i < len(nums); i++ {
		if nums[i] < nums[i-1] {
			return false
		}
	}
	return true
}

func checkPossibility(nums []int) bool {
	for i := 1; i < len(nums); i++ {
		if nums[i] < nums[i-1] {
			x := nums[i-1]
			nums[i-1] = nums[i]
			if ok(nums) {
				return true
			}
			nums[i-1] = x
			nums[i] = x
			if ok(nums) {
				return true
			}
			return false
		}
	}
	return true
}
func main() {
	fmt.Println(checkPossibility([]int{1, 2, 3, 3, 1}))
}
