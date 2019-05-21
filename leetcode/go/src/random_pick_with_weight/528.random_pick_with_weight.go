// 2018-10-25 11:31
package main

import "fmt"

import "math/rand"

type Solution struct {
	leftSum []int
}

func Constructor(w []int) Solution {
	ret := Solution{}
	ret.leftSum = make([]int, len(w))
	ret.leftSum[0] = w[0]
	for i := 1; i < len(w); i++ {
		ret.leftSum[i] = w[i] + ret.leftSum[i-1]
	}
	return ret
}

func (this *Solution) PickIndex() int {
	// 1 3 2 1
	// 1 4 6 7
	n := rand.Intn(this.leftSum[len(this.leftSum)-1])
	firstLarger := func(n int) int {
		lo, hi := 0, len(this.leftSum)-1
		for lo <= hi {
			mid := (lo + hi) / 2
			if this.leftSum[mid] <= n {
				lo = mid + 1
			} else {
				hi = mid - 1
			}
		}
		return lo
	}
	return firstLarger(n)
}

func main() {
	obj := Constructor([]int{3, 14, 1, 7})
	a, b := 0, 0
	for i := 0; i < 10000; i++ {
		x := obj.PickIndex()
		if x == 0 {
			a++
		} else if x == 2 {
			b++
		}
	}
	fmt.Println(a, b, float32(a)/float32(b))
}
