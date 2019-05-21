// 2017-11-14 18:54
package main

import (
	"fmt"
	"math"

	. "util/tree"
)

func max(nums ...int) int {
	ret := math.MinInt32
	for _, n := range nums {
		if ret < n {
			ret = n
		}
	}
	return ret
}

func maxDepth(root *TreeNode) int {
	if root == nil {
		return 0
	}
	return 1 + max(maxDepth(root.Left), maxDepth(root.Right))
	return 0
}
func main() {
	t := NewBinaryTree([]string{})
	fmt.Println(maxDepth(t))
}
