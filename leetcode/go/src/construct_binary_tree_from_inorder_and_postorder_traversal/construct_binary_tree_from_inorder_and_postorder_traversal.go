// 2018-01-22 13:26
package main

import (
	. "util/tree"
)

func buildTree(inorder []int, postorder []int) *TreeNode {
	if len(inorder) == 0 {
		return nil
	}
	rootVal := postorder[len(postorder)-1]
	i := 0
	for ; i < len(inorder); i++ {
		if inorder[i] == rootVal {
			break
		}
	}

	left := buildTree(inorder[:i], postorder[:i])
	right := buildTree(inorder[i+1:], postorder[i:len(postorder)-1])
	return &TreeNode{rootVal, left, right}
}
func main() {
	//
	//        6
	//    5
	// 1
	//       4
	//    2
	//       3
	t := buildTree([]int{3, 2, 4, 1, 5, 6}, []int{3, 4, 2, 6, 5, 1})
	t.Dump()
}
