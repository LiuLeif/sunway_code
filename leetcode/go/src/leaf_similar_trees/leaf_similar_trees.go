// 2018-10-17 09:52
package main

import (
	"fmt"
	. "util/tree"
)

func leafSimilar(root1 *TreeNode, root2 *TreeNode) bool {
	var getLeaf func(*TreeNode) []int
	getLeaf = func(root *TreeNode) []int {
		if root == nil {
			return []int{}
		}
		if root.Left == nil && root.Right == nil {
			return []int{root.Val}
		}
		return append(getLeaf(root.Left), getLeaf(root.Right)...)
	}

	leaf1, leaf2 := getLeaf(root1), getLeaf(root2)

	if len(leaf1) != len(leaf2) {
		return false
	}
	for i := 0; i < len(leaf1); i++ {
		if leaf1[i] != leaf2[i] {
			return false
		}
	}
	return true
}
func main() {
	t1 := NewBinaryTree([]string{"3", "5", "1", "6", "2", "9", "8", "null", "null", "7", "4"})
	t2 := NewBinaryTree([]string{"3", "5", "1", "6", "7", "4", "3", "null", "null", "null", "null", "null", "null", "9", "8"})
	fmt.Println(leafSimilar(t1, t2))
    
}
