// 2018-10-16 11:00
package main

import (
	. "util/tree"
)

func searchBST(root *TreeNode, val int) *TreeNode {
	if root == nil {
		return nil
	}
	if root.Val == val {
		return root
	}
	left := searchBST(root.Left, val)
	if left != nil {
		return left
	}
	right := searchBST(root.Right, val)
	if right != nil {
		return right
	}
	return nil
}

func main() {
	t := NewBinaryTree([]string{"4", "2", "7", "1", "3"})
	subTree := searchBST(t, 44)
	subTree.Dump()
}
