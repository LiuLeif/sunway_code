// 2018-10-31 14:10
package main

import "fmt"

type Pair struct {
	first, second int
}

type Stack []Pair

func (this *Stack) Push(x Pair) {
	*this = append(*this, x)
}

func (this *Stack) Pop() Pair {
	ret := (*this)[len(*this)-1]
	*this = (*this)[:len(*this)-1]
	return ret
}

func (this Stack) Empty() bool {
	return len(this) == 0
}

func (this Stack) Peek() Pair {
	return this[len(this)-1]
}

type StockSpanner struct {
	stack Stack
}

func Constructor() StockSpanner {
	ret := StockSpanner{}
	ret.stack = Stack{}
	return ret
}

func (this *StockSpanner) Next(price int) int {
	if len(this.stack) == 0 {
		this.stack.Push(Pair{price, 1})
		return 1
	}
	ret := 1
	for len(this.stack) != 0 && this.stack[len(this.stack)-1].first <= price {
		tmp := this.stack.Pop()
		ret += tmp.second
	}
	this.stack.Push(Pair{price, ret})
	return ret
}

func main() {
	span := Constructor()
	fmt.Println(span.Next(28))
	fmt.Println(span.Next(14))
	fmt.Println(span.Next(28))
}
