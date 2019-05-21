// 2018-12-11 16:07
package main

import "fmt"

func findRedundantDirectedConnection(edges [][]int) []int {
	parents := map[int][]int{}
	e1, e2 := []int{}, []int{}
	for _, e := range edges {
		if parents[e[1]] == nil {
			parents[e[1]] = []int{}
		}
		parents[e[1]] = append(parents[e[1]], e[0])
		if len(parents[e[1]]) == 2 {
			e1, e2 = []int{parents[e[1]][0], e[1]}, []int{parents[e[1]][1], e[1]}
		}
	}
	// fmt.Println(parents)
	// fmt.Println(e1, e2)
	parent := map[int]int{}
	rank := map[int]int{}

	find := func(a int) int {
		if _, ok := parent[a]; !ok {
			parent[a] = a
		}
		for parent[a] != a {
			a = parent[a]
		}
		return a
	}

	union := func(a, b int) bool {
		a, b = find(a), find(b)
		if a == b {
			return true
		}
		if rank[a] > rank[b] {
			parent[b] = a
		} else if rank[a] < rank[b] {
			parent[a] = b
		} else {
			parent[a] = b
			rank[a]++
		}
		return false
	}

	for _, e := range edges {
		if len(e1) != 0 && ((e[0] == e1[0] && e[1] == e1[1]) || (e[0] == e2[0] && e[1] == e2[1])) {
			continue
		}
		if union(e[1], e[0]) {
			return e
		}
	}

	if union(e1[1], e1[0]) {
		return e1
	}
	return e2
}

func main() {
	// fmt.Println(findRedundantDirectedConnection([][]int{{1, 2}, {2, 3}, {3, 4}, {4, 1}, {1, 5}}))
	fmt.Println(findRedundantDirectedConnection([][]int{{1, 2}, {1, 3}, {2, 3}}))
	fmt.Println(findRedundantDirectedConnection([][]int{{2, 1}, {3, 1}, {4, 2}, {1, 4}}))
}
