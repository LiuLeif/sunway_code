// 2018-10-22 10:16
package main

import (
	"fmt"
	"math"
)

func main() {
	fmt.Println(minFlipsMonoIncr("00011000"))
}

func minFlipsMonoIncr(S string) int {
	min := func(nums ...int) int {
		ret := math.MaxInt32
		for _, n := range nums {
			if ret > n {
				ret = n
			}
		}
		return ret
	}

	dpZero := make([]int, len(S))
	dpOne := make([]int, len(S))
	if S[0] == '0' {
		dpOne[0] = 1
	} else {
		dpZero[0] = 1
	}
	for i := 1; i < len(S); i++ {
		if S[i] == '0' {
			dpZero[i] = dpZero[i-1]
			dpOne[i] = dpOne[i-1] + 1
		} else {
			dpZero[i] = dpZero[i-1] + 1
			dpOne[i] = min(dpOne[i-1], dpZero[i-1])
		}
	}

	return min(dpOne[len(S)-1], dpZero[len(S)-1])
}
