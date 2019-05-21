// 2017-11-14 18:54
package main

import (
	"fmt"

	. "util/tree"
)

func findBottomLeftValue(root *TreeNode) int {
	queue := []*TreeNode{root, nil}
	ret := -1
	newRow := true
	for len(queue) != 1 {
		head := queue[0]
		queue = queue[1:]
		if newRow {
			ret = head.Val
		}
		newRow = false
		if head == nil {
			queue = append(queue, nil)
			newRow = true
		} else {
			if head.Left != nil {
				queue = append(queue, head.Left)
			}
			if head.Right != nil {
				queue = append(queue, head.Right)
			}
		}
	}
	return ret
}
func main() {
	t := NewBinaryTree([]string{"1", "3", "2"})
	fmt.Println(findBottomLeftValue(t))
}
