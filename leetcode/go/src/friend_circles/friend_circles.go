// 2017-11-30 11:48
package main

import "fmt"

func dfs(visited map[int]bool, matrix [][]int, i int) bool {
	if visited[i] {
		return false
	}
	visited[i] = true
	for j := 0; j < len(matrix[i]); j++ {
		dfs(visited, matrix, matrix[i][j])
	}
	return true
}

func findCircleNum(matrix [][]int) int {
	visited := make(map[int]bool)
	ret := 0
	data := make([][]int, len(matrix))
	for i := 0; i < len(matrix); i++ {
		data[i] = []int{}
	}

	for i := 0; i < len(matrix); i++ {
		for j := i + 1; j < len(matrix); j++ {
			if matrix[i][j] == 1 {
				data[i] = append(data[i], j)
				data[j] = append(data[j], i)
			}
		}
	}

	for i := 0; i < len(matrix); i++ {
		if dfs(visited, data, i) {
			ret++
		}
	}

	return ret
}
func main() {
	matrix := [][]int{
		{1, 0, 0, 1},
		{0, 1, 1, 0},
		{0, 1, 1, 1},
		{1, 0, 1, 1},
	}

	fmt.Println(findCircleNum(matrix))
}
