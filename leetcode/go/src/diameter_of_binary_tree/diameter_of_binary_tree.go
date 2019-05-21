// 2017-11-20 16:14
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

func length(ret *int, root *TreeNode) int {
	if root == nil {
		return 0
	}
	left := length(ret, root.Left)
	right := length(ret, root.Right)
	*ret = max(*ret, left+right)
	return max(left, right) + 1
}

func diameterOfBinaryTree(root *TreeNode) int {
	ret := 0
	length(&ret, root)
	return ret
}
func main() {
	t := NewBinaryTree([]string{"1", "2", "3", "4", "5"})
	fmt.Println(diameterOfBinaryTree(t))
}
