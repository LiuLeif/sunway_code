// 2017-12-21 15:30
package main

import (
	. "util/tree"
)

func constructMaximumBinaryTree(nums []int) *TreeNode {
	if len(nums) == 0 {
		return nil
	}
	maxValue := nums[0]
	maxIndex := 0
	for i := 1; i < len(nums); i++ {
		if nums[i] > maxValue {
			maxValue = nums[i]
			maxIndex = i
		}
	}
	left := constructMaximumBinaryTree(nums[:maxIndex])
	right := constructMaximumBinaryTree(nums[maxIndex+1:])
	return &TreeNode{maxValue, left, right}
}
func main() {
	t := constructMaximumBinaryTree([]int{3, 2, 1, 6, 0, 5})
	t.Dump()
}
