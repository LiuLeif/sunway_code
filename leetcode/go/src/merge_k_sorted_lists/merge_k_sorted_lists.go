// 2018-01-12 10:41
package main

import (
	"container/heap"
	. "util/list"
)

type Pair struct {
	v    int
	node *ListNode
}

type MinHeap []Pair

func (h MinHeap) Len() int {
	return len(h)
}

func (h MinHeap) Less(i, j int) bool {
	return h[i].v < h[j].v
}

func (h MinHeap) Swap(i, j int) {
	h[i], h[j] = h[j], h[i]
}

func (ph *MinHeap) Pop() interface{} {
	t := (*ph)[len(*ph)-1]
	*ph = (*ph)[:len(*ph)-1]
	return t
}

func (ph *MinHeap) Push(x interface{}) {
	*ph = append(*ph, x.(Pair))
}

func mergeKLists(lists []*ListNode) *ListNode {
	dummyHead := &ListNode{}
	curr := dummyHead
	queue := MinHeap{}
	for i := 0; i < len(lists); i++ {
		if lists[i] != nil {
			heap.Push(&queue, Pair{lists[i].Val, lists[i]})
		}
	}
	for len(queue) != 0 {
		tmp := heap.Pop(&queue).(Pair).node
		curr.Next = tmp
		curr = tmp
		if tmp.Next != nil {
			heap.Push(&queue, Pair{tmp.Next.Val, tmp.Next})
		}
	}
	return dummyHead.Next
}
func main() {
	l1 := NewLinkedList([]int{1, 2, 3, 4, 5})
	l2 := NewLinkedList([]int{3, 4, 5})
	l3 := NewLinkedList([]int{3, 4, 6, 7, 8})
	mergedList := mergeKLists([]*ListNode{l1, l2, l3})
	mergedList.Dump()
}
