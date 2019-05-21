// 2018-01-30 16:23
package main

import (
	. "util/list"
)

func deleteDuplicates(head *ListNode) *ListNode {
	if head == nil {
		return nil
	}
	prev := head
	curr := head.Next
	for curr != nil {
		if prev.Val == curr.Val {
			prev.Next, curr = curr.Next, curr.Next
		} else {
			prev, curr = curr, curr.Next
		}
	}
	return head
}
func main() {
	l := NewLinkedList([]int{1, 1})
	l = deleteDuplicates(l)
	l.Dump()
}
