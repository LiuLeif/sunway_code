// 2018-01-02 17:35
package main

import "fmt"

type MyStack []int

func (this *MyStack) Push(x int) {
	*this = append(*this, x)
}

func (this *MyStack) Peek() int {
	return (*this)[len(*this)-1]
}

func (this *MyStack) Pop() int {
	ret := this.Peek()
	*this = (*this)[:len(*this)-1]
	return ret
}

func (thiz MyStack) Empty() bool {
	return len(thiz) == 0
}

type MyQueue struct {
	inputStack, outputStack MyStack
}

func Constructor() MyQueue {
	return MyQueue{MyStack{}, MyStack{}}
}

func (this *MyQueue) Push(x int) {
	this.inputStack.Push(x)
}

func (this *MyQueue) Pop() int {
	this.Peek()
	return this.outputStack.Pop()
}

func (this *MyQueue) Peek() int {
	if this.outputStack.Empty() {
		for !this.inputStack.Empty() {
			this.outputStack.Push(this.inputStack.Pop())
		}
	}
	return this.outputStack.Peek()
}

func (this *MyQueue) Empty() bool {
	return this.inputStack.Empty() && this.outputStack.Empty()
}

func main() {
	queue := Constructor()
	queue.Push(1)
	queue.Push(2)
	fmt.Println(queue.Pop())
	fmt.Println(queue.Peek())
	fmt.Println(queue.Empty())
}
