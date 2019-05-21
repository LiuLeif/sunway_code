// 2017-12-28 14:52
package main

import (
	"fmt"
	"math/rand"
	"time"
	. "util/list"
)

type Solution struct {
	head *ListNode
}

func Constructor(head *ListNode) Solution {
	return Solution{head}
}

func (this *Solution) GetRandom() int {
	head := this.head
	ret := head
	n := 1
	for head != nil {
		x := rand.Intn(n)
		if x == 0 {
			ret = head
		}
		head = head.Next
		n++
	}
	return ret.Val
}

func main() {
	rand.Seed(time.Now().Unix())
	obj := Constructor(NewLinkedList([]int{1, 2}))
	fmt.Println(obj.GetRandom())
}
