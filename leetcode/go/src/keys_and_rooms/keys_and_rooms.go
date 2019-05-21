// 2018-10-17 11:04
package main

import "fmt"

func canVisitAllRooms(rooms [][]int) bool {
	N := len(rooms)
	visited := make([]bool, N)

	var dfs func(i int)
	dfs = func(i int) {
		if visited[i] {
			return
		}
		visited[i] = true
		for _, k := range rooms[i] {
			dfs(k)
		}
	}

	dfs(0)

	for _, v := range visited {
		if !v {
			return false
		}
	}
	return true
}
func main() {
	fmt.Println(canVisitAllRooms([][]int{
		{1, 3},
		{3, 0, 1},
		{2},
		{0},
	}))
}
