// 2018-12-03 14:26
package main

import (
	"fmt"
	. "util/tree"
)

/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
func flipEquiv(root1 *TreeNode, root2 *TreeNode) bool {
	if root1 == nil && root2 == nil {
		return true
	}
	if root1 == nil || root2 == nil {
		return false
	}
	if root1.Val != root2.Val {
		return false
	}

	a := flipEquiv(root1.Left, root2.Left)
	b := flipEquiv(root1.Right, root2.Right)

	c := flipEquiv(root1.Left, root2.Right)
	d := flipEquiv(root1.Right, root2.Left)

	if (a && b) || (c && d) {
		return true
	}
	return false
}

func main() {
	root1 := NewBinaryTree([]string{"1", "2", "3", "4", "5", "6", "null", "null", "null", "7", "8"})
	root2 := NewBinaryTree([]string{"1", "3", "2", "null", "6", "4", "5", "null", "null", "null", "null", "8", "7"})
	root1.Dump()
	root2.Dump()
	fmt.Println(flipEquiv(root1, root2))
}
