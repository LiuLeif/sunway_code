// 2018-10-18 12:01
package main

import (
	"fmt"
	"math"
)

func wordSubsets(A []string, B []string) []string {
	ret := []string{}
	max := func(nums ...int) int {
		ret := math.MinInt32
		for _, n := range nums {
			if ret < n {
				ret = n
			}
		}
		return ret
	}

	counter := map[rune]int{}
	for _, s := range B {
		tmpCounter := map[rune]int{}
		for _, c := range s {
			tmpCounter[c]++
		}
		for k, v := range tmpCounter {
			counter[k] = max(counter[k], v)
		}
	}
loop:
	for _, s := range A {
		tmpCounter := map[rune]int{}
		for _, c := range s {
			tmpCounter[c]++
		}
		for k, v := range counter {
			if tmpCounter[k] < v {
				continue loop
			}
		}
		ret = append(ret, s)
	}
	return ret
}
func main() {
	fmt.Println(wordSubsets([]string{"amazon", "apple", "facebook", "google", "leetcode"}, []string{"lo", "eo"}))
}
