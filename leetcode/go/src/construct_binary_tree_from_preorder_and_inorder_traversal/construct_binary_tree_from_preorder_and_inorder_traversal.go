// 2018-01-02 13:51
package main

import (
	. "util/tree"
)

func buildTree(preorder []int, inorder []int) *TreeNode {
	if len(preorder) == 0 {
		return nil
	}
	rootVal := preorder[0]
	leftInorder, rightInorder := inorder, inorder
	for i := 0; i < len(inorder); i++ {
		if inorder[i] != rootVal {
			continue
		}
		leftInorder = inorder[:i]
		rightInorder = inorder[i+1:]
		break
	}
	leftPreorder := preorder[1 : 1+len(leftInorder)]
	rightPreorder := preorder[1+len(leftInorder):]
	return &TreeNode{rootVal, buildTree(leftPreorder, leftInorder), buildTree(rightPreorder, rightInorder)}
}
func main() {
	t := buildTree([]int{1, 2, 3}, []int{2, 1, 3})
	t.Dump()
}
