// 2017-11-14 18:54
package main

import (
	"fmt"
	"math"
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

func count(m []int) int {
	max := 0
	all := 0
	for _, v := range m {
		all += v
		if v >= max {
			max = v
		}
	}
	return all - max
}

func characterReplacement(s string, k int) int {
	m := make([]int, 26)
	right := 0

	ret := 0
	for left := 0; left < len(s); left++ {
		for count(m) <= k && right < len(s) {
			m[s[right]-'A'] += 1
			right++
		}

		if count(m) > k {
			ret = max(ret, right-left-1)
		} else {
			ret = max(ret, right-left)
		}
		m[s[left]-'A'] -= 1
	}
	return ret
}

func main() {
	fmt.Println(characterReplacement("ABAA", 0))
}
