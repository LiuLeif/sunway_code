// 2018-10-24 15:01
package main

import "fmt"

func matrixScore(A [][]int) int {
	m, n := len(A), len(A[0])
	for i := 0; i < m; i++ {
		if A[i][0] == 0 {
			for j := 0; j < n; j++ {
				A[i][j] = (A[i][j] + 1) % 2
			}
		}
	}

	for j := 0; j < n; j++ {
		count := 0
		for i := 0; i < m; i++ {
			if A[i][j] == 1 {
				count++
			}
		}
		if count < m-count {
			for i := 0; i < m; i++ {
				A[i][j] = (A[i][j] + 1) % 2
			}
		}
	}

	ret := 0
	x := 1
	for j := n - 1; j >= 0; j-- {
		for i := 0; i < m; i++ {
			ret += x * A[i][j]
		}
		x *= 2
	}
	return ret
}

func main() {
	fmt.Println(matrixScore([][]int{
		{0, 1},
		{1, 1},
	}))
}
