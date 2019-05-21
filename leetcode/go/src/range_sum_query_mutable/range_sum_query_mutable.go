// 2018-01-04 17:04
package main

import "fmt"

type NumArray struct {
	nums []int
	root *Node
}

type Node struct {
	f, t        int
	val         int
	left, right *Node
}

func (this *NumArray) buildTree(f, t int) *Node {
	ret := Node{f, t, 0, nil, nil}
	if f == t {
		ret.val = this.nums[f]
		return &ret
	}
	mid := (f + t) / 2
	ret.left, ret.right = this.buildTree(f, mid), this.buildTree(mid+1, t)
	ret.val = ret.left.val + ret.right.val
	return &ret
}

func (this *NumArray) update(index, delta int, root *Node) {
	root.val += delta
	if root.f == root.t {
		return
	}
	mid := (root.f + root.t) / 2
	if index <= mid {
		this.update(index, delta, root.left)
	} else {
		this.update(index, delta, root.right)
	}
}

func Constructor(nums []int) NumArray {
	ret := NumArray{nums, nil}
	ret.root = ret.buildTree(0, len(nums)-1)
	return ret
}

func (this *NumArray) Update(i int, val int) {
	delta := val - this.nums[i]
	this.nums[i] = val
	if delta == 0 {
		return
	}
	this.update(i, delta, this.root)
}

func (this *NumArray) sumRange(i, j int, root *Node) int {
	if root.f == i && root.t == j {
		return root.val
	}
	mid := (root.f + root.t) / 2
	if j <= mid {
		return this.sumRange(i, j, root.left)
	}
	if i > mid {
		return this.sumRange(i, j, root.right)
	}
	return this.sumRange(i, mid, root.left) + this.sumRange(mid+1, j, root.right)
}

func (this *NumArray) SumRange(i int, j int) int {
	return this.sumRange(i, j, this.root)
}

func main() {
	array := Constructor([]int{7, 2, 7, 2, 0})
	array.Update(4, 6)
	array.Update(0, 2)
	array.Update(0, 9)
	array.Update(3, 8)
	fmt.Println(array.SumRange(0, 4))
}
