// 2018-10-16 13:39
package main

import "fmt"

func sortArrayByParityII(A []int) []int {
	p := 0
	for f := 1; f < len(A); f++ {
		if A[f]&1 == 0 {
			A[f], A[p+1] = A[p+1], A[f]
			A[p], A[p+1] = A[p+1], A[p]
			p++
		}
	}
	i := 1
	j := len(A) / 2
	for ; i < len(A)-1; i += 2 {
		A[i], A[j] = A[j], A[i]
		j++
	}
	return A
}
func main() {
	fmt.Println(sortArrayByParityII([]int{4, 2, 5, 7}))
}
