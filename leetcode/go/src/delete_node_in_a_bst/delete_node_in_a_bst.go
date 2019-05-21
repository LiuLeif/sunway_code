// 2018-01-19 13:11
package main

import (
	. "util/tree"
)

func inorder(parent, root *TreeNode, key int) (*TreeNode, *TreeNode) {
	if root == nil {
		return nil, nil
	}
	if root.Val == key {
		return parent, root
	}
	p, r := inorder(root, root.Left, key)
	if p != nil {
		return p, r
	}
	return inorder(root, root.Right, key)
}

func minNodeOf(root *TreeNode) *TreeNode {
	if root.Left == nil {
		return root
	}
	return minNodeOf(root.Left)
}

func deleteRoot(root *TreeNode) *TreeNode {
	if root.Left == nil && root.Right == nil {
		return nil
	}
	if root.Left == nil {
		return root.Right
	}
	if root.Right == nil {
		return root.Left
	}
	minNodeOf(root.Right).Left = root.Left
	return root.Right
}

func deleteNode(root *TreeNode, key int) *TreeNode {
	dummyHead := &TreeNode{}
	dummyHead.Left = root
	parent, target := inorder(dummyHead, root, key)
	if parent == nil {
		return root
	}
	if parent == dummyHead {
		parent.Left = deleteRoot(target)
	} else {
		if parent.Val > target.Val {
			parent.Left = deleteRoot(target)
		} else {
			parent.Right = deleteRoot(target)
		}
	}
	return dummyHead.Left
}

func deleteNodeRecursive(root *TreeNode, key int) *TreeNode {
	if root == nil {
		return nil
	}
	if key > root.Val {
		root.Right = deleteNodeRecursive(root.Right, key)
	} else if key < root.Val {
		root.Left = deleteNodeRecursive(root.Left, key)
	} else {
		root = deleteRoot(root)
	}
	return root
}

func main() {
	t := NewBinaryTree([]string{"5", "3", "6", "2", "4", "null", "7"})
	t.Dump()
	t = deleteNodeRecursive(t, 5)
	t.Dump()
}
