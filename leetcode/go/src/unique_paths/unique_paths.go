// 2017-12-20 17:40
package main

import "fmt"

type Pair struct {
	x, y int
}

func helper(m map[Pair]int, i, j, x, y int) int {
	if m[Pair{i, j}] != 0 {
		return m[Pair{i, j}]
	}
	if i == x && j == y {
		return 1
	}

	down := 0
	if i < x {
		down = helper(m, i+1, j, x, y)
	}
	right := 0
	if j < y {
		right = helper(m, i, j+1, x, y)
	}
	m[Pair{i, j}] = down + right
	return down + right
}

func uniquePaths(x int, y int) int {
	m := make(map[Pair]int)
	return helper(m, 0, 0, x-1, y-1)
}
func main() {
	fmt.Println(uniquePaths(10, 10))
}
