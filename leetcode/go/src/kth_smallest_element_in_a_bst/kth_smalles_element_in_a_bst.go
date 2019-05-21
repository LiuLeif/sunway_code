// 2017-11-24 17:37
package main

import (
	"fmt"

	. "util/tree"
)

func count(root *TreeNode) int {
	if root == nil {
		return 0
	}
	return 1 + count(root.Left) + count(root.Right)
}

func kthSmallest(root *TreeNode, k int) int {
	leftCount := count(root.Left)
	if leftCount == k-1 {
		return root.Val
	}
	if leftCount > k-1 {
		return kthSmallest(root.Left, k)
	} else {
		return kthSmallest(root.Right, k-leftCount-1)
	}
}
func main() {
	t := NewBinaryTree([]string{"3", "2", "4"})
	fmt.Println(kthSmallest(t, 3))
}
