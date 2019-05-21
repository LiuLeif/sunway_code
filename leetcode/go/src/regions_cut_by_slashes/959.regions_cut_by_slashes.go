// 2018-12-18 11:15
package main

import "fmt"

func regionsBySlashes(grid []string) int {
	//
	//    +-------------+
	//    |  \-   0    -|
	//    |    \-   -/  |
	//    | 1   --/  3  |
	//    |   --/  \-   |
	//    | -/   2   \- |
	//    +/-----------\+
	N := len(grid)
	ret := N * N * 4
	g := func(a, b, c int) int {
		return (a*N+b)*4 + c
	}
	parent := map[int]int{}

	find := func(x int) int {
		if _, ok := parent[x]; !ok {
			parent[x] = x
		}
		for parent[x] != x {
			x = parent[x]
		}
		return x
	}
	union := func(x, y int) {
		x, y = find(x), find(y)
		if x == y {
			return
		}
		parent[x] = y
		ret--
	}

	for i := 0; i < N; i++ {
		for j := 0; j < N; j++ {
			if i > 0 {
				union(g(i, j, 0), g(i-1, j, 2))
			}
			if j > 0 {
				union(g(i, j, 1), g(i, j-1, 3))
			}
			if grid[i][j] != '/' {
				union(g(i, j, 1), g(i, j, 2))
				union(g(i, j, 0), g(i, j, 3))
			}
			if grid[i][j] != '\\' {
				union(g(i, j, 0), g(i, j, 1))
				union(g(i, j, 2), g(i, j, 3))
			}
		}
	}
	return ret
}

func main() {
	fmt.Println(regionsBySlashes([]string{
		"/\\",
		"\\/",
	}))
}
