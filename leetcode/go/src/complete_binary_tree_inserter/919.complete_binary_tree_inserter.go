// 2018-11-09 12:10
package main

import (
	"fmt"
	. "util/tree"
)

type CBTInserter struct {
	root  *TreeNode
	queue []*TreeNode
}

func Constructor(root *TreeNode) CBTInserter {
	queue := []*TreeNode{}

	tmpQueue := []*TreeNode{root}
	for len(tmpQueue) != 0 {
		top := tmpQueue[0]
		tmpQueue = tmpQueue[1:]
		if top.Left == nil || top.Right == nil {
			queue = append(queue, top)
		}
		if top.Left != nil {
			tmpQueue = append(tmpQueue, top.Left)
		}
		if top.Right != nil {
			tmpQueue = append(tmpQueue, top.Right)
		}
	}
	return CBTInserter{root, queue}
}

func (this *CBTInserter) Insert(v int) int {
	top := this.queue[0]
	node := &TreeNode{v, nil, nil}
	if top.Left == nil {
		top.Left = node
	} else {
		top.Right = node
		this.queue = this.queue[1:]
	}
	this.queue = append(this.queue, node)
	return top.Val
}

func (this *CBTInserter) Get_root() *TreeNode {
	return this.root
}

func main() {
	root := NewBinaryTree([]string{"1", "2", "3", "4", "5", "6"})
	inserter := Constructor(root)
	fmt.Println(inserter.Insert(7))
	fmt.Println(inserter.Insert(8))
	inserter.Get_root().Dump()
}
