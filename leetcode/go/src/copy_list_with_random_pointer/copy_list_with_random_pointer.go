// 2018-01-16 15:45
package main

import "fmt"

type RandomListNode struct {
	Val    int
	Next   *RandomListNode
	Random *RandomListNode
}

func (this *RandomListNode) Dump() {
	ret := ""
	for this != nil {
		if this.Random != nil {
			ret += fmt.Sprintf("%d[%d] -> ", this.Val, this.Random.Val)
		} else {
			ret += fmt.Sprintf("%d -> ", this.Val)
		}

		this = this.Next
	}
	ret += "nil"
	fmt.Println(ret)
}
func NewLinkedList(Vals []int) *RandomListNode {
	var prev *RandomListNode = nil
	var root *RandomListNode = nil
	for _, v := range Vals {
		tmp := RandomListNode{v, nil, nil}
		if prev == nil {
			root = &tmp
		} else {
			prev.Next = &tmp
		}
		prev = &tmp
	}
	return root
}

type Pair struct {
	from, to *RandomListNode
}

func copyRandomListNoExtraSpace(head *RandomListNode) *RandomListNode {
	if head == nil {
		return nil
	}
	lastNode := head
	for lastNode != nil {
		copy := &RandomListNode{lastNode.Val, nil, nil}
		copy.Next = lastNode.Next
		lastNode.Next = copy
		lastNode = copy.Next
	}

	lastNode = head
	for lastNode != nil {
		if lastNode.Random != nil {
			lastNode.Next.Random = lastNode.Random.Next
		}
		lastNode = lastNode.Next.Next
	}

	lastNode = head
	dummyRet := &RandomListNode{}
	lastCopy := dummyRet

	for lastNode != nil {
		lastCopy.Next = lastNode.Next
		lastNode = lastNode.Next.Next
		lastCopy = lastCopy.Next
	}

	return dummyRet.Next
}

func copyRandomListTwoPass(head *RandomListNode) *RandomListNode {
	if head == nil {
		return nil
	}

	visited := map[*RandomListNode]*RandomListNode{}
	lastNode := head
	dummyRet := &RandomListNode{}
	lastCopy := dummyRet

	for lastNode != nil {
		copy := &RandomListNode{lastNode.Val, nil, nil}
		visited[lastNode] = copy
		lastCopy.Next = copy

		lastCopy = lastCopy.Next
		lastNode = lastNode.Next
	}

	lastNode = head
	lastCopy = dummyRet.Next
	for lastNode != nil {
		lastCopy.Random = visited[lastNode.Random]
		lastNode = lastNode.Next
		lastCopy = lastCopy.Next
	}
	return dummyRet.Next
}

func copyRandomList(head *RandomListNode) *RandomListNode {
	if head == nil {
		return nil
	}
	ret := &RandomListNode{head.Val, nil, nil}
	queue := []Pair{{head, ret}}
	visited := map[*RandomListNode]*RandomListNode{}
	visited[head] = ret
	for len(queue) != 0 {
		from, to := queue[0].from, queue[0].to
		queue = queue[1:]
		if from.Next != nil {
			if visited[from.Next] != nil {
				to.Next = visited[from.Next]
			} else {
				to.Next = &RandomListNode{from.Next.Val, nil, nil}
				visited[from.Next] = to.Next
				queue = append(queue, Pair{from.Next, to.Next})
			}
		}
		if from.Random != nil {
			if visited[from.Random] != nil {
				to.Random = visited[from.Random]
			} else {
				to.Random = &RandomListNode{from.Random.Val, nil, nil}
				visited[from.Random] = to.Random
				queue = append(queue, Pair{from.Random, to.Random})
			}
		}
	}

	return ret
}

func main() {
	l1 := NewLinkedList([]int{1, 2, 3, 4, 5})
	l1.Next.Next.Random = l1.Next
	l1.Next.Random = l1.Next.Next.Next
	l1.Dump()
	l2 := copyRandomListNoExtraSpace(l1)
	l2.Dump()
}
