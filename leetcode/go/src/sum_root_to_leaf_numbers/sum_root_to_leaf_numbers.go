// 2017-11-14 18:54
package main

import (
	"fmt"

	. "util/tree"
)

func walk(sum int, root *TreeNode) int {
	sum = sum*10 + root.Val
	if root.Left == nil && root.Right == nil {
		return sum
	}

	ret := 0
	if root.Left != nil {
		ret += walk(sum, root.Left)
	}
	if root.Right != nil {
		ret += walk(sum, root.Right)
	}
	return ret
}

func sumNumbers(root *TreeNode) int {
	if root == nil {
		return 0
	}
	return walk(0, root)
}
func main() {
	t := NewBinaryTree([]string{})
	fmt.Println(sumNumbers(t))
}
