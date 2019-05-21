// 2018-02-06 17:18
package main

import (
	"fmt"

	. "util/tree"
)

func pow(x, y int) int {
	ret := 1
	for i := 0; i < y; i++ {
		ret *= x
	}
	return ret
}

func rightHeight(root *TreeNode) int {
	ret := 0
	for root != nil {
		root = root.Right
		ret++
	}
	return ret
}
func leftHeight(root *TreeNode) int {
	ret := 0
	for root != nil {
		root = root.Left
		ret++
	}
	return ret
}

func countNodes(root *TreeNode) int {
	if root == nil {
		return 0
	}
	leftHeight, rightHeight := leftHeight(root.Left), rightHeight(root.Right)
	if leftHeight == rightHeight {
		return pow(2, leftHeight+1) - 1
	}
	return countNodes(root.Left) + countNodes(root.Right) + 1
}

func main() {
	t := NewBinaryTree([]string{"1", "2", "3", "4", "5", "6"})
	t.Dump()
	fmt.Println(countNodes(t))
}
