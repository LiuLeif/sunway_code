// 2018-02-06 17:55
package main

import (
	"fmt"
	. "util/tree"
)

func lowestCommonAncestor(root, a, b *TreeNode) *TreeNode {
	if a.Val > root.Val && b.Val > root.Val {
		return lowestCommonAncestor(root.Right, a, b)
	}

	if a.Val < root.Val && b.Val < root.Val {
		return lowestCommonAncestor(root.Left, a, b)
	}

	return root
}

// assume no duplicates in the BST
func main() {
	t := NewBinaryTree([]string{"2", "1", "3"})
	t.Dump()
	fmt.Println(lowestCommonAncestor(t, t.Left, t.Right))
}
