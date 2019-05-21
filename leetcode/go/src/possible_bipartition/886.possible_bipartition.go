// 2018-10-24 09:42
package main

import "fmt"

func possibleBipartition(N int, dislikes [][]int) bool {
	m := map[int][]int{}
	for _, dislike := range dislikes {
		if m[dislike[0]] == nil {
			m[dislike[0]] = []int{}
		}
		if m[dislike[1]] == nil {
			m[dislike[1]] = []int{}
		}
		m[dislike[0]] = append(m[dislike[0]], dislike[1])
		m[dislike[1]] = append(m[dislike[1]], dislike[0])
	}

	color := map[int]int{}
	queue := []int{}
	for i := 1; i < N+1; i++ {
		if color[i] == 0 {
			queue = append(queue, i)
		}
		for len(queue) != 0 {
			top := queue[0]
			queue = queue[1:]

			if color[top] == 0 {
				color[top] = 1
			}
			for _, n := range m[top] {
				if color[n] == color[top] {
					return false
				}
				if color[n] == 0 {
					color[n] = -color[top]
					queue = append(queue, n)
				}
			}
		}
	}

	return true
}

func main() {
	fmt.Println(possibleBipartition(10, [][]int{
		{4, 7},
		{4, 8},
		{2, 8},
		{8, 9},
		{1, 6},
		{5, 8},
		{1, 2},
		{6, 7},
		{3, 10},
		{8, 10},
		{1, 5},
		{7, 10},
		{1, 10},
		{3, 5},
		{3, 6},
		{1, 4},
		{3, 9},
		{2, 3},
		{1, 9},
		{7, 9},
		{2, 7},
		{6, 8},
		{5, 7},
		{3, 4},
	}))
}
