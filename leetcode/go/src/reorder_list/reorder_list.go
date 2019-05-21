// 2017-11-28 10:27
package main

import (
	. "util/list"
)

func doRecorder(earlyReturn *bool, head **ListNode, tail *ListNode) {
	if tail.Next != nil {
		doRecorder(earlyReturn, head, tail.Next)
	}
	if *earlyReturn {
		return
	}
	if (*head).Next == tail || (*head) == tail {
		tail.Next = nil
		*earlyReturn = true
	} else {
		*head, (*head).Next, tail.Next = (*head).Next, tail, (*head).Next
	}
}

func reorderList(head *ListNode) {
	if head == nil {
		return
	}
	earlyReturn := false
	doRecorder(&earlyReturn, &head, head)
}

func main() {
	// 1->2->3->4->nil
	// 1->2->3->nil
	l := NewLinkedList([]int{1, 2, 3, 4})
	reorderList(l)
	l.Dump()
}
