// 2017-11-24 11:05
package main

import (
	"fmt"
	"math"
)

type MinStack struct {
	Nums []int
	Min  int
}

/** initialize your data structure here. */
func Constructor() MinStack {
	ret := MinStack{}
	ret.Nums = []int{}
	ret.Min = math.MaxInt32
	return ret
}

func (this *MinStack) Push(x int) {
	if x <= this.Min {
		this.Nums = append(this.Nums, this.Min)
		this.Min = x
	}
	this.Nums = append(this.Nums, x)
}

func (this *MinStack) Pop() {
	top := this.Nums[len(this.Nums)-1]
	this.Nums = this.Nums[:len(this.Nums)-1]
	if top == this.Min {
		this.Min = this.Nums[len(this.Nums)-1]
		this.Nums = this.Nums[:len(this.Nums)-1]
	}
}

func (this *MinStack) Top() int {
	return this.Nums[len(this.Nums)-1]
}

func (this *MinStack) GetMin() int {
	return this.Min
}

/**
 * Your MinStack object will be instantiated and called as such:
 * obj := Constructor();
 * obj.Push(x);
 * obj.Pop();
 * param_3 := obj.Top();
 * param_4 := obj.GetMin();
 */
func main() {
	stack := Constructor()
	stack.Push(0)
	stack.Push(1)
	stack.Push(0)
	fmt.Println(stack.GetMin())
	stack.Pop()
	fmt.Println(stack.GetMin())
}
