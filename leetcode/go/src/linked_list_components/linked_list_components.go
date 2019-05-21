// 2018-10-18 11:22
package main

import (
	"fmt"
	. "util/list"
)

func numComponents(head *ListNode, G []int) int {
	counter := map[int]bool{}
	for _, v := range G {
		counter[v] = true
	}

	recording := false
	ret := 0
	for head != nil {
		if counter[head.Val] {
			recording = true
		} else {
			if recording {
				ret++
			}
			recording = false
		}
		head = head.Next
	}
	if recording {
		ret++
	}
	return ret
}
func main() {
	head := NewLinkedList([]int{0, 1, 2, 3})
	fmt.Println(numComponents(head, []int{0, 1}))
}
