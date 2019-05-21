// 2018-11-01 15:55
package main

import "fmt"

func escapeGhosts(ghosts [][]int, target []int) bool {
	abs := func(x int) int {
		if x < 0 {
			return -x
		}
		return x
	}
	distance := func(p1, p2 []int) int {
		return abs(p1[0]-p2[0]) + abs(p1[1]-p2[1])
	}

	d := distance([]int{0, 0}, target)
	if d == 0 {
		return true
	}
	for _, g := range ghosts {
		if distance(g, target) <= d {
			return false
		}
	}
	return true
}

func main() {
	fmt.Println(escapeGhosts([][]int{
		{2, 0},
	}, []int{1, 0}))
}
