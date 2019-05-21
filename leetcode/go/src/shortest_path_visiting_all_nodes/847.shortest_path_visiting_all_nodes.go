// 2018-11-05 17:06
package main

import "fmt"

type Tuple struct {
	index  int
	length int
	mask   int
}

func shortestPathLength(graph [][]int) int {
	queue := []Tuple{}
	N := len(graph)
	mask := (1 << uint(N)) - 1
	visited := map[Tuple]bool{}
	for i := 0; i < N; i++ {
		tmp := Tuple{i, 0, 1 << uint(i)}
		queue = append(queue, tmp)
		visited[tmp] = true
	}
	for len(queue) != 0 {
		top := queue[0]
		queue = queue[1:]
		if top.mask == mask {
			return top.length
		}
		for _, neighbor := range graph[top.index] {
			tmp := Tuple{neighbor, top.length + 1, top.mask | 1<<uint(neighbor)}
			if !visited[tmp] {
				visited[tmp] = true
				queue = append(queue, tmp)
			}
		}
	}
	return -1
}

func main() {
	fmt.Println(shortestPathLength([][]int{
		{1, 2, 3},
		{0},
		{0},
		{0},
	}))
}
