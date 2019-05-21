// 2018-10-31 09:30
package main

import "fmt"

type RLEIterator struct {
	index []int
	data  []int
	total int
}

func Constructor(A []int) RLEIterator {
	ret := RLEIterator{}
	ret.index = make([]int, len(A)/2)
	ret.data = make([]int, len(A)/2)

	index := 0
	for i := 0; i < len(A); i += 2 {
		if A[i] == 0 {
			continue
		}
		if i/2 == 0 {
			ret.index[index] = A[i]
		} else {
			ret.index[index] = ret.index[index-1] + A[i]
		}

		ret.data[index] = A[i+1]
		index++
	}

	ret.index = ret.index[:index]
	ret.data = ret.data[:index]

	return ret
}

func (this *RLEIterator) Next(n int) int {
	this.total += n
	ceil := func(target int) int {
		lo, hi := 0, len(this.index)-1
		for lo <= hi {
			mid := (lo + hi) / 2
			if this.index[mid] >= target {
				hi = mid - 1
			} else {
				lo = mid + 1
			}
		}
		if lo >= len(this.index) {
			return -1
		}
		return this.data[lo]
	}
	return ceil(this.total)
}

func main() {
	iter := Constructor([]int{3, 8, 0, 9, 2, 5})
	fmt.Println(iter)
	fmt.Println(iter.Next(2))
	fmt.Println(iter.Next(1))
	fmt.Println(iter.Next(1))
	fmt.Println(iter.Next(2))
}
