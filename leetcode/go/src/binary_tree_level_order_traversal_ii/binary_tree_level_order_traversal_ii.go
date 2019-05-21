// 2018-01-31 13:23
package main

import "fmt"
import (
	. "util/tree"
)

func levelOrderBottom(root *TreeNode) [][]int {
	ret := [][]int{}
	if root == nil {
		return ret
	}
	queue := []*TreeNode{root, nil}
	curr := []int{}
	for len(queue) != 1 {
		top := queue[0]
		queue = queue[1:]
		if top == nil {
			queue = append(queue, nil)
			ret = append(ret, curr)
			curr = []int{}
			continue
		}
		curr = append(curr, top.Val)
		if top.Left != nil {
			queue = append(queue, top.Left)
		}
		if top.Right != nil {
			queue = append(queue, top.Right)
		}
	}
	ret = append(ret, curr)
	for i := 0; i < len(ret)/2; i++ {
		ret[i], ret[len(ret)-i-1] = ret[len(ret)-i-1], ret[i]
	}
	return ret
}
func main() {
	root := NewBinaryTree([]string{"3", "9", "20", "null", "null", "15", "7"})
	root.Dump()
	fmt.Println(levelOrderBottom(root))
}
