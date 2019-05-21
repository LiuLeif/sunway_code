// 2018-02-26 10:22
package main

import "fmt"

func dfs(graph [][]int, start int, color map[int]bool, required bool) bool {
	if _, ok := color[start]; ok {
		return true
	}
	color[start] = required
	for _, neighbor := range graph[start] {
		if _, ok := color[neighbor]; !ok {
			if !dfs(graph, neighbor, color, !color[start]) {
				return false
			}
		} else {
			if color[neighbor] == color[start] {
				return false
			}
		}
	}
	return true
}

func isBipartite(graph [][]int) bool {
	color := map[int]bool{}
	for i := 0; i < len(graph); i++ {
		if !dfs(graph, i, color, true) {
			return false
		}
	}
	return true
}

func main() {
	graph := [][]int{
		{1},
		{0, 3},
		{3},
		{1, 2},
	}
	fmt.Println(isBipartite(graph))
}
