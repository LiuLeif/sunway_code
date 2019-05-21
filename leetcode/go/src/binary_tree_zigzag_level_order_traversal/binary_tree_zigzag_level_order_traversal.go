// 2017-11-14 18:54
package main

import (
	"fmt"

	. "util/tree"
)

func zigzagLevelOrder(root *TreeNode) [][]int {
	record := make([][]int, 0)
	if root == nil {
		return record
	}
	queue := make([]*TreeNode, 0)
	queue = append(queue, root, nil)
	curr := make([]int, 0)
	for len(queue) != 1 {
		t := queue[0]
		queue = queue[1:]
		if t == nil {
			record = append(record, curr)
			curr = make([]int, 0)
			queue = append(queue, nil)
		} else {
			curr = append(curr, t.Val)
			if t.Left != nil {
				queue = append(queue, t.Left)
			}
			if t.Right != nil {
				queue = append(queue, t.Right)
			}
		}
	}
	record = append(record, curr)
	for i := 1; i < len(record); i += 2 {
		l := len(record[i])
		for j := 0; j < l/2; j++ {
			record[i][j], record[i][l-j-1] = record[i][l-j-1], record[i][j]
		}
	}

	return record
}
func main() {
	// t := NewBinaryTree([]string{"3", "9", "20", "null", "null", "15", "7"})
	t := NewBinaryTree([]string{})
	fmt.Println(zigzagLevelOrder(t))
}
