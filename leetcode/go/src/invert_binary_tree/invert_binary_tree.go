// 2017-12-11 16:36
package main

import "fmt"

import (
	. "util/tree"
)

func invertTree(root *TreeNode) *TreeNode {
	if root != nil {
		root.Left, root.Right = root.Right, root.Left
		invertTree(root.Left)
		invertTree(root.Right)
	}
	return root
}
func main() {
	t := NewBinaryTree([]string{"1", "2", "3", "4", "5"})
	t = invertTree(t)
	fmt.Println(t.Left)
}
