// 2018-11-01 10:30
package main

import "fmt"

type Queue []int

func (this *Queue) Push(x int) {
	*this = append(*this, x)
}

func (this *Queue) Pop() int {
	ret := (*this)[0]
	*this = (*this)[1:]
	return ret
}

func (this Queue) Empty() bool {
	return len(this) == 0
}

func (this Queue) Peek() int {
	return this[0]
}

type MyStack struct {
	q1, q2 Queue
	top    int
}

//  q1: 3 2 1
//  q2:
//  3 2 1
/** Initialize your data structure here. */
func Constructor() MyStack {
	ret := MyStack{q1: Queue{}, q2: Queue{}}
	return ret
}

/** Push element x onto stack. */
func (this *MyStack) Push(x int) {
	this.q1.Push(x)
	this.top = x
}

/** Pop Removes the element on top of the stack and returns that element. */
func (this *MyStack) Pop() int {
	for !this.q1.Empty() {
		x := this.q1.Pop()
		if this.q1.Empty() {
			this.q1, this.q2 = this.q2, this.q1
			return x
		}
		this.top = x
		this.q2.Push(x)
	}
	return -1
}

/** Get the top element. */
func (this *MyStack) Top() int {
	return this.top
}

/** Returns whether the stack is empty. */
func (this *MyStack) Empty() bool {
	return this.q1.Empty()
}

func main() {
	stack := Constructor()
	stack.Push(1)
	stack.Push(2)
	fmt.Println(stack.Pop())
	fmt.Println(stack.Top())
}
