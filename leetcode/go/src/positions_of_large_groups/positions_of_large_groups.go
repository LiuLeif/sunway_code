// 2018-10-15 15:18
package main

import "fmt"

func largeGroupPositions(S string) [][]int {
	ret := [][]int{}
	prev := 0
	for i := 1; i < len(S); i++ {
		if S[i] != S[prev] {
			if i-prev >= 3 {
				ret = append(ret, []int{prev, i - 1})
			}
			prev = i
		}
	}
	if len(S)-prev >= 3 {
		ret = append(ret, []int{prev, len(S) - 1})
	}
	return ret
}
func main() {
	fmt.Println(largeGroupPositions("abbxxxxzzy"))
}
