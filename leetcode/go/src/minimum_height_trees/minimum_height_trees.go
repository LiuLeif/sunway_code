// 2017-12-07 14:07
package main

import "fmt"

func findMinHeightTrees(n int, edges [][]int) []int {
	ingress := make([]int, n)
	neighbors := make(map[int][]int)
	for i := 0; i < len(edges); i++ {
		a := edges[i][0]
		b := edges[i][1]
		ingress[a]++
		ingress[b]++
		if neighbors[a] == nil {
			neighbors[a] = []int{}
		}
		neighbors[a] = append(neighbors[a], b)
		if neighbors[b] == nil {
			neighbors[b] = []int{}
		}
		neighbors[b] = append(neighbors[b], a)
	}

	remain := map[int]bool{}
	for i := 0; i < n; i++ {
		remain[i] = true
	}

	queue := []int{-1}
	for i := 0; i < n; i++ {
		if ingress[i] == 1 {
			queue = append(queue, i)
		}
	}

	for len(queue) != 1 {
		top := queue[0]
		queue = queue[1:]
		if top == -1 {
			if len(remain) <= 2 {
				break
			}
			queue = append(queue, -1)
			continue
		}
		delete(remain, top)
		for _, v := range neighbors[top] {
			ingress[v]--
			if ingress[v] == 1 {
				queue = append(queue, v)
			}
		}
	}

	ret := []int{}
	for k, _ := range remain {
		ret = append(ret, k)
	}
	return ret
}
func main() {
	edges := [][]int{
		{0, 1},
	}
	fmt.Println(findMinHeightTrees(2, edges))
}
