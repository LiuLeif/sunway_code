// 2018-02-09 16:50
package main

import (
	"fmt"
	. "util/tree"
)

func sameTree(s, t *TreeNode) bool {
	if s == nil && t == nil {
		return true
	}
	if s == nil || t == nil {
		return false
	}
	if s.Val != t.Val {
		return false
	}
	return sameTree(s.Left, t.Left) && sameTree(s.Right, t.Right)
}

func isSubtree(s *TreeNode, t *TreeNode) bool {
	if sameTree(s, t) {
		return true
	}
	return s != nil && (isSubtree(s.Left, t) || isSubtree(s.Right, t))
}
func main() {
	s := NewBinaryTree([]string{"3", "4", "5", "1", "2"})
	t := NewBinaryTree([]string{"4", "1", "2"})
	s.Dump()
	t.Dump()
	fmt.Println(isSubtree(s, t))
}
