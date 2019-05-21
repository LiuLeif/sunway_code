// 2017-11-14 18:54
package main

import (
	"fmt"

	. "util/tree"
)

func isSameTree(p *TreeNode, q *TreeNode) bool {
	if p == nil && q == nil {
		return true
	}
	if p == nil || q == nil {
		return false
	}

	return p.Val == q.Val && isSameTree(p.Left, q.Left) && isSameTree(p.Right, q.Right)

}
func main() {
	t1 := NewBinaryTree([]string{"1", "2", "3"})
	t2 := NewBinaryTree([]string{"1", "2"})
	fmt.Println(isSameTree(t1, t2))
}
