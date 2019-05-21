// 2018-01-02 13:51
package main

import (
	"fmt"
	. "util/list"
)

func partition(head *ListNode, x int) *ListNode {
	l1, l2 := &ListNode{}, &ListNode{}
	t1, t2 := l1, l2
	for head != nil {
		if head.Val < x {
			t1.Next, t1 = head, head
		} else {
			t2.Next, t2 = head, head
		}
		head = head.Next
	}
	t2.Next = nil
	t1.Next = l2.Next
	return l1.Next
}
func main() {
	l := NewLinkedList([]int{1, 4, 3, 2, 5, 2})
	fmt.Println(partition(l, 4))
}
