// 2018-12-07 14:14
package main

import "fmt"

// 10,14
// 14,16

type SegTreeNode struct {
	lo, hi      int
	tracked     bool
	left, right *SegTreeNode
}

func SetTracking(root *SegTreeNode, left, right int, tracking bool) bool {
	if left <= root.lo && right >= root.hi {
		root.left = nil
		root.right = nil
		root.tracked = tracking
		return tracking
	}
	mid := (root.lo + root.hi) / 2
	if root.left == nil {
		root.left = &SegTreeNode{root.lo, mid, root.tracked, nil, nil}
	}
	if root.right == nil {
		root.right = &SegTreeNode{mid + 1, root.hi, root.tracked, nil, nil}
	}
	ltracking := root.left.tracked
	if left <= mid {
		ltracking = SetTracking(root.left, left, right, tracking)
	}
	rtracking := root.right.tracked
	if right > mid {
		rtracking = SetTracking(root.right, left, right, tracking)
	}
	root.tracked = ltracking && rtracking
	return root.tracked
}

func GetTracking(root *SegTreeNode, left, right int) bool {
	if root.left == nil && root.right == nil {
		return root.tracked
	}
	if left <= root.lo && right >= root.hi {
		return root.tracked
	}

	mid := (root.lo + root.hi) / 2
	rtracking, ltracking := true, true
	if left <= mid {
		ltracking = GetTracking(root.left, left, right)
	}
	if right > mid {
		rtracking = GetTracking(root.right, left, right)
	}
	return rtracking && ltracking
}

type RangeModule struct {
	root *SegTreeNode
}

func Constructor() RangeModule {
	return RangeModule{&SegTreeNode{1, 1000000000, false, nil, nil}}
}

func (this *RangeModule) AddRange(left int, right int) {
	SetTracking(this.root, left, right-1, true)
}

func (this *RangeModule) QueryRange(left int, right int) bool {
	return GetTracking(this.root, left, right-1)
}

func (this *RangeModule) RemoveRange(left int, right int) {
	SetTracking(this.root, left, right-1, false)
}

func main() {
	x := Constructor()
	x.AddRange(10, 20)
	x.RemoveRange(14, 16)
	fmt.Println(x.QueryRange(10, 14)) // true
	fmt.Println(x.QueryRange(13, 15)) // false
	fmt.Println(x.QueryRange(16, 17)) // true

}
