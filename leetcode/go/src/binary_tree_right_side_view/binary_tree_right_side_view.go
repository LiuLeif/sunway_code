// 2018-02-05 16:24
package main

import "fmt"
import (
	. "util/tree"
)

func dfs(root *TreeNode, ret *[]int, level int) {
	if root == nil {
		return
	}
	for level >= len(*ret) {
		*ret = append(*ret, 0)
	}
	(*ret)[level] = root.Val
	dfs(root.Left, ret, level+1)
	dfs(root.Right, ret, level+1)
}

func rightSideView(root *TreeNode) []int {
	ret := []int{}
	dfs(root, &ret, 0)
	return ret
}

func main() {
	t := NewBinaryTree([]string{"1", "2", "3", "null", "5", "null", "4"})
	t.Dump()
	fmt.Println(rightSideView(t))
}
