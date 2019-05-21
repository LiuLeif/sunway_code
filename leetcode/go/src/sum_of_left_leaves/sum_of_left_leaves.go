// 2017-12-29 14:47
package main

import "fmt"
import (
	. "util/tree"
)

func helper(root *TreeNode, isLeft bool) int {
	if root == nil {
		return 0
	}
	if root.Left == nil && root.Right == nil {
		if isLeft {
			return root.Val
		} else {
			return 0
		}
	}
	return helper(root.Left, true) + helper(root.Right, false)
}

func sumOfLeftLeaves(root *TreeNode) int {
	return helper(root, false)
}
func main() {
	root := NewBinaryTree([]string{"3", "9", "20", "null", "null", "15", "7"})
	fmt.Println(sumOfLeftLeaves(root))
}
