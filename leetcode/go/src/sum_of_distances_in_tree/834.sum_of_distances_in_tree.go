// 2018-11-20 17:57
package main

import (
	"fmt"
)

func sumOfDistancesInTree(N int, edges [][]int) []int {
	ret := make([]int, N)
	count := make([]int, N)

	conn := map[int][]int{}
	for _, e := range edges {
		if conn[e[0]] == nil {
			conn[e[0]] = []int{}
		}
		conn[e[0]] = append(conn[e[0]], e[1])
		if conn[e[1]] == nil {
			conn[e[1]] = []int{}
		}
		conn[e[1]] = append(conn[e[1]], e[0])
	}

	visited := map[int]bool{}
	visited2 := map[int]bool{}

	var dfs func(root int)
	dfs = func(root int) {
		visited[root] = true
		for _, n := range conn[root] {
			if visited[n] {
				continue
			}
			dfs(n)
			count[root] += count[n]
			ret[root] += ret[n] + count[n]
		}
		count[root] += 1
	}
	var dfs2 func(root int)
	dfs2 = func(root int) {
		visited2[root] = true
		for _, n := range conn[root] {
			if visited2[n] {
				continue
			}
			ret[n] = ret[root] - count[n] + N - count[n]
			dfs2(n)
		}
	}

	dfs(0)
	dfs2(0)
	return ret
}

func main() {
	fmt.Println(sumOfDistancesInTree(6, [][]int{
		{0, 1}, {0, 2}, {2, 3}, {2, 4}, {2, 5},
	}))
	fmt.Println(sumOfDistancesInTree(3, [][]int{
		{2, 1}, {0, 2},
	}))
}
