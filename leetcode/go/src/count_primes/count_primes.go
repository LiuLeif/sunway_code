// 2017-11-26 21:10
package main

import (
	"fmt"
	"math"
)

func countPrimes(n int) int {
	if n <= 1 {
		return 0
	}
	notPrime := make([]bool, n)
	notPrime[1] = true
	for i := 2; i < int(math.Sqrt(float64(n)))+1; i++ {
		if notPrime[i] {
			continue
		}
		for j := 2; i*j < n; j++ {
			notPrime[i*j] = true
		}
	}

	ret := 0
	for i := 1; i < n; i++ {
		if !notPrime[i] {
			ret++
		}
	}

	return ret
}
func main() {
	fmt.Println(countPrimes(100))
}
