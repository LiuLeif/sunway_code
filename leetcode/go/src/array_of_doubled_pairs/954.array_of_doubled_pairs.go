// 2018-12-10 15:36
package main

import (
	"fmt"
	"sort"
)

func canReorderDoubled(A []int) bool {
	// 1 2 4 4 8 16
	x, y := []int{}, []int{}
	for _, v := range A {
		if v >= 0 {
			x = append(x, v)
		} else {
			y = append(y, -v)
		}
	}
	sort.Ints(x)
	sort.Ints(y)

	canReorder := func(x []int) bool {
		// 1 2 4 4 8 16
		count := map[int]int{}
		for i := 0; i < len(x); i++ {
			count[x[i]]++
		}
		for i := 0; i < len(x); i++ {
			if count[x[i]] == 0 {
				continue
			}
			count[x[i]]--
			if count[x[i]*2] == 0 {
				return false
			}
			count[x[i]*2]--
		}
		return true
	}

	return canReorder(x) && canReorder(y)
}

func main() {
	fmt.Println(canReorderDoubled([]int{4, -2, 2, -4}))
	fmt.Println(canReorderDoubled([]int{1, 2, 4, 16, 8, 4, 8, 8}))
}
