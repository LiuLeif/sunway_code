// 2017-11-15 10:19
package main

import (
	"fmt"
	"sort"
)

func dfs(ret *[][]int, record []int, candidates []int, target int, start int) {
	if target == 0 {
		*ret = append(*ret, append([]int{}, record...))
		return
	}
	for i := start; i < len(candidates); i++ {
		if i > start && candidates[i] == candidates[i-1] {
			continue
		}

		if target >= candidates[i] {
			record = append(record, candidates[i])
			dfs(ret, record, candidates, target-candidates[i], i+1)
			record = record[:len(record)-1]
		}
	}
}

func combinationSum2(candidates []int, target int) [][]int {
	sort.Ints(candidates)
	ret := [][]int{}
	dfs(&ret, []int{}, candidates, target, 0)
	return ret
}
func main() {
	fmt.Println(combinationSum2([]int{10, 1, 2, 7, 6, 1, 5}, 8))
}
