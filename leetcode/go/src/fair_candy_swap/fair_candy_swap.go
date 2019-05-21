// 2018-10-15 15:39
package main

import "fmt"

func fairCandySwap(A []int, B []int) []int {
	counter := map[int]int{}
	sum_b := 0
	sum_a := 0

	for i := 0; i < len(B); i++ {
		counter[B[i]] = i
		sum_b += B[i]
	}

	for i := 0; i < len(A); i++ {
		sum_a += A[i]
	}

	if sum_a > sum_b {
		for i := 0; i < len(A); i++ {
			if index, ok := counter[A[i]-(sum_a-sum_b)/2]; ok {
				return []int{A[i], B[index]}
			}
		}
	} else {
		for i := 0; i < len(A); i++ {
			if index, ok := counter[A[i]+(sum_b-sum_a)/2]; ok {
				return []int{A[i], B[index]}
			}
		}
	}
	return []int{}
}
func main() {
	fmt.Println(fairCandySwap([]int{2}, []int{1, 3}))
}
