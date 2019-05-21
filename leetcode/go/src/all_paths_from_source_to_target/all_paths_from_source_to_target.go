// 2018-03-30 13:44
package main

import "fmt"

func allPathsSourceTarget(graph [][]int) [][]int {
	ret := [][]int{}

	var dfs func([]int, int)
	dfs = func(rec []int, f int) {
		rec = append(rec, f)
		if f == len(graph)-1 {
			ret = append(ret, append([]int{}, rec...))
			return
		}
		for _, v := range graph[f] {
			dfs(rec, v)
		}
		rec = rec[:len(rec)-1]
	}

	dfs([]int{}, 0)
	return ret
}
func main() {
	fmt.Println(allPathsSourceTarget([][]int{
		{1, 2, 3},
		{3},
		{3},
		{},
	}))
}
