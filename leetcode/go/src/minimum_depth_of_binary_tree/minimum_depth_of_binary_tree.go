// 2017-12-05 16:02
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

func minDepth(root *TreeNode) int {
	if root == nil {
		return 0
	}
	if root.Left == nil && root.Right == nil {
		return 1
	}
	left := math.MaxInt32
	if root.Left != nil {
		left = 1 + minDepth(root.Left)
	}
	right := math.MaxInt32
	if root.Right != nil {
		right = 1 + minDepth(root.Right)
	}
	return min(left, right)
}
func main() {
	t := NewBinaryTree([]string{"1", "2", "3"})
	fmt.Println(minDepth(t))
}
