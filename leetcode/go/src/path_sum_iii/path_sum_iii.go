// 2017-11-14 18:54
package main

import (
	"fmt"
	. "util/tree"
)

func pathSumFrom(root *TreeNode, sum int) int {
	if root == nil {
		return 0
	}
	ret := 0
	if root.Val == sum {
		ret = 1
	}
	ret += pathSumFrom(root.Left, sum-root.Val) + pathSumFrom(root.Right, sum-root.Val)
	return ret
}

func pathSum(root *TreeNode, sum int) int {
	if root == nil {
		return 0
	}
	return pathSumFrom(root, sum) + pathSum(root.Left, sum) + pathSum(root.Right, sum)
}
func main() {
	t := NewBinaryTree([]string{"1", "null", "2", "null", "3", "null", "3", "null", "5"})
	t.Dump()
	fmt.Println(pathSum(t, 3))
}
