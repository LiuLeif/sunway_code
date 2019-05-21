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

func findMaxLength(nums []int) int {
	m := make(map[int]int)
	m[0] = 0
	count := 0
	ret := 0
	for i, v := range nums {
		if v == 0 {
			count--
		} else {
			count++
		}
		_, ok := m[count]
		if !ok {
			m[count] = i + 1
		}

		ret = max(ret, i-m[count]+1)
	}

	return ret
}
func main() {
	fmt.Println(findMaxLength([]int{0, 1}))
}
