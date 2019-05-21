// 2017-12-27 17:05
package main

import "fmt"
import (
	. "util/list"
)

func swapPairs(head *ListNode) *ListNode {
	dummyHead := &ListNode{0, head}
	prevTail, tail := dummyHead, head
	for tail != nil && tail.Next != nil {
		prevTail, tail, prevTail.Next, tail.Next.Next = tail, tail.Next.Next, tail.Next, tail
	}
	prevTail.Next = tail
	return dummyHead.Next
}
func main() {
	l := NewLinkedList([]int{1, 2, 3, 4, 5})
	fmt.Println(swapPairs(l))
}
