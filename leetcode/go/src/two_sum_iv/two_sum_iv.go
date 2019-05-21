// 2017-11-21 10:50
package main

import (
	"fmt"

	. "util/tree"
)

func findOtherValue(root *TreeNode, k int, skipNode *TreeNode) bool {
	if root == nil {
		return false
	}
	if k == root.Val && root != skipNode {
		return true
	}

	if k > root.Val {
		return findOtherValue(root.Right, k, skipNode)
	} else {
		return findOtherValue(root.Left, k, skipNode)
	}
}

func doFind(origRoot, root *TreeNode, k int) bool {
	if root == nil {
		return false
	}

	ret := findOtherValue(origRoot, k-root.Val, root) || findOtherValue(origRoot, k-root.Val, root)
	if ret {
		return true
	}

	return doFind(origRoot, root.Left, k) || doFind(origRoot, root.Right, k)
}

func findTarget(root *TreeNode, k int) bool {
	return doFind(root, root, k)
}
func main() {
	t := NewBinaryTree([]string{"5", "3", "6", "2", "4", "null", "7"})
	fmt.Println(findTarget(t, 7))
}
