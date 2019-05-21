// 2018-01-12 10:01
package main

import (
	"fmt"
	. "util/list"
)

func hasCycle(head *ListNode) bool {
	if head == nil {
		return false
	}
	slow, fast := head, head.Next
	for slow != fast && fast != nil && fast.Next != nil {
		slow = slow.Next
		fast = fast.Next.Next
	}
	return slow == fast
}

func main() {
	// 1 2 3 4
	a := &ListNode{1, nil}
	b := &ListNode{2, nil}
	c := &ListNode{3, nil}
	d := &ListNode{4, nil}
	a.Next = b
	b.Next = c
	c.Next = d
	d.Next = c
	fmt.Println(hasCycle(a))
}
