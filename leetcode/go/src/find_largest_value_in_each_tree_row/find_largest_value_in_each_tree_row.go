// 2018-02-02 13:41
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

func largestValues(root *TreeNode) []int {
	ret := []int{}
	var dfs func(*TreeNode, int)
	dfs = func(root *TreeNode, level int) {
		if root == nil {
			return
		}
		for level >= len(ret) {
			ret = append(ret, math.MinInt32)
		}
		ret[level] = max(root.Val, ret[level])
		dfs(root.Left, level+1)
		dfs(root.Right, level+1)
	}
	dfs(root, 0)
	return ret
}

func main() {
	t := NewBinaryTree([]string{"1", "3", "2", "5", "3", "null", "9"})
	fmt.Println(largestValues(t))
}
