// 2017-11-17 17:12
package main

import "fmt"

func fourSumCount(A []int, B []int, C []int, D []int) int {
	sum := map[int]int{}
	for _, v := range A {
		for _, k := range B {
			sum[v+k]++
		}
	}
	ret := 0
	for _, v := range C {
		for _, k := range D {
			x := -(v + k)
			ret += sum[x]
		}
	}

	return ret
}
func main() {
	fmt.Println(fourSumCount([]int{1}, []int{-2}, []int{-1}, []int{0}))
}
