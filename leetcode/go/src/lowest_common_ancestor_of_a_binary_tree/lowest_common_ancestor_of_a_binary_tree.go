// 2017-12-13 13:52
package main

import (
	"fmt"
	"math"
	. "util/tree"
)

func lowestCommonAncestorRecursive(root, p, q *TreeNode) *TreeNode {
	if root == nil || root == q || root == p {
		return root
	}
	left := lowestCommonAncestor(root.Left, p, q)
	right := lowestCommonAncestor(root.Right, p, q)
	if left == nil {
		return right
	}
	if right == nil {
		return left
	}
	return root
}

func min(nums ...int) int {
	ret := math.MaxInt32
	for _, n := range nums {
		if ret > n {
			ret = n
		}
	}
	return ret
}

func lowestCommonAncestor(root, p, q *TreeNode) *TreeNode {
	path1 := path(root, p)
	path2 := path(root, q)
	fmt.Println(path1, path2)
	len := min(len(path1), len(path2))
	ret := root
	for i := 0; i < len; i++ {
		if path1[i] != path2[i] {
			break
		}
		ret = path1[i]
	}
	return ret
}

func path(root, p *TreeNode) []*TreeNode {
	if root == nil {
		return nil
	}
	ret := []*TreeNode{root}
	if root == p {
		return ret
	}
	left := path(root.Left, p)
	if left != nil {
		return append(ret, left...)
	}
	right := path(root.Right, p)
	if right != nil {
		return append(ret, right...)
	}
	return nil
}

func main() {
	t := NewBinaryTree([]string{"1", "2", "3", "4"})
	t.Dump()
	fmt.Println(lowestCommonAncestor(t, t.Left, t.Right))
}
