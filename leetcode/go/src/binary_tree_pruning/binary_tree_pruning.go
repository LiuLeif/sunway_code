// 2018-04-13 11:42
package main

import (
	. "util/tree"
)

func pruneTree(root *TreeNode) *TreeNode {
	if root == nil {
		return nil
	}
	root.Left = pruneTree(root.Left)
	root.Right = pruneTree(root.Right)
	if root.Left == nil && root.Right == nil && root.Val == 0 {
		return nil
	}
	return root
}
func main() {
	root := NewBinaryTree([]string{"1", "0", "0", "1"})
	root = pruneTree(root)
	root.Dump()
}
