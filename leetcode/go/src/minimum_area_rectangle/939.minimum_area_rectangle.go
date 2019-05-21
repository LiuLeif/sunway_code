// 2018-11-12 11:31
package main

import (
	"fmt"
	"math"
)

func minAreaRect(points [][]int) int {
	min := func(nums ...int) int {
		ret := math.MaxInt32
		for _, n := range nums {
			if ret > n {
				ret = n
			}
		}
		return ret
	}
	abs := func(x int) int {
		if x >= 0 {
			return x
		}
		return -x
	}

	area := func(p1, p2 []int) int {
		return abs(p1[0]-p2[0]) * abs(p1[1]-p2[1])
	}

	m := map[[2]int]bool{}
	for _, p := range points {
		m[[2]int{p[0], p[1]}] = true
	}

	ret := math.MaxInt32
	for _, p1 := range points {
		for _, p2 := range points {
			if p1[0] == p2[0] || p1[1] == p2[1] {
				continue
			}
			p3 := [2]int{p1[0], p2[1]}
			p4 := [2]int{p2[0], p1[1]}
			if m[p3] && m[p4] {
				ret = min(ret, area(p1, p2))
			}
		}
	}
	if ret == math.MaxInt32 {
		return 0
	}
	return ret
}

func main() {
	fmt.Println(minAreaRect([][]int{{0, 1}, {1, 3}, {3, 3}, {4, 4}, {1, 4}, {2, 3}, {1, 0}, {3, 4}}))
}
