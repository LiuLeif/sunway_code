// 2017-12-26 16:09
package main

import "fmt"
import (
	. "util/list"
)

func insertionSortList(head *ListNode) *ListNode {
	dummyHead := &ListNode{0, nil}
	dummyHead.Next = nil
	for head != nil {
		t := dummyHead
		for ; t.Next != nil && t.Next.Val <= head.Val; t = t.Next {
		}
		t.Next, head, head.Next = head, head.Next, t.Next
	}
	return dummyHead.Next
}
func main() {
	l := NewLinkedList([]int{1, 2, 1})
	fmt.Println(insertionSortList(l))
}
