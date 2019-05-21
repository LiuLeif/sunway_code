// 2017-11-29 15:31
package main

import (
	"fmt"
	. "util/tree"
)

func hasPathSum(root *TreeNode, sum int) bool {
	if root == nil {
		return false
	}
	if sum == root.Val && root.Left == nil && root.Right == nil {
		return true
	}
	if root.Left != nil {
		t := hasPathSum(root.Left, sum-root.Val)
		if t {
			return true
		}
	}
	if root.Right != nil {
		t := hasPathSum(root.Right, sum-root.Val)
		if t {
			return true
		}
	}
	return false
}
func main() {
	t := NewBinaryTree([]string{"1"})
	fmt.Println(hasPathSum(t, 22))
}
