// 2018-01-10 14:21
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

func inorder(root *TreeNode, f func(v int)) {
	if root == nil {
		return
	}
	inorder(root.Left, f)
	f(root.Val)
	inorder(root.Right, f)
}

func findMode(root *TreeNode) []int {
	ret := []int{}
	currVal, currCount, maxCount := 0, 0, 0
	inorder(root, func(v int) {
		if currVal == v {
			currCount++
		} else {
			currVal = v
			currCount = 1
		}
		maxCount = max(maxCount, currCount)
	})

	currVal, currCount = 0, 0
	inorder(root, func(v int) {
		if currVal == v {
			currCount++
		} else {
			currVal = v
			currCount = 1
		}
		if currCount == maxCount {
			ret = append(ret, v)
		}
	})
	return ret
}
func main() {
	t := NewBinaryTree([]string{"1", "null", "2", "2", "3", "null", "null", "3", "3"})
	t.Dump()
	fmt.Println(findMode(t))
}
