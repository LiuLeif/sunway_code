// 2017-11-14 18:54
package main

import "fmt"

func dfs(ret *[][]int, record []int, candidates []int, target int, start int) {
	if target == 0 {
		*ret = append(*ret, append([]int(nil), record...))
		return
	}

	for i := start; i < len(candidates); i++ {
		if target >= candidates[i] {
			record = append(record, candidates[i])
			dfs(ret, record, candidates, target-candidates[i], i)
			record = record[:len(record)-1]
		}
	}

}

func combinationSum(candidates []int, target int) [][]int {
	ret := [][]int{}
	dfs(&ret, []int{}, candidates, target, 0)
	return ret
}
func main() {
	fmt.Println(combinationSum([]int{2, 3, 6, 7}, 1))
}
