// 2018-01-08 11:04
package main

import (
	. "util/tree"
)

func sortedArrayToBST(nums []int) *TreeNode {
	if len(nums) == 0 {
		return nil
	}
	mid := (len(nums) - 1) / 2
	root := TreeNode{nums[mid], sortedArrayToBST(nums[:mid]), sortedArrayToBST(nums[mid+1:])}
	return &root
}
func main() {
	t := sortedArrayToBST([]int{-10, -3, 0, 5, 9})
	t.Dump()
}
