// 2018-11-19 19:02
package main

import "fmt"

func minDeletionSize(A []string) int {
	// cba
	// daf
	// ghi
	//
	R, C := len(A), len(A[0])
	ret := 0
loop:
	for i := 0; i < C; i++ {
		for j := 1; j < R; j++ {
			if A[j][i] < A[j-1][i] {
				ret++
				continue loop
			}
		}
	}
	return ret
}

func main() {
	fmt.Println(minDeletionSize([]string{"zyx", "wvu", "tsr"}))
}
