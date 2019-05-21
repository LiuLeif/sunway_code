// 2017-12-20 14:33
package main

import (
	"fmt"
	"math"
)

func isSquare(n int) bool {
	x := int(math.Sqrt(float64(n)))
	return x*x == n
}

func judgeSquareSum(n int) bool {
	i := 0
	for i*i <= n {
		if isSquare(n - i*i) {
			return true
		}
		i += 1
	}
	return false
}

func judgeSquareSumHash(n int) bool {
	m := make(map[int]bool)
	for i := 0; i <= int(math.Sqrt(float64(n))); i++ {
		m[i*i] = true
		if m[n-i*i] {
			return true
		}
	}
	return false
}

func main() {
	fmt.Println(judgeSquareSumHash(5))
}
