// 2018-12-18 13:10
package main

import (
	"fmt"
	. "util/tree"
)

type Queue []*TreeNode

func (this *Queue) Poll() *TreeNode {
	tmp := (*this)[0]
	*this = (*this)[1:]
	return tmp
}

func (this *Queue) Offer(x *TreeNode) {
	*this = append(*this, x)
}

func (this *Queue) Empty() bool {
	return len(*this) == 0
}

func (this *Queue) Peek() *TreeNode {
	return (*this)[0]
}
func isCompleteTree(root *TreeNode) bool {
	q := Queue{root}
	isTail := false
	for !q.Empty() {
		top := q.Poll()
		if isTail && top != nil {
			return false
		}
		if top == nil {
			isTail = true
		} else {
			q.Offer(top.Left)
			q.Offer(top.Right)
		}
	}
	return true
}

func main() {
	root := NewBinaryTree([]string{"1", "2", "3", "4", "null", "6"})
	fmt.Println(isCompleteTree(root))
}
