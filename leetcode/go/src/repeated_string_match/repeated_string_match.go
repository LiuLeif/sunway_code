// 2018-01-09 16:58
package main

import (
	"fmt"
	"math"
)

func repeatedStringMatch(A string, B string) int {
	for i := 0; i < len(A); i++ {
		j := 0
		for ; j < len(B); j++ {
			ai := i + j
			if ai >= len(A) {
				ai %= len(A)
			}
			if A[ai] != B[j] {
				break
			}
		}
		if j == len(B) {
			return int(math.Ceil(float64(i+j) / float64(len(A))))
		}
	}
	return -1
}
func main() {
	fmt.Println(repeatedStringMatch("abcd", "cdabcdab"))
}
