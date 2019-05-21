// 2018-01-12 10:18
package main

import (
	"fmt"
	. "util/list"
)

func detectCycle(head *ListNode) *ListNode {
	//               +-----y---+
	//               v         |
	// 假设 1---x--->2----n--->3
	// slow, fast 在 3 相遇, 其中 slow 走了 x+n 步,
	// fast 走了 x+n+y+n-1 步 (因为 fast 的起点在 1.Next),
	// 2(x+n)=x+y+2n-1 -> x=y-1, 所以分别人 1 和 3.Next 以相同的
	// 步数走, 最终会在会合到 2
	if head == nil {
		return nil
	}
	slow, fast := head, head.Next
	for slow != fast && fast != nil && fast.Next != nil {
		slow = slow.Next
		fast = fast.Next.Next
	}
	if slow != fast {
		return nil
	}

	slow, fast = head, fast.Next
	for slow != fast {
		slow = slow.Next
		fast = fast.Next
	}

	return slow
}

func main() {
	a := &ListNode{1, nil}
	b := &ListNode{2, nil}
	c := &ListNode{3, nil}
	d := &ListNode{4, nil}
	a.Next = b
	b.Next = c
	c.Next = d
	d.Next = b
	fmt.Println(detectCycle(a))
}
