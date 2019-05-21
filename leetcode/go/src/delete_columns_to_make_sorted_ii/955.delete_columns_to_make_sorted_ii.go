// 2018-12-10 16:31
package main

import "fmt"

func minDeletionSize(A []string) int {
	M, N := len(A), len(A[0])
	ignored := make([]bool, M)
	ret := 0
loop:
	for j := 0; j < N; j++ {
		for i := 1; i < M; i++ {
			if ignored[i-1] {
				continue
			}
			if A[i][j] < A[i-1][j] {
				ret++
				continue loop
			}
		}
		for i := 1; i < M; i++ {
			if ignored[i-1] {
				continue
			}
			// a b b c
			if A[i][j] > A[i-1][j] {
				ignored[i-1] = true
			} else {
				ignored[i-1] = false
			}
		}
	}
	return ret
}

func main() {
	fmt.Println(minDeletionSize([]string{"ca", "de", "xy"}))
	fmt.Println(minDeletionSize([]string{"xc", "yb", "za"}))
	fmt.Println(minDeletionSize([]string{"zyx", "wvu", "tsr"}))
	fmt.Println(minDeletionSize([]string{"abx", "agz", "bgc", "bfc"}))
}
