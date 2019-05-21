// 2018-10-11 19:15
package main

import "fmt"
import "sort"

func sortArrayByParity2(A []int) []int {
	sort.Slice(A, func(i, j int) bool {
		return A[i]%2 < A[j]%2
	})
	return A
}

func sortArrayByParity(A []int) []int {
	even := []int{}
	odd := []int{}
	for i := 0; i < len(A); i++ {
		if A[i]%2 == 0 {
			even = append(even, A[i])
		} else {
			odd = append(odd, A[i])
		}
	}
	return append(even, odd...)
}

func main() {
	fmt.Println(sortArrayByParity2([]int{3, 1, 2, 4}))
}
