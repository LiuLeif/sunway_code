// 2018-02-05 18:28
package main

import "fmt"
import (
	. "util/list"
)

func len(head *ListNode) int {
	ret := 0
	for head != nil {
		ret++
		head = head.Next
	}
	return ret
}

func intersectionNode(l1, l2 *ListNode) *ListNode {
	len1, len2 := len(l1), len(l2)
	if len1 < len2 {
		for i := 0; i < len2-len1; i++ {
			l2 = l2.Next
		}
	} else {
		for i := 0; i < len1-len2; i++ {
			l1 = l1.Next
		}
	}
	for l1 != l2 {
		l1 = l1.Next
		l2 = l2.Next
	}
	return l1
}

func main() {
	l1 := NewLinkedList([]int{1, 2, 3, 4, 5})
	l2 := NewLinkedList([]int{11, 22, 33})
	l2.Next.Next = l1.Next.Next.Next
	l1.Dump()
	l2.Dump()
	// should return 4
	fmt.Println(intersectionNode(l1, nil))
}
