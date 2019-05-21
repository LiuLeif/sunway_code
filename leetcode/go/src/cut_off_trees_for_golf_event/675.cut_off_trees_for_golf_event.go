// 2018-12-11 13:43
package main

import (
	"container/heap"
	"fmt"
)

type Tuple struct {
	point  []int
	height int
}

type MinHeap []Tuple

func (h MinHeap) Len() int {
	return len(h)
}

func (h MinHeap) Less(i, j int) bool {
	return h[i].height < h[j].height
}

func (h MinHeap) Swap(i, j int) {
	h[i], h[j] = h[j], h[i]
}

func (ph *MinHeap) Pop() interface{} {
	t := (*ph)[len(*ph)-1]
	*ph = (*ph)[:len(*ph)-1]
	return t
}

func (ph *MinHeap) Push(x interface{}) {
	*ph = append(*ph, x.(Tuple))
}

func cutOffTree(forest [][]int) int {
	queue := MinHeap{}
	M := len(forest)
	N := len(forest[0])
	for i := 0; i < M; i++ {
		for j := 0; j < N; j++ {
			if forest[i][j] > 1 {
				heap.Push(&queue, Tuple{[]int{i, j}, forest[i][j]})
			}
		}
	}

	bfs := func(f, t []int, visited [][]bool) int {
		queue := [][]int{[]int{f[0], f[1], 0}}
		for len(queue) != 0 {
			top := queue[0]
			queue = queue[1:]
			if top[0] == t[0] && top[1] == t[1] {
				return top[2]
			}
			directions := [][2]int{{-1, 0}, {1, 0}, {0, -1}, {0, 1}}
			for _, d := range directions {
				x, y := top[0]+d[0], top[1]+d[1]
				if x < 0 || x >= M || y < 0 || y >= N {
					continue
				}
				if forest[x][y] == 0 {
					continue
				}
				if visited[x][y] {
					continue
				}
				visited[x][y] = true
				queue = append(queue, []int{x, y, top[2] + 1})
			}
		}
		return -1
	}

	ret := 0
	curr := []int{0, 0}
	for len(queue) != 0 {
		top := heap.Pop(&queue).(Tuple).point
		visited := make([][]bool, M)
		for i := 0; i < M; i++ {
			visited[i] = make([]bool, N)
		}
		visited[curr[0]][curr[1]] = true
		tmp := bfs(curr, top, visited)
		if tmp == -1 {
			return -1
		}
		ret += tmp
		curr = top
	}
	return ret
}

func main() {
	fmt.Println(cutOffTree([][]int{
		{1, 2, 3},
		{0, 0, 4},
		{7, 6, 5},
	}))
}
