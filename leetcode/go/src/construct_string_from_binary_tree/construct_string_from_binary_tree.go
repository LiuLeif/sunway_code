// 2018-02-02 14:22
package main

import (
	"fmt"
	"strconv"

	. "util/tree"
)

func tree2str(t *TreeNode) string {
	if t == nil {
		return ""
	}
	ret := strconv.Itoa(t.Val)
	if t.Left == nil && t.Right == nil {
		return ret
	}
	ret += "(" + tree2str(t.Left) + ")"
	if t.Right != nil {
		ret += "(" + tree2str(t.Right) + ")"
	}
	return ret
}
func main() {
	t := NewBinaryTree([]string{"1", "2", "3", "4"})
	fmt.Println(tree2str(t))
}
