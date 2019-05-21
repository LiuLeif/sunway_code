// 2017-11-14 18:54
package main

import (
	. "util/list"
)

func mergeTwoLists(l1 *ListNode, l2 *ListNode) *ListNode {
	fakeHead := &ListNode{0, nil}
	prev := fakeHead
	for l1 != nil && l2 != nil {
		if l1.Val <= l2.Val {
			prev.Next = l1
			l1 = l1.Next
		} else {
			prev.Next = l2
			l2 = l2.Next
		}
		prev = prev.Next
	}
	if l1 == nil {
		prev.Next = l2
	} else {
		prev.Next = l1
	}
	return fakeHead.Next
}
func main() {
	l1 := NewLinkedList([]int{1})
	l2 := NewLinkedList([]int{2})
	l3 := mergeTwoLists(l1, l2)
	l3.Dump()
}
