// 2017-12-29 14:47
package main

import (
	. "util/tree"
)

func helper(sum int, root *TreeNode) int {
	if root == nil {
		return sum
	}
	sumOfRight := helper(sum, root.Right)
	root.Val += sumOfRight
	return helper(root.Val, root.Left)
}

func convertBST(root *TreeNode) *TreeNode {
	helper(0, root)
	return root
}
func main() {
	t := NewBinaryTree([]string{"5", "2", "13"})
	t = convertBST(t)
	t.Dump()
}
