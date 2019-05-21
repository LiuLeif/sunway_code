// 2017-12-14 16:21
package main

import (
	"fmt"
	. "util/list"
)

func deleteDuplicates(head *ListNode) *ListNode {
	dummyHead := &ListNode{}
	dummyHead.Next = head
	prev := dummyHead
	for head != nil {
		for head.Next != nil && head.Val == head.Next.Val {
			head = head.Next
		}
		if prev.Next == head {
			prev = head
		} else {
			prev.Next = head.Next
		}
		head = head.Next
	}
	return dummyHead.Next
}
func main() {
	l := NewLinkedList([]int{1, 2, 2, 3})
	fmt.Println(deleteDuplicates(l))
}
