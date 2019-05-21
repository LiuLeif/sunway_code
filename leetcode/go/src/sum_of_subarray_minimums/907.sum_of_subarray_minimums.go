// 2018-10-22 10:52
package main

import (
	"fmt"
	"math"
)

type Stack []int

func (this *Stack) Push(x int) {
	*this = append(*this, x)
}

func (this *Stack) Pop() int {
	ret := (*this)[len(*this)-1]
	*this = (*this)[:len(*this)-1]
	return ret
}

func (this Stack) Empty() bool {
	return len(this) == 0
}

func (this Stack) Top() int {
	return this[len(this)-1]
}

func NewStack() Stack {
	return Stack{}
}

func sumSubarrayMins(A []int) int {
	leftSmaller := make([]int, len(A))
	rightSmaller := make([]int, len(A))
	MOD := int(math.Pow(10, 9)) + 7
	stack := NewStack()
	for i := 0; i < len(A); i++ {
		for (!stack.Empty()) && A[stack.Top()] >= A[i] {
			stack.Pop()
		}
		if stack.Empty() {
			leftSmaller[i] = -1
		} else {
			leftSmaller[i] = stack.Top()
		}
		stack.Push(i)
	}

	stack = NewStack()
	for i := len(A) - 1; i >= 0; i-- {
		for (!stack.Empty()) && A[stack.Top()] > A[i] {
			stack.Pop()
		}
		if stack.Empty() {
			rightSmaller[i] = len(A)
		} else {
			rightSmaller[i] = stack.Top()
		}
		stack.Push(i)
	}

	ret := 0

	for i := 0; i < len(A); i++ {
		left := i - leftSmaller[i]
		right := rightSmaller[i] - i
		ret += left * right * A[i]
		ret %= MOD
	}
	return ret
}

func main() {
	// 3 1 2 4 3 1
	fmt.Println(sumSubarrayMins([]int{71, 55, 82, 55}))
}
