// 2018-11-09 14:08
package main

import "fmt"

func flipLights(n int, m int) int {
	if m == 0 {
		return 1
	}
	if n >= 4 {
		n = 4
	}
	a := func(x int) int {
		for i := 0; i < n; i++ {
			x ^= (1 << uint(i))
		}
		return x
	}

	b := func(x int) int {
		for i := 0; i*2 < n; i++ {
			x ^= (1 << uint(i*2))
		}
		return x
	}
	c := func(x int) int {
		for i := 0; i*2+1 < n; i++ {
			x ^= (1 << uint(i*2+1))
		}
		return x
	}
	d := func(x int) int {
		for i := 0; i*3 < n; i++ {
			x ^= (1 << uint(i*3))
		}
		return x
	}
	queue := []int{0, -1}
	visited := map[int]bool{}

	for len(queue) != 1 {
		top := queue[0]
		queue = queue[1:]
		if top == -1 {
			m--
			if m == 0 {
				break
			}
			visited = map[int]bool{}
			queue = append(queue, -1)
			continue
		}
		for _, t := range []int{a(top), b(top), c(top), d(top)} {
			if !visited[t] {
				visited[t] = true
				queue = append(queue, t)
			}
		}
	}
	return len(queue)
}

func main() {
	fmt.Println(flipLights(1, 1))
}
