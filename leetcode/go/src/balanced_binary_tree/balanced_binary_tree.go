// 2017-12-07 19:09
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

func dfs(root *TreeNode) (bool, int) {
	if root == nil {
		return true, 0
	}
	balancedLeft, heightLeft := dfs(root.Left)
	balancedRight, heightRight := dfs(root.Right)
	if !balancedLeft || !balancedRight {
		return false, 0
	}
	delta := heightLeft - heightRight
	if delta < 0 {
		delta = -delta
	}
	if delta > 1 {
		return false, 0
	}
	return true, max(heightLeft, heightRight) + 1
}

func isBalanced(root *TreeNode) bool {
	ret, _ := dfs(root)
	return ret
}
func main() {
	t := NewBinaryTree([]string{"1", "2", "null", "3"})
	fmt.Println(isBalanced(t))
}
