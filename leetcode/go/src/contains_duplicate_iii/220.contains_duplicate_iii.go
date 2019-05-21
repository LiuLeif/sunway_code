// 2018-10-29 15:21
package main

import "fmt"
import . "util/tree"

func rangeExist(root *TreeNode, f, t int) bool {
	v, ok := ceil(root, f)
	if !ok {
		return false
	}

	return v <= t
}

func ceil(root *TreeNode, target int) (int, bool) {
	if root == nil {
		return 0, false
	}
	if target > root.Val {
		return ceil(root.Right, target)
	}
	ret, ok := ceil(root.Left, target)
	if ok {
		return ret, true
	}
	return root.Val, true
}

func remove(root *TreeNode, target int) *TreeNode {
	if root == nil {
		return nil
	}
	if root.Val == target {
		if root.Right == nil {
			return root.Left
		}
		if root.Left == nil {
			return root.Right
		}
		return add(root.Right, root.Left)
	}
	root.Left = remove(root.Left, target)
	root.Right = remove(root.Right, target)
	return root
}

func add(root *TreeNode, target *TreeNode) *TreeNode {
	if root == nil {
		return target
	}
	if target.Val >= root.Val {
		root.Right = add(root.Right, target)
	} else {
		root.Left = add(root.Left, target)
	}
	return root
}

func pretty(head *TreeNode, prefix string, isLeft bool) string {
	if head == nil {
		return ""
	}
	var ret string
	if head.Right != nil {
		newPrefix := prefix
		if isLeft {
			newPrefix += "| "
		} else {
			newPrefix += "  "
		}
		ret += pretty(head.Right, newPrefix, false)
	}

	if isLeft {
		ret += fmt.Sprintf("%s└─%d\n", prefix, head.Val)
	} else {
		ret += fmt.Sprintf("%s┌─%d\n", prefix, head.Val)
	}

	if head.Left != nil {
		newPrefix := prefix
		if isLeft {
			newPrefix += "  "
		} else {
			newPrefix += "| "
		}
		ret += pretty(head.Left, newPrefix, true)
	}
	return ret
}

func containsNearbyAlmostDuplicate(nums []int, k int, t int) bool {
	root := (*TreeNode)(nil)

	for i := 0; i < len(nums); i++ {
		if rangeExist(root, nums[i]-t, nums[i]+t) {
			return true
		}
		root = add(root, &TreeNode{nums[i], nil, nil})
		if i >= k {
			root = remove(root, nums[i-k])
		}
	}
	return false
}

func main() {
	fmt.Println(containsNearbyAlmostDuplicate([]int{7, 1, 3}, 2, 3))
}
