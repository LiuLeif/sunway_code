// 2017-12-06 10:40
package main

import "fmt"

type Node struct {
	Key, Val   int
	Prev, Next *Node
}

func MakeNode(key, val int) *Node {
	this := Node{}
	this.Prev = &this
	this.Next = &this
	this.Key = key
	this.Val = val
	return &this
}

type LRUCache struct {
	Capacity   int
	Count      int
	Mapping    map[int]*Node
	Head, Tail *Node
}

func Constructor(capacity int) LRUCache {
	this := LRUCache{}
	this.Capacity = capacity
	this.Count = 0
	this.Mapping = map[int]*Node{}
	this.Head = MakeNode(0, 0)
	this.Tail = MakeNode(0, 0)
	this.Head.Next = this.Tail
	this.Tail.Prev = this.Head
	return this
}

func (this *LRUCache) AddNode(node *Node) {
	this.Head.Next, node.Next, this.Head.Next.Prev, node.Prev = node, this.Head.Next, node, this.Head
}

func (this *LRUCache) RemoveNode(node *Node) {
	node.Prev.Next, node.Next.Prev = node.Next, node.Prev
}

func (this *LRUCache) Get(key int) int {
	ret := this.Mapping[key]
	if ret == nil {
		return -1
	}
	this.RemoveNode(ret)
	this.AddNode(ret)
	return ret.Val
}

func (this *LRUCache) Put(key int, value int) {
	node := this.Mapping[key]
	if node != nil {
		node.Val = value
		this.RemoveNode(node)
		this.AddNode(node)
		return
	}

	node = MakeNode(key, value)
	this.Mapping[key] = node
	this.Count++
	this.AddNode(node)

	if this.Count == this.Capacity+1 {
		delete(this.Mapping, this.Tail.Prev.Key)
		this.RemoveNode(this.Tail.Prev)
		this.Count--
	}
}

func main() {
	lru := Constructor(2)
	lru.Put(1, 1)
	lru.Put(2, 2)
	fmt.Println(lru.Get(1))
	lru.Put(3, 3)
	fmt.Println(lru.Get(2))
}
