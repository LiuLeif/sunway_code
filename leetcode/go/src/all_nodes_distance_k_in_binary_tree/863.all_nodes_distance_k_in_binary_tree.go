// 2018-10-26 12:00
package main

import (
	"fmt"
	. "util/tree"
)

func emplaceChildren(root *TreeNode, k int, ret *[]int) {
	if root == nil {
		return
	}
	if k == 0 {
		*ret = append(*ret, root.Val)
	}
	emplaceChildren(root.Left, k-1, ret)
	emplaceChildren(root.Right, k-1, ret)
}

func getDistanceToTarget(root, target *TreeNode, ret *[]int, k int) int {
	if root == nil {
		return -1
	}
	if root == target {
		return 0
	}
	leftDistance := getDistanceToTarget(root.Left, target, ret, k)
	if leftDistance >= 0 {
		distance := leftDistance + 1
		if k == distance {
			*ret = append(*ret, root.Val)
		}
		emplaceChildren(root.Right, k-distance-1, ret)
		return distance
	}
	rightDistance := getDistanceToTarget(root.Right, target, ret, k)
	if rightDistance >= 0 {
		distance := rightDistance + 1
		if k == distance {
			*ret = append(*ret, root.Val)
		}
		emplaceChildren(root.Left, k-distance-1, ret)
		return distance
	}
	return -1
}

func distanceK(root *TreeNode, target *TreeNode, K int) []int {
	ret := []int{}
	emplaceChildren(target, K, &ret)
	getDistanceToTarget(root, target, &ret, K)
	return ret
}

func main() {
	root := NewBinaryTree([]string{"3", "5", "1", "6", "2", "0", "8", "null", "null", "7", "4"})
	fmt.Println(distanceK(root, root.Left, 2))
}
