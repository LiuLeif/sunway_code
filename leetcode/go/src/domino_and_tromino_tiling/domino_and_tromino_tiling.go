// 2018-03-30 14:03
package main

import (
	"fmt"
)

func numTilings(N int) int {
	// XYZ XXZ XYY XXY XYY
	// XYZ YYZ XZZ XYY XXY

	// 1. A  2. A A  3. A A
	//    A     A A     A

	memA := map[int]int{}
	memB := map[int]int{}

	var numTilingsTypeA func(int) int
	var numTilingsTypeB func(int) int

	numTilingsTypeA = func(N int) int {
		if N == 0 {
			return 1
		}
		if N <= 2 {
			return N
		}
		if cache, ok := memA[N]; ok {
			return cache
		}
		ret := numTilingsTypeA(N-1) + numTilingsTypeA(N-2) + numTilingsTypeB(N-1)*2
		ret %= (1e9 + 7)
		memA[N] = ret
		return ret
	}

	numTilingsTypeB = func(N int) int {
		if N == 0 {
			return 1
		}
		if N == 1 {
			return 0
		}
		if N == 2 {
			return 1
		}
		if cache, ok := memB[N]; ok {
			return cache
		}
		ret := numTilingsTypeA(N-2) + numTilingsTypeB(N-1)
		ret %= (1e9 + 7)
		memB[N] = ret
		return ret
	}
	return numTilingsTypeA(N)
}
func main() {
	fmt.Println(numTilings(3))
}
