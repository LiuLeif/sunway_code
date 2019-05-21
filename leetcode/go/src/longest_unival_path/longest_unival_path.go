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

func walk(root *TreeNode, ret *int) int {
	if root == nil {
		return 0
	}
	left := walk(root.Left, ret)
	right := walk(root.Right, ret)
	if left == 0 || root.Left.Val != root.Val {
		left = 0
	}
	if right == 0 || root.Right.Val != root.Val {
		right = 0
	}
	*ret = max(*ret, 1+left+right)
	return 1 + max(left, right)
}

func longestUnivaluePath(root *TreeNode) int {
	if root == nil {
		return 0
	}
	ret := 0
	walk(root, &ret)
	return ret - 1
}
func main() {
	// t := NewBinaryTree([]string{"1", "4", "5", "4", "4", "5"})
	// t := NewBinaryTree([]string{"5", "4", "5", "1", "1", "null", "5"})
	t := NewBinaryTree([]string{"5", "5", "5", "5"})
	fmt.Println(longestUnivaluePath(t))
}
