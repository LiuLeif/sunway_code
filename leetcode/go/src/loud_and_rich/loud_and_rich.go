// 2018-10-18 15:37
package main

import "fmt"

func loudAndRich(richer [][]int, quiet []int) []int {
	ret := make([]int, len(quiet))
	conn := map[int]map[int]bool{}

	cache := make([]int, len(quiet))

	var dfs func(int) int
	dfs = func(start int) int {
		if cache[start] != 0 {
			return cache[start]
		}
		ret := start
		for k, _ := range conn[start] {
			tmp := dfs(k)
			if quiet[tmp] < quiet[ret] {
				ret = tmp
			}
		}
		cache[start] = ret
		return ret
	}

	for i := 0; i < len(richer); i++ {
		x := richer[i][0]
		y := richer[i][1]
		if conn[y] == nil {
			conn[y] = map[int]bool{}
		}
		conn[y][x] = true
	}

	for i := 0; i < len(quiet); i++ {
		ret[i] = dfs(i)
	}
	return ret
}
func main() {
	fmt.Println(loudAndRich([][]int{
		{1, 0},
		{2, 1},
		{3, 1},
		{3, 7},
		{4, 3},
		{5, 3},
		{6, 3},
	}, []int{3, 2, 5, 4, 6, 1, 7, 0}))
}
