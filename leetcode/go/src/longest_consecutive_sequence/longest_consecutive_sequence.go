// 2017-11-14 18:54
package main

import (
	"fmt"
	"math"

	. "util/UnionFindSet"
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

func longestConsecutiveUFS(nums []int) int {
	ufs := NewUnionFindSet(nums)
	m := make(map[int]bool)
	for _, v := range nums {
		if m[v] {
			continue
		}
		m[v] = true
		if m[v-1] == true {
			ufs.Union(v-1, v)
		}
		if m[v+1] == true {
			ufs.Union(v+1, v)
		}
	}
	count := make(map[int]int)
	ret := 0
	m = make(map[int]bool)
	for _, v := range nums {
		if m[v] {
			continue
		}
		m[v] = true
		root := ufs.Find(v)
		count[root]++
		ret = max(ret, count[root])
	}

	return ret
}

func longestConsecutive(nums []int) int {
	m := make(map[int]int)
	ret := 0
	for _, v := range nums {
		if m[v] != 0 {
			continue
		}
		left, right := m[v-1], m[v+1]
		sum := left + right + 1
		m[v] = sum
		if ret < sum {
			ret = sum
		}
		m[v-left], m[v+right] = sum, sum
	}
	return ret
}
func main() {
	fmt.Println(longestConsecutiveUFS([]int{0, 2, 3, 1}))
	fmt.Println(longestConsecutiveUFS([]int{-2, -3, -3, 7, -3, 0, 5, 0, -8, -4, -1, 2}))
}
