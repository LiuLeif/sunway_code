// 2018-01-04 17:04
package main

import "fmt"
import (
	. "util/tree"
)

func averageOfLevels(root *TreeNode) []float64 {
	ret := []float64{}
	queue := []*TreeNode{root, nil}
	levelSum := 0
	levelCount := 0
	for len(queue) != 1 {
		top := queue[0]
		queue = queue[1:]
		if top == nil {
			ret = append(ret, float64(levelSum)/float64(levelCount))
			levelSum, levelCount = 0, 0
			queue = append(queue, nil)
			continue
		}
		levelSum += top.Val
		levelCount += 1
		if top.Left != nil {
			queue = append(queue, top.Left)
		}
		if top.Right != nil {
			queue = append(queue, top.Right)
		}
	}
	ret = append(ret, float64(levelSum)/float64(levelCount))
	return ret
}
func main() {
	t := NewBinaryTree([]string{"3", "4", "5"})
	fmt.Println(averageOfLevels(t))
}
