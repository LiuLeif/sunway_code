// 2017-12-22 15:50
package main

import "fmt"
import (
	. "util/tree"
)

func reverse(root *TreeNode) {
	if root == nil {
		return
	}
	root.Left, root.Right = root.Right, root.Left
	reverse(root.Left)
	reverse(root.Right)
}
func same(t1, t2 *TreeNode) bool {
	if t1 != nil && t2 != nil {
		return t1.Val == t2.Val && same(t1.Left, t2.Left) && same(t1.Right, t2.Right)
	} else {
		if t1 == nil && t2 == nil {
			return true
		} else {
			return false
		}
	}
}
func isSymmetric(root *TreeNode) bool {
	if root == nil {
		return true
	}
	reverse(root.Right)
	return same(root.Left, root.Right)
}

func helper(left, right *TreeNode) bool {
	if left != nil && right != nil {
		if left.Val != right.Val {
			return false
		}
	} else {
		if left == nil && right == nil {
			return true
		} else {
			return false
		}
	}
	return helper(left.Right, right.Left) && helper(left.Left, right.Right)
}

func isSymmetric2(root *TreeNode) bool {
	if root == nil {
		return true
	}
	return helper(root.Left, root.Right)
}

func main() {
	t := NewBinaryTree([]string{"1", "2", "2", "3", "4", "4", "3"})
	t.Dump()
	fmt.Println(isSymmetric2(t))
}
