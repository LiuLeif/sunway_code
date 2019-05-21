// 2018-10-16 11:13
package main

import (
	. "util/list"
)

func middleNode(head *ListNode) *ListNode {
	slow, fast := head, head

	for slow != nil && fast != nil && fast.Next != nil {
		slow = slow.Next
		fast = fast.Next
		if fast != nil {
			fast = fast.Next
		}
	}
	return slow
}
func main() {
	head := NewLinkedList([]int{1, 2, 3, 4})
	head.Dump()
	m := middleNode(head)
	m.Dump()
}
