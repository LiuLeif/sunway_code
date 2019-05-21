// 2018-10-24 16:22
package main

import "fmt"
import "math/rand"
import "time"

type Solution struct {
	index int
	m, n  int
	hash  map[int]int
}

func Constructor(n_rows int, n_cols int) Solution {
	ret := Solution{}
	ret.index = n_rows*n_cols - 1
	ret.m = n_rows
	ret.n = n_cols
	ret.hash = map[int]int{}
	return ret
}

func (this *Solution) Flip() []int {
	getVal := func(index int) int {
		if x, ok := this.hash[index]; ok {
			return x
		}
		return index
	}
	setVal := func(i, j int) {
		this.hash[i] = getVal(j)
	}

	target := rand.Intn(this.index + 1)
	ret := getVal(target)
	setVal(target, this.index)
	this.index--
	return []int{ret / this.n, ret % this.n}
}

func (this *Solution) Reset() {
	this.index = this.m*this.n - 1
	this.hash = map[int]int{}
}

func main() {
	rand.Seed(time.Now().Unix())
	x := Constructor(2, 3)
	fmt.Println(x.Flip())
	fmt.Println(x.Flip())
	fmt.Println(x.Flip())
	fmt.Println(x.Flip())
	fmt.Println(x.Flip())
	fmt.Println(x.Flip())
}
