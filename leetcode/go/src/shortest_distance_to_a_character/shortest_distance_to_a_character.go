// 2018-10-16 10:22
package main

import "fmt"

func shortestToChar(S string, C byte) []int {
	min := func(a, b int) int {
		if a > b {
			return b
		}
		return a
	}
	ret := []int{}
	pending := 0
	first := true
	for i := 0; i < len(S); i++ {
		if S[i] == C {
			for i := 0; i < pending; i++ {
				if first {
					ret = append(ret, pending-i)
				} else {
					ret = append(ret, min(i+1, pending-i))
				}
			}
			first = false
			pending = 0
			ret = append(ret, 0)
		} else {
			pending++
		}
	}
	for i := 0; i < pending; i++ {
		ret = append(ret, i+1)
	}
	return ret
}
func main() {
	fmt.Println(shortestToChar("aaab", 'b'))
}
