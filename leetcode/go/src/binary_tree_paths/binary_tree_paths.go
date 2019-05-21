// 2018-01-10 10:33
package main

import (
	"fmt"
	"strconv"

	. "util/tree"
)

func binaryTreePaths(root *TreeNode) []string {
	if root == nil {
		return nil
	}
	left := binaryTreePaths(root.Left)
	right := binaryTreePaths(root.Right)
	if left == nil && right == nil {
		return []string{strconv.Itoa(root.Val)}
	}
	ret := []string{}
	if left != nil {
		for i := 0; i < len(left); i++ {
			ret = append(ret, strconv.Itoa(root.Val)+"->"+left[i])
		}
	}
	if right != nil {
		for i := 0; i < len(right); i++ {
			ret = append(ret, strconv.Itoa(root.Val)+"->"+right[i])
		}
	}
	return ret
}
func main() {
	t := NewBinaryTree([]string{"1", "2", "3", "null", "5"})
	t.Dump()
	fmt.Println(binaryTreePaths(t))
}
