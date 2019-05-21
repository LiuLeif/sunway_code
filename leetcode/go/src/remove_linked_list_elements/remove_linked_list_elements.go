// 2017-11-14 18:54
package main

import (
	. "util/list"
)

func removeElements(head *ListNode, val int) *ListNode {
	dummyHead := &ListNode{0, nil}
	prev := dummyHead
	curr := head
	prev.Next = curr
	for curr != nil {
		if curr.Val == val {
			prev.Next = curr.Next
		} else {
			prev = curr
		}
		curr = curr.Next
	}
	return dummyHead.Next
}
func main() {
	l := NewLinkedList([]int{1, 1, 2, 1})
	l = removeElements(l, 1)
	l.Dump()
}
