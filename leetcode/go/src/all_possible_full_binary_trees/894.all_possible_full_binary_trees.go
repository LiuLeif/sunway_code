// 2018-11-01 19:17
package main

import (
	. "util/tree"
)

func allPossibleFBT(N int) []*TreeNode {
	if N&1 == 0 {
		return []*TreeNode{}
	}
	ret := []*TreeNode{}
	if N == 0 || N == 2 {
		return []*TreeNode{}
	}
	if N == 1 {
		return []*TreeNode{&TreeNode{0, nil, nil}}
	}
	if N == 3 {
		return []*TreeNode{&TreeNode{0, &TreeNode{0, nil, nil}, &TreeNode{0, nil, nil}}}
	}

	for i := 0; i < N; i++ {
		right := allPossibleFBT(i)
		left := allPossibleFBT(N - 1 - i)
		for _, x := range left {
			for _, y := range right {
				ret = append(ret, &TreeNode{0, x, y})
			}
		}
	}
	return ret
}

func main() {
	nodes := allPossibleFBT(17)
	for _, node := range nodes {
		node.Dump()
	}
}
