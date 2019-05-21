// 2017-11-20 11:39
package main

import (
	"fmt"

	. "util/tree"
)

func doFlattern(root *TreeNode) *TreeNode {
	if root.Left == nil && root.Right == nil {
		return root
	}

	var left *TreeNode = root
	if root.Left != nil {
		left = doFlattern(root.Left)
	}
	var right *TreeNode = nil
	if root.Right != nil {
		right = doFlattern(root.Right)
	}

	left.Right = root.Right
	if root.Left != nil {
		root.Right = root.Left
	}
	root.Left = nil
	if right != nil {
		return right
	}
	return left
}

func flatten(root *TreeNode) {
	if root == nil {
		return
	}
	doFlattern(root)
}
func main() {
	t := NewBinaryTree([]string{"1"})

	flatten(t)
	fmt.Println(t)
	fmt.Println(t.Right)
	// fmt.Println(t.Right.Right)
	// fmt.Println(t.Right.Right.Right)
	// fmt.Println(t.Right.Right.Right.Right)
	// fmt.Println(t.Right.Right.Right.Right.Right)
	// fmt.Println(t.Right.Right.Right.Right.Right.Right)
}
