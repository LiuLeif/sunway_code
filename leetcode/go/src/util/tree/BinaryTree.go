// 2017-12-19 17:32
package BinaryTree

import (
	"fmt"
	"strconv"
)

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func prettyPrintTree(head *TreeNode, prefix string, isLeft bool) string {
	if head == nil {
		return ""
	}
	var ret string = ""
	if head.Right != nil {
		newPrefix := prefix
		if isLeft {
			newPrefix += "| "
		} else {
			newPrefix += "  "
		}
		ret += prettyPrintTree(head.Right, newPrefix, false)
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
		ret += prettyPrintTree(head.Left, newPrefix, true)
	}
	return ret
}

func (this *TreeNode) String() string {
	if this == nil {
		return "nil"
	}
	return strconv.Itoa(this.Val)
}

func (this *TreeNode) Dump() {
	fmt.Println(prettyPrintTree(this, "", true))
}

func NewBinaryTree(nodes []string) *TreeNode {
	queue := make([]*TreeNode, 0)
	if len(nodes) == 0 {
		return nil
	}
	x, _ := strconv.Atoi(nodes[0])
	root := &TreeNode{x, nil, nil}
	i := 1
	queue = append(queue, root)
	for i < len(nodes) {
		p := queue[0]
		queue = queue[1:]
		if nodes[i] != "null" {
			x, _ := strconv.Atoi(nodes[i])
			tmp := &TreeNode{x, nil, nil}
			p.Left = tmp
			queue = append(queue, tmp)
		}
		i += 1
		if i < len(nodes) && nodes[i] != "null" {
			x, _ := strconv.Atoi(nodes[i])
			tmp := &TreeNode{x, nil, nil}
			p.Right = tmp
			queue = append(queue, tmp)
		}
		i += 1
	}
	return root
}
