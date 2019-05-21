// 2018-12-12 14:23
package main

import (
	"fmt"
	"math"
)

func abs(x int) int {
	if x < 0 {
		return -x
	}
	return x
}

func max(nums ...int) int {
	ret := math.MinInt32
	for _, n := range nums {
		if ret < n {
			ret = n
		}
	}
	return ret
}

func tallestBillboard(rods []int) int {
	total := 0
	for _, v := range rods {
		total += v
	}
	total /= 2
	m := map[[2]int]int{}
	var dfs func(s1, s2, index int) int
	dfs = func(s1, s2, index int) int {
		if s1 > total || s2 > total {
			return 0
		}
		if index >= len(rods) {
			if s1 == s2 {
				return s1
			}
			return 0
		}
		if c, ok := m[[2]int{index, abs(s1 - s2)}]; ok {
			if c < 0 {
				return 0
			}
			return c + max(s1, s2)
		}
		ret := max(
			dfs(s1+rods[index], s2, index+1),
			dfs(s1, s2+rods[index], index+1),
			dfs(s1, s2, index+1),
		)
		// s1: 50, s2: 30, ret: 200
		m[[2]int{index, abs(s1 - s2)}] = ret - max(s1, s2)
		return ret
	}
	return dfs(0, 0, 0)
}

func main() {
	fmt.Println(tallestBillboard([]int{1, 2, 3, 4, 5, 6}))
	fmt.Println(tallestBillboard([]int{1, 2, 4, 8, 16, 32, 64, 128, 256, 512, 50, 50, 50, 150, 150, 150, 100, 100, 100, 123}))
	fmt.Println(tallestBillboard([]int{2, 4, 8, 16}))
}
