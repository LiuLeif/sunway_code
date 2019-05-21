// 2018-11-26 18:12
package main

import "fmt"

func removeStones(stones [][]int) int {
	// 0 0
	// 0   0
	//   0 0

	parent := map[int]int{}

	find := func(a int) int {
		if _, ok := parent[a]; !ok {
			parent[a] = a
		}
		for parent[a] != a {
			a = parent[a]
		}
		return a
	}

	union := func(a, b int) {
		pa, pb := find(a), find(b)
		parent[pa] = pb
	}

	rows := map[int][]int{}
	cols := map[int][]int{}
	for _, stone := range stones {
		if rows[stone[0]] == nil {
			rows[stone[0]] = []int{}
		}
		if cols[stone[1]] == nil {
			cols[stone[1]] = []int{}
		}
		rows[stone[0]] = append(rows[stone[0]], stone[1])
		cols[stone[1]] = append(cols[stone[1]], stone[0])
	}

	for r, _ := range rows {
		c0 := rows[r][0]
		for _, c := range rows[r][1:] {
			union(c0+10000*r, c+10000*r)
		}
	}
	for c, _ := range cols {
		r0 := cols[c][0]
		for _, r := range cols[c][1:] {
			union(c+10000*r0, c+10000*r)
		}
	}

	islands := map[int]bool{}
	for _, stone := range stones {
		islands[find(stone[0]*10000+stone[1])] = true
	}

	return len(stones) - len(islands)
}

func removeStones2(stones [][]int) int {
	// 0 0
	// 0   0
	//   0 0

	rows := map[int]map[int]bool{}
	cols := map[int]map[int]bool{}
	for _, stone := range stones {
		if rows[stone[0]] == nil {
			rows[stone[0]] = map[int]bool{}
		}
		if cols[stone[1]] == nil {
			cols[stone[1]] = map[int]bool{}
		}
		rows[stone[0]][stone[1]] = true
		cols[stone[1]][stone[0]] = true
	}

	island := 0
	var dfs func(i, j int)
	dfs = func(i, j int) {
		delete(rows[i], j)
		delete(cols[j], i)
		if len(rows[i]) != 0 {
			for jj, _ := range rows[i] {
				dfs(i, jj)
			}
		}
		if len(cols[j]) != 0 {
			for ii, _ := range cols[j] {
				dfs(ii, j)
			}
		}
	}

	for row, cols := range rows {
		for col, _ := range cols {
			if rows[row][col] {
				island++
				dfs(row, col)
			}
		}
	}
	return len(stones) - island
}

func main() {
	fmt.Println(removeStones([][]int{{0, 0}, {0, 1}, {1, 0}, {1, 2}, {2, 1}, {2, 2}}))
	fmt.Println(removeStones([][]int{{0, 0}, {0, 2}, {1, 1}, {2, 0}, {2, 2}}))
	fmt.Println(removeStones([][]int{{0, 0}}))
	fmt.Println(removeStones([][]int{{0, 0}, {0, 1}, {1, 0}, {1, 1}, {2, 1}, {2, 2}, {3, 2}, {3, 3}, {3, 4}, {4, 3}, {4, 4}}))
}
