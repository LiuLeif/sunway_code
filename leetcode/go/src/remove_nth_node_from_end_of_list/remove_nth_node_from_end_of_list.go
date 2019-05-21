// 2017-12-07 19:09
package main

import (
	"fmt"
	. "util/list"
)

func removeNthFromEnd(head *ListNode, n int) *ListNode {
	dummyHead := &ListNode{}
	dummyHead.Next = head
	slow, fast := dummyHead, head
	for i := 0; i < n; i++ {
		fast = fast.Next
	}
	for fast != nil {
		slow = slow.Next
		fast = fast.Next
	}
	slow.Next = slow.Next.Next
	return dummyHead.Next
}
func main() {
	l := NewLinkedList([]int{1, 2, 3})
	l = removeNthFromEnd(l, 1)
	fmt.Println(l)
}
