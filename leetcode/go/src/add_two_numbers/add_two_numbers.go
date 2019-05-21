// 2017-11-14 18:54
package main

import (
	"fmt"
	. "util/list"
)

/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */
func addTwoNumbers(l1 *ListNode, l2 *ListNode) *ListNode {
	var dummyHead *ListNode = new(ListNode)
	curr := dummyHead
	carry := 0
	for l1 != nil || l2 != nil {
		x := 0
		if l1 != nil {
			x = l1.Val
		}
		y := 0
		if l2 != nil {
			y = l2.Val
		}

		curr.Next = &ListNode{(x + y + carry) % 10, nil}
		curr = curr.Next
		carry = (x + y + carry) / 10

		if l1 != nil {
			l1 = l1.Next
		}
		if l2 != nil {
			l2 = l2.Next
		}
	}
	if carry == 1 {
		curr.Next = &ListNode{1, nil}
	}
	return dummyHead.Next
}

func main() {
	var l1 *ListNode = NewLinkedList([]int{5})
	var l2 *ListNode = NewLinkedList([]int{5})
	var l3 = addTwoNumbers(l1, l2)
	fmt.Println(l3)
}
