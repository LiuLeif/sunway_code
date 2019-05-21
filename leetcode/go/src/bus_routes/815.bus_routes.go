// 2018-11-12 15:37
package main

import "fmt"

func numBusesToDestination(routes [][]int, S int, T int) int {
	if S == T {
		return 0
	}

	N := len(routes)
	m := map[int][]int{}
	neighbor := map[int]map[int]bool{}
	for i := 0; i < N; i++ {
		for j := 0; j < len(routes[i]); j++ {
			stop := routes[i][j]
			if m[routes[i][j]] == nil {
				m[stop] = []int{}
			}
			if neighbor[i] == nil {
				neighbor[i] = map[int]bool{}
			}
			for k := 0; k < len(m[stop]); k++ {
				neighbor[i][m[stop][k]] = true
				neighbor[m[stop][k]][i] = true
			}
			m[stop] = append(m[stop], i)
		}
	}
	queue := m[S]
	visited := map[int]bool{}
	for _, v := range queue {
		visited[v] = true
	}

	target := map[int]bool{}
	for i := 0; i < len(m[T]); i++ {
		target[m[T][i]] = true
	}

	// fmt.Println(neighbor)
	// fmt.Println(queue, target)

	queue = append(queue, -1)
	ret := 1
	for len(queue) != 1 {
		// fmt.Println(queue)
		top := queue[0]
		queue = queue[1:]

		if top == -1 {
			ret++
			queue = append(queue, -1)
			continue
		}

		if target[top] {
			return ret
		}

		for k, _ := range neighbor[top] {
			if !visited[k] {
				visited[k] = true
				queue = append(queue, k)
			}
		}
	}
	return -1
}

func main() {
	fmt.Println(numBusesToDestination([][]int{{7, 12}, {4, 5, 15}, {6}, {15, 19}, {9, 12, 13}}, 15, 12))
	fmt.Println(numBusesToDestination([][]int{{1, 2, 7}, {3, 6, 7}}, 1, 6))
	fmt.Println(numBusesToDestination(
		[][]int{
			{1, 9, 12, 20, 23, 24, 35, 38},
			{10, 21, 24, 31, 32, 34, 37, 38, 43},
			{10, 19, 28, 37},
			{8},
			{14, 19},
			{11, 17, 23, 31, 41, 43, 44},
			{21, 26, 29, 33},
			{5, 11, 33, 41},
			{4, 5, 8, 9, 24, 44},
		},
		37, 28))
}
