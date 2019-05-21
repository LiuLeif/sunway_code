// 2018-01-12 12:47
package main

import (
	. "util/tree"
)

func mergeTrees(t1 *TreeNode, t2 *TreeNode) *TreeNode {
	if t1 == nil && t2 == nil {
		return nil
	}
	ret := &TreeNode{}
	v1, v2 := 0, 0
	var l1, r1, l2, r2 *TreeNode = nil, nil, nil, nil

	if t1 != nil {
		v1 = t1.Val
		l1 = t1.Left
		r1 = t1.Right
	}
	if t2 != nil {
		v2 = t2.Val
		l2 = t2.Left
		r2 = t2.Right
	}
	ret.Val = v1 + v2
	ret.Left = mergeTrees(l1, l2)
	ret.Right = mergeTrees(r1, r2)
	return ret
}
func main() {
	t1 := NewBinaryTree([]string{"1", "3", "2", "5"})
	t2 := NewBinaryTree([]string{"2", "1", "3", "null", "4", "null", "7"})
	t := mergeTrees(t1, t2)
	t.Dump()
}
