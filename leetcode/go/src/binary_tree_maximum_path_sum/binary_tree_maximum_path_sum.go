// 2017-12-19 15:43
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

func helper(root *TreeNode) (int, int) {
	if root == nil {
		return math.MinInt32, math.MinInt32
	}
	leftMaxPath, leftMaxBranch := helper(root.Left)
	rightMaxPath, rightMaxBranch := helper(root.Right)
	maxPath := max(leftMaxPath, rightMaxPath, leftMaxBranch+root.Val+rightMaxBranch, leftMaxBranch+root.Val, rightMaxBranch+root.Val, root.Val)
	maxBranch := max(leftMaxBranch+root.Val, rightMaxBranch+root.Val, root.Val)
	return maxPath, maxBranch
}

func maxPathSum(root *TreeNode) int {
	x, _ := helper(root)
	return x
}
func main() {
	t := NewBinaryTree([]string{"1", "2"})
	t.Dump()
	fmt.Println(maxPathSum(t))
}
