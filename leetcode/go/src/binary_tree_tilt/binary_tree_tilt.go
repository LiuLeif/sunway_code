// 2018-01-09 13:47
package main

import (
	"fmt"
	. "util/tree"
)

func helper(root *TreeNode) (sum int, tilt int) {
	if root == nil {
		return 0, 0
	}
	ls, lt := helper(root.Left)
	rs, rt := helper(root.Right)
	sum = ls + rs + root.Val
	tilt = ls - rs
	if tilt < 0 {
		tilt = -tilt
	}
	tilt += lt + rt
	return
}

func findTilt(root *TreeNode) int {
	_, tilt := helper(root)
	return tilt
}
func main() {
	t := NewBinaryTree([]string{"1", "2", "4"})
	t.Dump()
	fmt.Println(findTilt(t))
}
