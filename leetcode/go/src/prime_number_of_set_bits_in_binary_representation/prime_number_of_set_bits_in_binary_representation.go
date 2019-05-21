// 2018-02-24 14:02
package main

import "fmt"

func countPrimeSetBits(L int, R int) int {
	ret := 0
	primes := map[int]bool{
		2: true, 3: true, 5: true, 7: true, 11: true, 13: true, 17: true, 19: true,
	}
	for i := L; i <= R; i++ {
		bits := 0
		for n := i; n > 0; n &= n - 1 {
			bits++
		}
		if primes[bits] {
			ret++
		}
	}
	return ret
}
func main() {
	fmt.Println(countPrimeSetBits(10, 15))
}
