// 2018-01-22 16:03
package main

import (
	. "util/tree"
)

func helper(root *TreeNode, v, d int, level int) {
	if level == d-1 {
		t := root.Left
		root.Left = &TreeNode{v, nil, nil}
		root.Left.Left = t

		t = root.Right
		root.Right = &TreeNode{v, nil, nil}
		root.Right.Right = t

		return
	}
	if root.Left != nil {
		helper(root.Left, v, d, level+1)
	}
	if root.Right != nil {
		helper(root.Right, v, d, level+1)
	}
}

func addOneRow(root *TreeNode, v int, d int) *TreeNode {
	dummyRoot := &TreeNode{0, root, nil}
	helper(dummyRoot, v, d, 0)
	return dummyRoot.Left
}
func main() {
	t := NewBinaryTree([]string{"4", "2", "6", "3", "1", "5", "null"})
	t.Dump()
	t = addOneRow(t, 1, 4)
	t.Dump()
}
