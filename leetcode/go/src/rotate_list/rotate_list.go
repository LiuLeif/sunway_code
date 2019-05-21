// 2017-12-22 15:50
package main

import (
	"fmt"
	. "util/list"
)

func reverse(head *ListNode) *ListNode {
	if head == nil {
		return head
	}
	var prev *ListNode = nil
	for head != nil {
		head, head.Next, prev = head.Next, prev, head
	}
	return prev
}

func rotateRight(head *ListNode, k int) *ListNode {
	if head == nil {
		return nil
	}
	len := 0
	t := head
	for t != nil {
		len++
		t = t.Next
	}
	k %= len
	if k == 0 {
		return head
	}

	head = reverse(head)
	firstHead := head
	for i := 0; i < k-1; i++ {
		head = head.Next
	}
	l2 := reverse(head.Next)
	head.Next = nil

	l1 := reverse(firstHead)
	firstHead.Next = l2
	return l1
}

func main() {
	l := NewLinkedList([]int{1, 2, 3, 4, 5})
	fmt.Println(rotateRight(l, 1))
}
