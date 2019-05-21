// 2018-11-13 13:39
package main

import (
	"fmt"
	"math"
)

func bestRotation(A []int) int {
	// 2 3 1 4 0 = 2
	// 0 1 2 3 4
	N := len(A)
	change := make([]int, N)
	for i := 0; i < N; i++ {
		change[i] = 1
	}

	for i := 0; i < N; i++ {
		index := (i - A[i] + 1 + N) % N
		change[index] -= 1
	}

	for i := 1; i < N; i++ {
		change[i] += change[i-1]
	}

	maxVal := math.MinInt32
	maxIndex := 0
	for i := 0; i < N; i++ {
		if change[i] > maxVal {
			maxVal = change[i]
			maxIndex = i
		}
	}
	return maxIndex
}

func main() {
	fmt.Println(bestRotation([]int{2, 3, 1, 4, 0}))
}
