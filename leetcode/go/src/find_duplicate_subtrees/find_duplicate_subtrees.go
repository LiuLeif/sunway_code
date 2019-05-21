// 2018-01-08 13:47
package main

import "fmt"
import (
	. "util/tree"
)

func inorder(root *TreeNode, m map[string]int, ret *[]*TreeNode) string {
	if root == nil {
		return "#"
	}
	s := fmt.Sprintf("%d,%s,%s", root.Val, inorder(root.Left, m, ret), inorder(root.Right, m, ret))
	if m[s] == 1 {
		*ret = append(*ret, root)
	}
	m[s] += 1
	return s
}

func findDuplicateSubtrees(root *TreeNode) []*TreeNode {
	ret := []*TreeNode{}
	inorder(root, map[string]int{}, &ret)
	return ret
}

func main() {
	// 4 2 4 2 4 1 1
	root := NewBinaryTree([]string{"1", "2", "3", "4", "null", "2", "4", "null", "null", "4"})
	root.Dump()
	fmt.Println(findDuplicateSubtrees(root))
}
