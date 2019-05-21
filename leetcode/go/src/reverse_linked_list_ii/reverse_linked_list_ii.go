// 2017-12-06 15:10
package main

import (
	. "util/list"
)

func reverseBetween(head *ListNode, m int, n int) *ListNode {
	if head == nil {
		return nil
	}
	dummyHead := &ListNode{}
	dummyHead.Next = head
	pre := dummyHead
	for i := 0; i < m-1; i++ {
		pre = pre.Next
	}

	start := pre.Next
	then := start.Next
	for i := 0; i < n-m; i++ {
		pre.Next, start.Next, then, then.Next = then, then.Next, then.Next, pre.Next
	}

	return dummyHead.Next
}
func main() {
	l := NewLinkedList([]int{1, 2, 3, 4, 5})
	l2 := reverseBetween(l, 2, 4)
	l2.Dump()
}
