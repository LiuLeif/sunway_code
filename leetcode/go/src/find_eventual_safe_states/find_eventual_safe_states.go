// 2018-03-22 09:18
package main

import "fmt"

func eventualSafeNodes(graph [][]int) []int {
	ret := []int{}
	mem := map[int]bool{}

	var dfs func(int) bool
	dfs = func(start int) bool {
		if isSafe, exist := mem[start]; exist {
			return isSafe
		}
		neighbors := graph[start]
		if len(neighbors) == 0 {
			return true
		}
		mem[start] = false
		for _, n := range neighbors {
			if dfs(n) == false {
				return false
			}
		}
		mem[start] = true
		return true
	}
	for i := 0; i < len(graph); i++ {
		if dfs(i) {
			ret = append(ret, i)
		}
	}
	return ret
}
func main() {
	fmt.Println(eventualSafeNodes([][]int{
		{1, 2},
		{2, 3},
		{5},
		{0},
		{5},
		{},
		{},
	}))
}
