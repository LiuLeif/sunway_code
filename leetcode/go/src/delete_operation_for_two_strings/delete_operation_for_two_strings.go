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

func minDistance(w1 string, w2 string) int {
	m := make([][]int, len(w1)+1)
	for i := 0; i < len(w1)+1; i++ {
		m[i] = make([]int, len(w2)+1)
	}

	mm := 0
	for i := 1; i < len(w1)+1; i++ {
		for j := 1; j < len(w2)+1; j++ {
			if w1[i-1] == w2[j-1] {
				m[i][j] = m[i-1][j-1] + 1
			} else {
				m[i][j] = max(m[i][j-1], m[i-1][j])
			}
			mm = max(mm, m[i][j])
		}
	}
	return len(w1) + len(w2) - mm*2
}
func main() {
	fmt.Println(minDistance("sea", ""))
}
