// 2018-01-05 13:31
package main

import "fmt"
import (
	. "util/tree"
)

func pathSum(root *TreeNode, sum int) [][]int {
	ret := [][]int{}
	if root == nil {
		return ret
	}
	if root.Left == nil && root.Right == nil && sum == root.Val {
		ret = append(ret, append([]int{}, root.Val))
		return ret
	}
	if root.Left != nil {
		left := pathSum(root.Left, sum-root.Val)
		for i := 0; i < len(left); i++ {
			ret = append(ret, append([]int{root.Val}, left[i]...))
		}
	}
	if root.Right != nil {
		right := pathSum(root.Right, sum-root.Val)
		for i := 0; i < len(right); i++ {
			ret = append(ret, append([]int{root.Val}, right[i]...))
		}
	}

	return ret
}
func main() {
	t := NewBinaryTree([]string{"5", "4", "8", "11", "null", "13", "4", "7", "2", "null", "null", "5", "1"})
	fmt.Println(pathSum(t, 22))
}
