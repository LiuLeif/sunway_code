// 2017-12-13 13:52
package main

import (
	"fmt"
	. "util/list"
)

func sortList(head *ListNode) *ListNode {
	if head == nil || head.Next == nil {
		return head
	}

	slow, fast := head, head.Next
	for fast != nil && fast.Next != nil {
		slow = slow.Next
		fast = fast.Next.Next
	}
	l1 := head
	l2 := slow.Next
	slow.Next = nil

	// merge l1 & l2
	l1 = sortList(l1)
	l2 = sortList(l2)
	dummHead := &ListNode{}
	curr := dummHead
	for l1 != nil || l2 != nil {
		var t *ListNode = nil
		if l1 == nil {
			t = l2
			l2 = l2.Next
		} else if l2 == nil {
			t = l1
			l1 = l1.Next
		} else {
			if l1.Val > l2.Val {
				t = l2
				l2 = l2.Next
			} else {
				t = l1
				l1 = l1.Next
			}
		}
		curr.Next = t
		curr = curr.Next
	}
	return dummHead.Next
}
func main() {
	l := NewLinkedList([]int{1, 1, 2, 0})
	fmt.Println(sortList(l))
}
