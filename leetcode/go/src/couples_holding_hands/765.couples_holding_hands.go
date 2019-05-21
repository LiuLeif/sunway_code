// 2018-11-07 10:36
package main

import "fmt"

func minSwapsCouples(row []int) int {
	N := len(row)
	conn := map[int]map[int]bool{}

	for i := 0; i < N; i += 2 {
		t1 := row[i] / 2
		t2 := row[i+1] / 2
		if t1 != t2 {
			if conn[t1] == nil {
				conn[t1] = map[int]bool{}
			}
			conn[t1][t2] = true
			if conn[t2] == nil {
				conn[t2] = map[int]bool{}
			}
			conn[t2][t1] = true
		}
	}
	visited := make([]bool, N)
	count := 0

	var dfs func(x int)
	dfs = func(x int) {
		if visited[x] {
			return
		}
		visited[x] = true
		for k, _ := range conn[x] {
			dfs(k)
		}
	}
	for i := 0; i < N/2; i++ {
		if !visited[i] {
			count++
			dfs(i)
		}
	}
	return N/2 - count
}

func main() {
	fmt.Println(minSwapsCouples([]int{
		0, 2, 1, 3,
	}))
}

// 0 2 4 3 5 1
