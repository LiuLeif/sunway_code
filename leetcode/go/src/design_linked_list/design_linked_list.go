// 2018-10-15 17:52
package main

import "fmt"

type Node struct {
	val  int
	next *Node
}

type MyLinkedList struct {
	head *Node
}

func Constructor() MyLinkedList {
	return MyLinkedList{head: nil}
}

func (this *MyLinkedList) Get(index int) int {
	head := this.head
	for index != 0 && head != nil {
		index--
		head = head.next
	}
	if head == nil {
		return -1
	}
	return head.val
}

func (this *MyLinkedList) AddAtHead(val int) {
	this.head = &Node{val, this.head}
}

func (this *MyLinkedList) AddAtTail(val int) {
	dummyHead := &Node{0, this.head}
	prev, head := dummyHead, this.head
	for head != nil {
		prev, head = head, head.next
	}
	prev.next = &Node{val, nil}
	this.head = dummyHead.next
}

func (this *MyLinkedList) AddAtIndex(index int, val int) {
	dummyHead := &Node{0, this.head}
	prev, head := dummyHead, this.head
	for index != 0 && head != nil {
		index--
		prev, head = head, head.next
	}
	if head == nil && index != 0 {
		return
	}
	node := &Node{val, prev.next}
	prev.next, node.next = node, prev.next
	this.head = dummyHead.next
}

func (this *MyLinkedList) DeleteAtIndex(index int) {
	dummyHead := &Node{0, this.head}
	prev, head := dummyHead, this.head
	for index != 0 && head != nil {
		index--
		prev, head = head, head.next
	}
	if head == nil {
		return
	}
	prev.next = head.next
	this.head = dummyHead.next
}

func (this *MyLinkedList) Dump() {
	head := this.head
	for head != nil {
		fmt.Printf("%d -> ", head.val)
		head = head.next
	}
	fmt.Println()
}

func main() {
	obj := Constructor()
	obj.AddAtHead(10)
	obj.AddAtTail(5)
	obj.AddAtIndex(1, 2)
	obj.DeleteAtIndex(10)
	obj.Dump()
}
