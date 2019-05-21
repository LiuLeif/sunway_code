// 2018-10-23 16:05
package main

import "fmt"

type MyCircularDeque struct {
	data []int
	k    int
}

/** Initialize your data structure here. Set the size of the deque to be k. */
func Constructor(k int) MyCircularDeque {
	return MyCircularDeque{[]int{}, k}
}

/** Adds an item at the front of Deque. Return true if the operation is successful. */
func (this *MyCircularDeque) InsertFront(value int) bool {
	if len(this.data) == this.k {
		return false
	}
	this.data = append([]int{value}, this.data...)
	return true
}

/** Adds an item at the rear of Deque. Return true if the operation is successful. */
func (this *MyCircularDeque) InsertLast(value int) bool {
	if len(this.data) == this.k {
		return false
	}
	this.data = append(this.data, value)
	return true
}

/** Deletes an item from the front of Deque. Return true if the operation is successful. */
func (this *MyCircularDeque) DeleteFront() bool {
	if len(this.data) == 0 {
		return false
	}
	this.data = this.data[1:]
	return true
}

/** Deletes an item from the rear of Deque. Return true if the operation is successful. */
func (this *MyCircularDeque) DeleteLast() bool {
	if len(this.data) == 0 {
		return false
	}
	this.data = this.data[:len(this.data)-1]
	return true
}

/** Get the front item from the deque. */
func (this *MyCircularDeque) GetFront() int {
	if len(this.data) == 0 {
		return -1
	}
	return this.data[0]
}

/** Get the last item from the deque. */
func (this *MyCircularDeque) GetRear() int {
	if len(this.data) == 0 {
		return -1
	}
	return this.data[len(this.data)-1]
}

/** Checks whether the circular deque is empty or not. */
func (this *MyCircularDeque) IsEmpty() bool {
	return len(this.data) == 0
}

/** Checks whether the circular deque is full or not. */
func (this *MyCircularDeque) IsFull() bool {
	return len(this.data) == this.k
}

func main() {
	deque := Constructor(10)
	deque.InsertFront(1)
	deque.InsertLast(2)
	deque.DeleteFront()
	deque.DeleteLast()
	fmt.Println(deque.GetFront())
	fmt.Println(deque.GetRear())
	fmt.Println(deque.IsEmpty())
	fmt.Println(deque.IsFull())
}
