// 2017-11-14 18:54
package main

import (
	"fmt"
	"math"
	"sort"

	. "util/tree"
)

func walk(root *TreeNode, record *[]int) {
	if root == nil {
		return
	}
	*record = append(*record, root.Val)
	walk(root.Left, record)
	walk(root.Right, record)
	return
}

func getMinimumDifference(root *TreeNode) int {
	record := make([]int, 0)
	walk(root, &record)
	sort.Ints(record)
	ret := math.MaxInt32
	for i := 1; i < len(record); i++ {
		if ret >= record[i]-record[i-1] {
			ret = record[i] - record[i-1]
		}
	}
	return ret
}

func main() {
	t := NewBinaryTree([]string{"1", "null", "5", "3"})
	fmt.Println(getMinimumDifference(t))
}
