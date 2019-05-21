// 2018-02-27 15:47
package main

import (
	"fmt"
	"sort"
)

func isIdealPermutation(A []int) bool {
	for i := 0; i < len(A)-1; {
		if A[i] > A[i+1] {
			A[i], A[i+1] = A[i+1], A[i]
			i += 2
		} else {
			i++
		}
	}
	return sort.IntsAreSorted(A)
}
func main() {
	fmt.Println(isIdealPermutation([]int{1, 0, 2}))
	fmt.Println(isIdealPermutation([]int{1, 2, 0}))
	fmt.Println(isIdealPermutation([]int{0, 2, 1}))
}
