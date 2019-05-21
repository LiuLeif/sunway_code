// 2017-11-20 10:54
package main

import (
	"fmt"

	. "util/list"
)

// recursive approach
var origHead *ListNode

func check(tail *ListNode) bool {
	if tail == nil {
		return true
	}
	ret := check(tail.Next)

	headVal := origHead.Val
	origHead = origHead.Next

	if !ret || headVal != tail.Val {
		return false
	}
	return true
}

func isPalindrome(head *ListNode) bool {
	origHead = head
	return check(head)
}

// stack approach
func isPalindromeByStack(head *ListNode) bool {
	stack := []int{}
	origHead := head
	for head != nil {
		stack = append(stack, head.Val)
		head = head.Next
	}
	for origHead != nil {
		t := stack[len(stack)-1]
		stack = stack[:len(stack)-1]
		if t != origHead.Val {
			return false
		}
		origHead = origHead.Next
	}
	return true
}

// reverse approach
func isPalindromeByReverse(head *ListNode) bool {
	slow, fast := head, head
	for fast != nil && fast.Next != nil {
		slow = slow.Next
		fast = fast.Next.Next
	}
	// reverse [slow:]
	var prev *ListNode = nil
	curr := slow
	for curr != nil {
		tmp := curr.Next
		curr.Next = prev
		prev = curr
		curr = tmp
	}
	for head != nil && prev != nil {
		if head.Val != prev.Val {
			return false
		}
		head = head.Next
		prev = prev.Next
	}

	return true
}

func main() {
	l := NewLinkedList([]int{1, 1})
	fmt.Println(isPalindromeByStack(l))
}
