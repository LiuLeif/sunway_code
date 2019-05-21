// 2017-11-16 14:53
package main

import (
	"fmt"
	"math"
)

func min(nums ...int) int {
	ret := math.MaxInt32
	for _, n := range nums {
		if ret > n {
			ret = n
		}
	}
	return ret
}

func countBinarySubstrings(s string) int {
	groups := []int{}
	current := 1
	for i := 1; i < len(s); i++ {
		if s[i] == s[i-1] {
			current++
		} else {
			groups = append(groups, current)
			current = 1
		}
	}

	groups = append(groups, current)
	ret := 0
	for i := 1; i < len(groups); i++ {
		ret += min(groups[i], groups[i-1])
	}

	return ret
}
func main() {
	fmt.Println(countBinarySubstrings("00110011"))
	fmt.Println(countBinarySubstrings("001011"))
}
