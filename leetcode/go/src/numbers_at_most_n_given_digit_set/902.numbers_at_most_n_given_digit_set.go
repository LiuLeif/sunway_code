// 2018-11-05 13:26
package main

import (
	"fmt"
	"strconv"
)

func atMostNGivenDigitSet(D []string, N int) int {
	// 1 3 5 7, 100
	// 1 3 5 7 11 13 15 17
	pow := func(a, b int) int {
		ret := 1
		for i := 0; i < b; i++ {
			ret *= a
		}
		return ret
	}

	numD := make([]int, len(D))
	for i := 0; i < len(D); i++ {
		t, _ := strconv.Atoi(D[i])
		numD[i] = t
	}

	var helper func(strN string) int
	helper = func(strN string) int {
		if len(strN) == 0 {
			return 1
		}
		head, _ := strconv.Atoi(string(strN[0]))
		ret := 0

		count := 0
		contains := false
		for _, v := range numD {
			if v < head {
				count++
			}
			if v == head {
				contains = true
			}
		}
		ret += count * (pow(len(numD), len(strN)-1))
		if contains {
			ret += helper(strN[1:])
		}
		return ret
	}

	strN := strconv.Itoa(N)
	ret := helper(strN)
	for i := 1; i < len((strN)); i++ {
		ret += pow(len(numD), i)
	}
	return ret
}

func main() {
	fmt.Println(atMostNGivenDigitSet([]string{"1", "3", "5", "7"}, 100))
}
