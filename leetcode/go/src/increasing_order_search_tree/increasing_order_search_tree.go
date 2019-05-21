// 2018-10-16 14:12
package main

import (
	. "util/tree"
)

func increasingBST(root *TreeNode) *TreeNode {
	var traverse func(root *TreeNode) (*TreeNode, *TreeNode)
	traverse = func(root *TreeNode) (*TreeNode, *TreeNode) {
		if root == nil {
			return nil, nil
		}
		leftHead, leftTail := traverse(root.Left)
		rightHead, rightTail := traverse(root.Right)
		root.Left = nil
		root.Right = rightHead

		if leftHead == nil {
			leftHead = root
		} else {
			leftTail.Right = root
		}

		if rightTail == nil {
			rightTail = root
		}
		return leftHead, rightTail
	}
	ret, _ := traverse(root)
	return ret
}
func main() {
	root := NewBinaryTree([]string{"5", "3", "6"})
	root.Dump()
	root = increasingBST(root)
	root.Dump()
}
