// 2018-11-12 15:00
package main

import (
	"fmt"
	. "util/tree"
)

func rangeSumBST(root *TreeNode, L int, R int) int {
	if root == nil {
		return 0
	}
	ret := 0
	if root.Val > R {
		ret += rangeSumBST(root.Left, L, R)
	} else if root.Val < L {
		ret += rangeSumBST(root.Right, L, R)
	} else {
		ret += root.Val + rangeSumBST(root.Left, L, root.Val) + rangeSumBST(root.Right, root.Val, R)
	}
	return ret
}

func main() {
	root := NewBinaryTree([]string{"10", "5", "15", "3", "7", "13", "18", "1", "null", "6"})
	fmt.Println(rangeSumBST(root, 6, 10))
}
