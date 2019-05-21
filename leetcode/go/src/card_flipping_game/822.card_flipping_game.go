// 2018-11-08 14:04
package main

import (
	"fmt"
	"sort"
)

func flipgame(fronts []int, backs []int) int {
	counter := map[int]bool{}
	all := []int{}
	for i := 0; i < len(fronts); i++ {
		if fronts[i] == backs[i] {
			counter[fronts[i]] = true
		}
		all = append(all, fronts[i])
		all = append(all, backs[i])
	}
	sort.Ints(all)
	for i := 0; i < len(all); i++ {
		if !counter[all[i]] {
			return all[i]
		}
	}
	return 0
}

func main() {
	fmt.Println(flipgame([]int{1, 2, 4, 4, 7}, []int{1, 3, 4, 1, 3}))
}
