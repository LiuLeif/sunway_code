// 2018-01-16 15:09
package main

import (
	. "util/list"
)

func reverse(head *ListNode) *ListNode {
	var prev *ListNode = nil
	for head != nil {
		head.Next, prev, head = prev, head, head.Next
	}
	return prev
}

func addTwoNumbers(l1 *ListNode, l2 *ListNode) *ListNode {
	l1, l2 = reverse(l1), reverse(l2)
	ret := &ListNode{}
	curr := ret
	carry := 0
	for l1 != nil || l2 != nil {
		v1, v2 := 0, 0
		if l1 != nil {
			v1 = l1.Val
			l1 = l1.Next
		}
		if l2 != nil {
			v2 = l2.Val
			l2 = l2.Next
		}
		total := v1 + v2 + carry
		carry = total / 10
		curr.Next = &ListNode{total % 10, nil}
		curr = curr.Next
	}
	if carry != 0 {
		curr.Next = &ListNode{carry, nil}
	} else {
		curr.Next = nil
	}

	return reverse(ret.Next)
}
func main() {
	l1 := NewLinkedList([]int{7, 2, 4, 3})
	l2 := NewLinkedList([]int{5, 6, 4})
	l3 := addTwoNumbers(l1, l2)
	l3.Dump()
}
