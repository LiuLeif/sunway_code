// 2017-12-06 15:10
package main

import (
	. "util/list"
)

func oddEvenList(head *ListNode) *ListNode {
	odd := true
	oddHead := &ListNode{}
	oddTail := oddHead
	evenHead := &ListNode{}
	evenTail := evenHead
	for head != nil {
		if odd {
			oddTail.Next = head
			oddTail = head
			odd = false
		} else {
			evenTail.Next = head
			evenTail = head
			odd = true
		}
		head = head.Next
	}
	oddTail.Next = evenHead.Next
	evenTail.Next = nil
	return oddHead.Next
}
func main() {
	l := NewLinkedList([]int{1, 2, 3, 4, 5, 6})
	l2 := oddEvenList(l)
	l2.Dump()
}
