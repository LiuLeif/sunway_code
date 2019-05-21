// 2018-10-27 19:58
package main

import (
	. "util/tree"
)

func constructFromPrePost(pre []int, post []int) *TreeNode {
	if len(pre) == 0 {
		return nil
	}
	root := TreeNode{Val: pre[0]}
	if len(pre) == 1 {
		return &root
	}

	last := 0
	for last = 0; last < len(post); last++ {
		if post[last] == pre[1] {
			break
		}
	}

	last++
	root.Left = constructFromPrePost(pre[1:last+1], post[:last])
	root.Right = constructFromPrePost(pre[last+1:], post[last:len(post)-1])
	return &root
}

func main() {
	root := constructFromPrePost([]int{1, 2, 4, 5, 3, 6, 7}, []int{4, 5, 2, 6, 7, 3, 1})
	root.Dump()

}
