// 2017-11-28 10:27
package main

import (
	"fmt"
	. "util/tree"
)

const (
	DOWN = iota
	LEFT
	RIGHT
)

func inorderTraversal(root *TreeNode) []int {
	ret := []int{}
	if root == nil {
		return ret
	}
	stack := []*TreeNode{root}
	// 0 down, 1 left, 2 right
	directions := []int{DOWN}

	for len(stack) != 0 {
		top := stack[len(stack)-1]

		dict := directions[len(directions)-1]
		directions = directions[:len(directions)-1]

		if dict == DOWN {
			directions = append(directions, LEFT)
			if top.Left != nil {
				stack = append(stack, top.Left)
				directions = append(directions, DOWN)
			}
		} else if dict == LEFT {
			directions = append(directions, RIGHT)
			ret = append(ret, top.Val)
			if top.Right != nil {
				stack = append(stack, top.Right)
				directions = append(directions, DOWN)
			}
		} else if dict == RIGHT {
			stack = stack[:len(stack)-1]
		}
	}
	return ret
}

func main() {
	t := NewBinaryTree([]string{"1", "null", "2", "3"})
	fmt.Println(inorderTraversal(t))
}
