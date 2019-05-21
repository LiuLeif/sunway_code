// 2018-01-09 10:45
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

func min(nums ...int) int {
	ret := math.MaxInt32
	for _, n := range nums {
		if ret > n {
			ret = n
		}
	}
	return ret
}

func findSecondMinimumValue(root *TreeNode) int {
	if root == nil || root.Left == nil {
		return -1
	}
	l := root.Left.Val
	if root.Val == root.Left.Val {
		l = findSecondMinimumValue(root.Left)
	}
	r := root.Right.Val
	if root.Val == root.Right.Val {
		r = findSecondMinimumValue(root.Right)
	}

	if l == -1 && r == -1 {
		return -1
	}
	if l != -1 && r != -1 {
		return min(l, r)
	}
	return max(l, r)
}
func main() {
	t := NewBinaryTree([]string{"5", "6", "6"})
	t.Dump()
	fmt.Println(findSecondMinimumValue(t))
}
