// 2018-10-17 11:28
package main

import (
	. "util/tree"
)

func insertIntoBST(root *TreeNode, val int) *TreeNode {
	if root == nil {
		return &TreeNode{val, nil, nil}
	}
	if val > root.Val {
		root.Right = insertIntoBST(root.Right, val)
	} else if val < root.Val {
		root.Left = insertIntoBST(root.Left, val)
	}
	return root
}
func main() {
	root := NewBinaryTree([]string{"4", "2", "7", "1", "3"})
	root.Dump()
	root = insertIntoBST(root, 5)
	root.Dump()
}
