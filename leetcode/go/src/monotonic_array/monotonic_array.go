// 2018-10-12 15:00
package main

import "fmt"

func isMonotonic(A []int) bool {
	if len(A) <= 2 {
		return true
	}

	decreasing := func() bool {
		for i := 1; i < len(A); i++ {
			if A[i] > A[i-1] {
				return false
			}
		}
		return true
	}

	increasing := func() bool {
		for i := 1; i < len(A); i++ {
			if A[i] < A[i-1] {
				return false
			}
		}
		return true
	}

	return increasing() || decreasing()
}

func main() {
	fmt.Println(isMonotonic([]int{2, 2, 2}))
}
