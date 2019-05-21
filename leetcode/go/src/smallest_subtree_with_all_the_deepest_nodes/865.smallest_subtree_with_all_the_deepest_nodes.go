// 2018-10-26 08:25
package main

import . "util/tree"

/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
func subtreeWithAllDeepest(root *TreeNode) *TreeNode {
	var dfs func(root *TreeNode) (*TreeNode, int)
	dfs = func(root *TreeNode) (*TreeNode, int) {
		if root == nil {
			return nil, 0
		}
		leftRoot, leftLen := dfs(root.Left)
		rightRoot, rightLen := dfs(root.Right)
		if leftLen == rightLen {
			return root, leftLen + 1
		}
		if leftLen > rightLen {
			return leftRoot, leftLen + 1
		}
		return rightRoot, rightLen + 1
	}
	ret, _ := dfs(root)
	return ret
}

func main() {
	root := NewBinaryTree([]string{"3", "5", "1", "6", "2", "0", "8", "null", "null", "7", "4"})
	ret := subtreeWithAllDeepest(root)
	ret.Dump()
}
