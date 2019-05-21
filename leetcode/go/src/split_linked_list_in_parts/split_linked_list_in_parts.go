// 2018-01-22 10:19
package main

import (
	. "util/list"
)

func getLen(root *ListNode) int {
	ret := 0
	for root != nil {
		ret++
		root = root.Next
	}
	return ret
}

func split(len, k int) []int {
	ret := make([]int, k)
	part := len / k
	mod := len % k
	for i := 0; i < k; i++ {
		if mod != 0 {
			ret[i] = part + 1
			mod--
		} else {
			ret[i] = part
		}
	}
	return ret
}

func splitListToParts(root *ListNode, k int) []*ListNode {
	ret := []*ListNode{}
	len := getLen(root)
	for _, p := range split(len, k) {
		if p == 0 {
			ret = append(ret, nil)
			continue
		}
		ret = append(ret, root)
		for i := 0; i < p-1; i++ {
			root = root.Next
		}
		root, root.Next = root.Next, nil
	}
	return ret
}
func main() {
	head := NewLinkedList([]int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10})
	parts := splitListToParts(head, 3)
	for _, p := range parts {
		p.Dump()
	}
}
