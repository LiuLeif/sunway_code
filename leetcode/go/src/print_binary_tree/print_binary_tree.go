// 2018-02-02 14:52
package main

import (
	"fmt"
	"math"
	"strconv"

	. "util/tree"
)

func max(nums ...int) int {
	ret := math.MinInt32
	for _, n := range nums {
		if ret < n {
			ret = n
		}
	}
	return ret
}

func getHeight(root *TreeNode) int {
	if root == nil {
		return 0
	}
	return max(getHeight(root.Left), getHeight(root.Right)) + 1
}

func getWidth(height int) int {
	return int(math.Pow(2.0, float64(height-1)))*2 - 1
}

func dfs(root *TreeNode, ret [][]string, level, lo, hi int) {
	if root == nil {
		return
	}
	mid := (lo + hi) / 2
	ret[level][mid] = strconv.Itoa(root.Val)
	dfs(root.Left, ret, level+1, lo, mid-1)
	dfs(root.Right, ret, level+1, mid+1, hi)
}

func printTree(root *TreeNode) [][]string {
	height := getHeight(root)
	ret := make([][]string, height)
	width := getWidth(height)
	for i := 0; i < height; i++ {
		ret[i] = make([]string, width)
	}
	dfs(root, ret, 0, 0, width-1)
	return ret
}
func main() {
	t := NewBinaryTree([]string{"1", "2", "5", "3", "null", "null", "null", "4"})
	fmt.Println(printTree(t))
}
