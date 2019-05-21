// 2017-12-19 15:43
package main

import (
	"math"
	. "util/tree"
)

func traverse(first, second, prev **TreeNode, root *TreeNode) {
	if root == nil {
		return
	}
	traverse(first, second, prev, root.Left)

	if *first == nil && (*prev).Val >= root.Val {
		*first = *prev
	}
	if *first != nil && (*prev).Val >= root.Val {
		*second = root
	}
	*prev = root

	traverse(first, second, prev, root.Right)
}

func recoverTree(root *TreeNode) {
	var first *TreeNode
	var second *TreeNode
	var prev *TreeNode = &TreeNode{math.MinInt32, nil, nil}
	traverse(&first, &second, &prev, root)
	first.Val, second.Val = second.Val, first.Val
}
func main() {
	t := NewBinaryTree([]string{"2", "4", "3", "null", "null", "2", "1"})
	recoverTree(t)
	t.Dump()
}
