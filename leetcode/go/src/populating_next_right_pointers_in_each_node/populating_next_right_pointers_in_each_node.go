// 2018-01-15 11:11
package main

import (
	"fmt"
	"strconv"
)

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
	Next  *TreeNode
}

func NewBinaryTree(nodes []string) *TreeNode {
	queue := make([]*TreeNode, 0)
	if len(nodes) == 0 {
		return nil
	}
	x, _ := strconv.Atoi(nodes[0])
	root := &TreeNode{x, nil, nil, nil}
	i := 1
	queue = append(queue, root)
	for i < len(nodes) {
		p := queue[0]
		queue = queue[1:]
		if nodes[i] != "null" {
			x, _ := strconv.Atoi(nodes[i])
			tmp := &TreeNode{x, nil, nil, nil}
			p.Left = tmp
			queue = append(queue, tmp)
		}
		i += 1
		if i < len(nodes) && nodes[i] != "null" {
			x, _ := strconv.Atoi(nodes[i])
			tmp := &TreeNode{x, nil, nil, nil}
			p.Right = tmp
			queue = append(queue, tmp)
		}
		i += 1
	}
	return root
}

func dump(root *TreeNode) {
	ret := ""
	for root != nil {
		ret += fmt.Sprintf("%d -> ", root.Val)
		root = root.Next
	}
	ret += "nil"
	fmt.Println(ret)
}

func connect(root *TreeNode) {
	queue := []*TreeNode{root, nil}
	var prev *TreeNode = nil
	for len(queue) != 1 {
		top := queue[0]
		queue = queue[1:]
		if top == nil {
			prev = nil
			queue = append(queue, nil)
			continue
		}
		if top.Left != nil {
			queue = append(queue, top.Left)
		}
		if top.Right != nil {
			queue = append(queue, top.Right)
		}

		if prev == nil {
			prev = top
			continue
		}
		prev.Next = top
		prev = top
	}
}

func helper(a, b *TreeNode) {
	if a == nil || b == nil {
		return
	}
	a.Next = b
	helper(a.Left, a.Right)
	helper(a.Right, b.Left)
	helper(b.Left, b.Right)
}

func connectRecursive(root *TreeNode) {
	if root == nil {
		return
	}
	helper(root.Left, root.Right)
}

func main() {
	t := NewBinaryTree([]string{"1", "2", "3", "4", "5", "6", "7"})
	// connect(t)
	connectRecursive(t)
	dump(t.Left.Left)
}
