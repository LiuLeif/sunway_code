// 2017-11-27 15:46
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

func doRob(m1, m2 map[*TreeNode]int, root *TreeNode, mustSkip bool) int {
	if root == nil {
		return 0
	}
	if mustSkip && m1[root] != 0 {
		return m1[root]
	}
	if !mustSkip && m2[root] != 0 {
		return m2[root]
	}
	if mustSkip {
		x := doRob(m1, m2, root.Left, false) + doRob(m1, m2, root.Right, false)
		m1[root] = x
		return x
	} else {
		x := max(doRob(m1, m2, root.Left, true)+doRob(m1, m2, root.Right, true)+root.Val, doRob(m1, m2, root.Left, false)+doRob(m1, m2, root.Right, false))
		m2[root] = x
		return x
	}
}

func rob(root *TreeNode) int {
	m1 := make(map[*TreeNode]int)
	m2 := make(map[*TreeNode]int)
	return doRob(m1, m2, root, false)
}
func main() {
	// t := NewBinaryTree([]string{"3", "4", "5", "1", "3", "null", "1"})
	t := NewBinaryTree([]string{"4", "1", "null", "2", "null", "3"})
	fmt.Println(rob(t))
}
