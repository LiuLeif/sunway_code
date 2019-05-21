// 2018-02-09 11:31
package main

import (
	. "util/list"
)

func deleteNode(node *ListNode) {
	node.Val = node.Next.Val
	node.Next = node.Next.Next
}

func main() {
	l := NewLinkedList([]int{1, 2, 3, 4, 5})
	deleteNode(l)
	l.Dump()
}
