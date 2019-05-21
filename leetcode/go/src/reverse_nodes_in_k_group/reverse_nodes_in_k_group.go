// 2017-12-27 17:05
package main

import (
	"fmt"
	. "util/list"
)

func reverse(head *ListNode) *ListNode {
	if head == nil {
		return nil
	}
	var prev *ListNode = nil
	for head != nil {
		head.Next, head, prev = prev, head.Next, head
	}
	return prev
}

func reverseKGroup(head *ListNode, k int) *ListNode {

	t := head
	for i := 0; i < k-1; i++ {
		if t == nil {
			return head
		}
		t = t.Next
	}
	if t == nil {
		return head
	}

	next := reverseKGroup(t.Next, k)
	t.Next = nil
	newHead := reverse(head)
	head.Next = next
	return newHead
}
func main() {
	// 1 2 3 4 5 6
	head := NewLinkedList([]int{1, 2, 3, 4, 5, 6, 7, 8})
	fmt.Println(reverseKGroup(head, 3))
}
