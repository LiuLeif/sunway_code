// 2018-01-17 18:08
package main

import (
	. "util/tree"
)

func dupTree(root *TreeNode) *TreeNode {
	if root == nil {
		return nil
	}
	return &TreeNode{root.Val, dupTree(root.Left), dupTree(root.Right)}
}

func helper(nums []int) []*TreeNode {
	if len(nums) == 0 {
		return []*TreeNode{nil}
	}
	// 1,2,3
	ret := []*TreeNode{}
	for i := 0; i < len(nums); i++ {
		root := &TreeNode{nums[i], nil, nil}
		left := helper(nums[:i])
		right := helper(nums[i+1:])
		for _, l := range left {
			for _, r := range right {
				root.Left = l
				root.Right = r
				ret = append(ret, dupTree(root))
			}
		}
	}
	return ret
}

func generateTrees(n int) []*TreeNode {
	// 1,2,3
	if n == 0 {
		return []*TreeNode{}
	}
	nums := make([]int, n)
	for i := 0; i < n; i++ {
		nums[i] = i + 1
	}
	return helper(nums)
}
func main() {
	trees := generateTrees(3)
	for i := 0; i < len(trees); i++ {
		trees[i].Dump()
	}
}
