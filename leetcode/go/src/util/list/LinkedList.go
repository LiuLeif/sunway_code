// 2017-12-19 17:32
package LinkedList

import "fmt"

type ListNode struct {
	Val  int
	Next *ListNode
}

func NewLinkedList(Vals []int) *ListNode {
	var prev *ListNode = nil
	var root *ListNode = nil
	for _, v := range Vals {
		tmp := ListNode{v, nil}
		if prev == nil {
			root = &tmp
		} else {
			prev.Next = &tmp
		}
		prev = &tmp
	}
	return root
}

func (this *ListNode) Dump() {
	ret := ""
	for this != nil {
		ret += fmt.Sprintf("%d -> ", this.Val)
		this = this.Next
	}
	ret += "nil"
	fmt.Println(ret)
}

func ReverseLinkedList(node *ListNode) *ListNode {
	var prev *ListNode = nil
	curr := node
	for curr != nil {
		tmp := curr.Next
		curr.Next = prev
		prev = curr
		curr = tmp
	}
	return prev
}
