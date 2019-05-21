// 2017-12-29 14:47
package main

import (
	"fmt"
	"math"
)

func primes(n int) []int {
	if n < 3 {
		return []int{}
	}
	skipped := make([]bool, n)
	skipped[0] = true
	skipped[1] = true
	for i := 2; i <= int(math.Sqrt(float64(n))); i++ {
		if skipped[i] {
			continue
		}
		for j := 2; i*j < n; j++ {
			skipped[i*j] = true
		}
	}
	ret := []int{}
	for i := 0; i < n; i++ {
		if !skipped[i] {
			ret = append(ret, i)
		}
	}
	return ret
}

func main() {
	fmt.Println(primes(100))
}
