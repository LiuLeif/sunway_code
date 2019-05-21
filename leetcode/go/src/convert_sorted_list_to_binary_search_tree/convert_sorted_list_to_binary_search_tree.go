// 2017-12-08 16:55
package main

import (
	"fmt"
	. "util/list"
	. "util/tree"
)

func sortedListToBST(head *ListNode) *TreeNode {
	if head == nil {
		return nil
	}
	slow, fast := head, head
	var prev *ListNode = nil
	for fast != nil && fast.Next != nil {
		prev = slow
		slow = slow.Next
		fast = fast.Next.Next
	}

	ret := &TreeNode{slow.Val, nil, nil}
	if prev != nil {
		prev.Next = nil
	} else {
		head = nil
	}

	ret.Left = sortedListToBST(head)
	ret.Right = sortedListToBST(slow.Next)

	return ret
}
func main() {
	l := NewLinkedList([]int{1, 2, 3})
	t := sortedListToBST(l)
	fmt.Println(t)
}
