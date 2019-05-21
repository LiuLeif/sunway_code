// 2018-02-13 14:58
package main

import (
	"fmt"
	"math"

	. "util/tree"
)

func min(nums ...int) int {
	ret := math.MaxInt32
	for _, n := range nums {
		if ret > n {
			ret = n
		}
	}
	return ret
}

func minDiffInBST(root *TreeNode) int {
	prev := -1
	ret := math.MaxInt32
	var inorder func(*TreeNode)
	inorder = func(root *TreeNode) {
		if root == nil {
			return
		}
		inorder(root.Left)
		if prev != -1 {
			ret = min(ret, root.Val-prev)
		}
		prev = root.Val
		inorder(root.Right)
	}

	inorder(root)
	return ret
}
func main() {
	root := NewBinaryTree([]string{"90", "69", "null", "49", "89", "null", "52"})
	root.Dump()
	fmt.Println(minDiffInBST(root))
}
