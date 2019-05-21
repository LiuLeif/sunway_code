// 2017-11-14 18:54
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

type Pair struct {
	node  *TreeNode
	order int
}

func widthOfBinaryTree(root *TreeNode) int {
	if root == nil {
		return 0
	}

	ret := 0
	left, right := 0, 1
	queue := []Pair{Pair{root, 1}, Pair{nil, 0}}
	for len(queue) != 1 {
		t := queue[0]
		queue = queue[1:]
		if t.node == nil {
			queue = append(queue, Pair{nil, 0})
			ret = max(ret, right-left+1)
			left = 0
			right = 0
			continue
		}

		if left == 0 {
			left = t.order
		}
		right = t.order
		if t.node.Left != nil {
			queue = append(queue, Pair{t.node.Left, t.order * 2})
		}
		if t.node.Right != nil {
			queue = append(queue, Pair{t.node.Right, t.order*2 + 1})
		}
	}
	ret = max(ret, right-left+1)
	return ret
}

func main() {
	t := NewBinaryTree([]string{"1", "2", "3"})
	fmt.Println(widthOfBinaryTree(t))
}
