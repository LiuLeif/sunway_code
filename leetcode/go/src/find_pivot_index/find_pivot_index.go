// 2017-12-15 17:52
package main

import "fmt"

func pivotIndex(nums []int) int {
	leftSum := make([]int, len(nums))
	for i := len(nums) - 2; i >= 0; i-- {
		leftSum[i] = leftSum[i+1] + nums[i+1]
	}
	rightSum := make([]int, len(nums))
	for i := 1; i < len(nums); i++ {
		rightSum[i] = rightSum[i-1] + nums[i-1]
	}
	for i := 0; i < len(nums); i++ {
		if leftSum[i] == rightSum[i] {
			return i
		}
	}
	return -1
}
func main() {
	fmt.Println(pivotIndex([]int{1, 2, 3}))
}
