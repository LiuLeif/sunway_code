// 2018-02-07 15:10
package main

import (
	"container/heap"
	"fmt"
)

type MinHeap []*Node

func (h MinHeap) Len() int {
	return len(h)
}

func (h MinHeap) Less(i, j int) bool {
	if h[i].Count < h[j].Count {
		return true
	}
	if h[i].Count > h[j].Count {
		return false
	}
	return h[i].Stamp < h[j].Stamp
}

func (h MinHeap) Swap(i, j int) {
	h[i].Index = j
	h[j].Index = i
	h[i], h[j] = h[j], h[i]
}

func (ph *MinHeap) Pop() interface{} {
	t := (*ph)[len(*ph)-1]
	*ph = (*ph)[:len(*ph)-1]
	return t
}

func (ph *MinHeap) Push(x interface{}) {
	item := x.(*Node)
	item.Index = len(*ph)
	*ph = append(*ph, item)
}

type Node struct {
	Key, Val int
	Count    int
	Index    int
	Stamp    int
}

func MakeNode(key, val int) *Node {
	this := Node{}
	this.Key = key
	this.Val = val
	this.Count = 1
	return &this
}

type LFUCache struct {
	Capacity int
	Count    int
	Mapping  map[int]*Node
	Heap     MinHeap
	Stamp    int
}

func Constructor(capacity int) LFUCache {
	this := LFUCache{}
	this.Capacity = capacity
	this.Count = 0
	this.Mapping = map[int]*Node{}
	this.Heap = MinHeap{}
	this.Stamp = 1
	return this
}

func (this *LFUCache) Get(key int) int {
	if this.Capacity == 0 {
		return -1
	}
	this.Stamp++
	node := this.Mapping[key]
	if node == nil {
		return -1
	}
	node.Count++
	node.Stamp = this.Stamp
	heap.Fix(&this.Heap, node.Index)
	return node.Val
}

func (this *LFUCache) Put(key int, value int) {
	if this.Capacity == 0 {
		return
	}
	this.Stamp++
	node := this.Mapping[key]
	if node != nil {
		node.Val = value
		node.Count++
		node.Stamp = this.Stamp
		heap.Fix(&this.Heap, node.Index)
		return
	}

	if this.Count == this.Capacity {
		top := heap.Pop(&this.Heap).(*Node)
		this.Count--
		delete(this.Mapping, top.Key)
	}

	node = MakeNode(key, value)
	node.Stamp = this.Stamp
	this.Mapping[key] = node
	this.Count++
	heap.Push(&this.Heap, node)

}

func main() {
	obj := Constructor(2)
	obj.Put(1, 1)
	obj.Put(2, 2)
	fmt.Println(obj.Get(1)) // 1
	obj.Put(3, 3)
	fmt.Println(obj.Get(2)) // -1
	fmt.Println(obj.Get(3)) // 3
	obj.Put(4, 4)
	fmt.Println(obj.Get(1)) // -1
	fmt.Println(obj.Get(3))
	fmt.Println(obj.Get(4))
}
