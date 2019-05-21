// 2018-02-06 10:16
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

type InOrderIterator struct {
	stack     []*TreeNode
	direction []int
	next      int
	hasNext   bool
}

func NewInOrderIterator(root *TreeNode) *InOrderIterator {
	ret := InOrderIterator{}
	if root == nil {
		return &ret
	}
	ret.stack = []*TreeNode{root}
	ret.direction = []int{DOWN}
	ret.next, ret.hasNext = ret.getNext()
	return &ret
}

func (this *InOrderIterator) HasNext() bool {
	return this.hasNext
}

func (this *InOrderIterator) Next() int {
	ret := this.next
	this.next, this.hasNext = this.getNext()
	return ret
}

func (this *InOrderIterator) getNext() (int, bool) {
	for len(this.stack) != 0 {
		top := this.stack[len(this.stack)-1]

		dict := this.direction[len(this.direction)-1]
		this.direction = this.direction[:len(this.direction)-1]

		if dict == DOWN {
			this.direction = append(this.direction, LEFT)
			if top.Left != nil {
				this.stack = append(this.stack, top.Left)
				this.direction = append(this.direction, DOWN)
			}
		} else if dict == LEFT {
			this.direction = append(this.direction, RIGHT)
			if top.Right != nil {
				this.stack = append(this.stack, top.Right)
				this.direction = append(this.direction, DOWN)
			}
			return top.Val, true
		} else if dict == RIGHT {
			this.stack = this.stack[:len(this.stack)-1]
		}
	}
	return -1, false
}

func main() {
	t := NewBinaryTree([]string{"1", "2", "3", "4", "5"})
	t.Dump()
	iter := NewInOrderIterator(t)
	for iter.HasNext() {
		fmt.Println(iter.Next())
	}
}
