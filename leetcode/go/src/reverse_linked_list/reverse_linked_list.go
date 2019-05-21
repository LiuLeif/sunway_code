// 2017-11-14 18:54
package main

import (
	. "util/list"
)

func reverseList2(head *ListNode) *ListNode {
	// 1 2 3 4
	if head == nil {
		return nil
	}
	dummyHead := &ListNode{}
	dummyHead.Next = head
	end := head
	current := head.Next
	for current != nil {
		dummyHead.Next, current, current.Next, end.Next = current, current.Next, dummyHead.Next, current.Next
	}
	return dummyHead.Next
}

func reverseList(head *ListNode) *ListNode {
	var prev *ListNode
	for head != nil {
		tmp := head.Next
		head.Next = prev
		prev = head
		head = tmp
	}
	return prev
}

func reverseListRecursive(head *ListNode) (*ListNode, *ListNode) {
	if head == nil {
		return nil, nil
	}

	if head.Next == nil {
		return head, head
	}

	h, t := reverseListRecursive(head.Next)
	head.Next = nil
	t.Next = head
	return h, head
}

func main() {
	l := NewLinkedList([]int{1, 2, 3, 4})
	l2 := reverseList2(l)
	l2.Dump()
}
