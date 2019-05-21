// 2018-01-29 12:13
package main

import (
	"fmt"
	"math"
	. "util/tree"
)

func helper(m map[int]int, root *TreeNode) int {
	if root == nil {
		return 0
	}

	if root.Left == nil && root.Right == nil {
		m[root.Val]++
		return root.Val
	}
	ret := root.Val + helper(m, root.Left) + helper(m, root.Right)
	m[ret]++
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

func findFrequentTreeSum(root *TreeNode) []int {
	ret := []int{}
	m := map[int]int{}
	helper(m, root)
	maxFreq := 0
	for _, v := range m {
		maxFreq = max(maxFreq, v)
	}
	for k, v := range m {
		if v == maxFreq {
			ret = append(ret, k)
		}
	}
	return ret
}
func main() {
	t := NewBinaryTree([]string{"5", "2", "-3"})
	fmt.Println(findFrequentTreeSum(t))
}
