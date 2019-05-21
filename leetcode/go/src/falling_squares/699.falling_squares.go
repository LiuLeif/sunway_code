// 2018-11-20 10:46
package main

import (
	"fmt"
	"math"
	"sort"
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

type Node struct {
	lo, hi      int
	left, right *Node
	val         int
}

func buildTree(lo, hi int) *Node {
	node := Node{lo, hi, nil, nil, 0}
	if lo == hi {
		return &node
	}
	mid := (lo + hi) / 2
	node.left = buildTree(lo, mid)
	node.right = buildTree(mid+1, hi)
	return &node
}

func query(root *Node, lo, hi int) int {
	if root.lo == lo && root.hi == hi {
		return root.val
	}
	ret := 0
	mid := (root.lo + root.hi) / 2
	if hi <= mid {
		ret = query(root.left, lo, hi)
	} else if lo > mid {
		ret = query(root.right, lo, hi)
	} else {
		ret = max(query(root.left, lo, mid), query(root.right, mid+1, hi))
	}
	return ret
}

func insert(root *Node, lo, hi, val int) {
	if lo == root.lo && hi == root.hi && lo == hi {
		root.val = val
		return
	}
	mid := (root.lo + root.hi) / 2
	if hi <= mid {
		insert(root.left, lo, hi, val)
	} else if lo > mid {
		insert(root.right, lo, hi, val)
	} else {
		insert(root.left, lo, mid, val)
		insert(root.right, mid+1, hi, val)
	}
	root.val = max(root.val, root.left.val, root.right.val)
}

func fallingSquares(positions [][]int) []int {
	ret := []int{}
	// discreate
	m := map[int]bool{}
	for _, pos := range positions {
		m[pos[0]] = true
		m[pos[0]+pos[1]-1] = true
	}
	keys := []int{}
	for k, _ := range m {
		keys = append(keys, k)
	}
	sort.Ints(keys)
	N := len(keys)
	forwardMapping := map[int]int{}
	for i, k := range keys {
		forwardMapping[k] = i
	}
	// discreate done
	root := buildTree(0, N-1)
	for i := 0; i < len(positions); i++ {
		pos := positions[i]
		lo, hi := forwardMapping[pos[0]], forwardMapping[pos[0]+pos[1]-1]
		largest := query(root, lo, hi)
		insert(root, lo, hi, largest+pos[1])
		ret = append(ret, query(root, 0, N-1))
	}
	return ret
}

func main() {
	fmt.Println(fallingSquares([][]int{{1, 2}, {2, 3}, {6, 1}}))
	fmt.Println(fallingSquares([][]int{{9, 7}, {1, 9}, {3, 1}}))
	fmt.Println(fallingSquares([][]int{{1, 5}, {2, 2}, {7, 5}}))
}
