// 2018-10-30 22:23
package main

import (
	"fmt"
)

func kthGrammar(N int, K int) int {
	if N == 1 {
		return 0
	}
	index := (K-1)/2 + 1
	offset := (K - 1) % 2

	x := kthGrammar(N-1, index)
	if x == offset {
		return 0
	}
	return 1
}

func main() {
	fmt.Println(kthGrammar(4, 6))
}
