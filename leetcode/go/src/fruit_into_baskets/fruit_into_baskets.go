// 2018-10-17 12:02
package main

import (
	"fmt"
)

func totalFruit(tree []int) int {
	type1, type2 := -1, -1
	start1, start2 := 0, 0
	end1, end2 := 0, 0
	ret := 0

	min := func(a, b int) int {
		if a > b {
			return b
		}
		return a
	}
	max := func(a, b int) int {
		if a > b {
			return a
		}
		return b
	}

	for i := 0; i < len(tree); i++ {
		if type1 == -1 {
			type1 = tree[i]
			start1 = i
			end1 = i
			continue
		}
		if tree[i] == type1 {
			end1 = i
			continue
		}

		if type2 == -1 {
			type2 = tree[i]
			start2 = i
			end2 = i
			continue
		}
		if tree[i] == type2 {
			end2 = i
			continue
		}

		ret = max(ret, i-min(start1, start2))
		// fmt.Println(type1, type2, start1, start2, end1, end2, i, ret)
		if end2 > end1 {
			type1 = tree[i]
			start1 = i
			start2 = end1 + 1
			end1 = i
		} else {
			type2 = tree[i]
			start2 = i
			start1 = end2 + 1
			end2 = i
		}
	}

	ret = max(ret, len(tree)-min(start1, start2))
	return ret
}

func main() {
	fmt.Println(totalFruit([]int{1, 0, 3, 4, 3, 4}))
}
