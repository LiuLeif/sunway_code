// 2018-11-19 13:42
package main

import "fmt"

func validMountainArray(A []int) bool {
	firstDecreasingRev := func() int {
		for i := len(A) - 2; i >= 0; i-- {
			if A[i] <= A[i+1] {
				return i + 1
			}
		}
		return -1
	}

	firstDecreasing := func() int {
		for i := 1; i < len(A); i++ {
			if A[i] <= A[i-1] {
				return i - 1
			}
		}
		return -2
	}

	return firstDecreasing() == firstDecreasingRev()
}

func main() {
	fmt.Println(validMountainArray([]int{0, 3, 2, 1}))
	fmt.Println(validMountainArray([]int{3, 5, 5}))
	fmt.Println(validMountainArray([]int{14, 82, 89, 84, 79, 70, 70, 68, 67, 66, 63, 60, 58, 54, 44, 43, 32, 28, 26, 25, 22, 15, 13, 12, 10, 8, 7, 5, 4, 3}))
}
