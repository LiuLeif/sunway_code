// 2018-01-02 13:51
package main

import (
	. "util/tree"
)

func trimBST(root *TreeNode, L int, R int) *TreeNode {
	switch {
	case root == nil:
		return nil
	case root.Val > R:
		return trimBST(root.Left, L, R)
	case root.Val < L:
		return trimBST(root.Right, L, R)
	default:
		root.Left = trimBST(root.Left, L, R)
		root.Right = trimBST(root.Right, L, R)
		return root
	}
}
func main() {
	t := NewBinaryTree([]string{"3", "0", "4", "null", "2", "null", "null", "1"})
	t2 := trimBST(t, 1, 3)
	t2.Dump()
}
