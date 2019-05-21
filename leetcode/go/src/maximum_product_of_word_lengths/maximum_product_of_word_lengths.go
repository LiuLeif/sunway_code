// 2017-12-19 15:43
package main

import (
	"fmt"
	"math"
	"sort"
)

func max(nums ...int) int {
	ret := math.MinInt32
	for _, n := range nums {
		if ret < n {
			ret = n
		}
	}
	return ret
}

func maxProduct(words []string) int {
	sort.Slice(words, func(i, j int) bool {
		return len(words[i]) > len(words[j])
	})
	m := make(map[string]int)
	for _, w := range words {
		t := 0
		for i := 0; i < len(w); i++ {
			t |= 1 << (w[i] - 'a')
		}
		m[w] = t
	}
	ret := 0
	for i := 0; i < len(words); i++ {
		if len(words[i])*len(words[i]) <= ret {
			break
		}
		for j := i + 1; j < len(words); j++ {
			if m[words[i]]&m[words[j]] == 0 {
				ret = max(ret, len(words[i])*len(words[j]))
				break
			}
		}

	}

	return ret
}
func main() {
	// abcw, xtfn
	fmt.Println(maxProduct([]string{"abcw", "xyz"}))
}
