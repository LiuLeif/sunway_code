// 2018-01-11 10:09
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

func max(nums ...int) int {
	ret := math.MinInt32
	for _, n := range nums {
		if ret < n {
			ret = n
		}
	}
	return ret
}

func helper(root *TreeNode) (ret bool, minVal int, maxVal int) {
	minVal, maxVal = root.Val, root.Val
	if root.Left != nil {
		left, leftMin, leftMax := helper(root.Left)
		if !left {
			return false, 0, 0
		}
		if root.Val <= leftMax {
			return false, 0, 0
		}
		minVal = min(minVal, leftMin)
	}
	if root.Right != nil {
		right, rightMin, rightMax := helper(root.Right)
		if !right {
			return false, 0, 0
		}
		if root.Val >= rightMin {
			return false, 0, 0
		}
		maxVal = max(maxVal, rightMax)
	}
	return true, minVal, maxVal
}

func isValidBST(root *TreeNode) bool {
	if root == nil {
		return true
	}
	ret, _, _ := helper(root)
	return ret
}
func main() {
	t := NewBinaryTree([]string{"2", "1", "3"})
	fmt.Println(isValidBST(t))
}
