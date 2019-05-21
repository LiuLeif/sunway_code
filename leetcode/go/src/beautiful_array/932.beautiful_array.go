// 2018-10-29 10:40
package main

import "fmt"

func helper(x []int) []int {
	if len(x) <= 2 {
		return x
	}
	even := []int{}
	odd := []int{}
	for i := 0; i < len(x); i++ {
		if i&1 == 0 {
			even = append(even, x[i])
		} else {
			odd = append(odd, x[i])
		}
	}
	return append(helper(even), helper(odd)...)
}

func beautifulArray(N int) []int {
	// 1 2 3 4 5
	// 1 3 5 2 4
	//
	ret := make([]int, N)
	for i := 0; i < N; i++ {
		ret[i] = i + 1
	}
	ret = helper(ret)
	return ret
}

func main() {
	fmt.Println(beautifulArray(9))
}
