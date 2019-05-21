// 2017-12-01 11:55
package main

import "fmt"

func bfs(matrix [][]int, queue [][2]int, ocean [][]bool) {
	for len(queue) != 0 {
		top := queue[0]
		queue = queue[1:]
		x, y := top[0], top[1]
		if x > 0 && matrix[x][y] <= matrix[x-1][y] && !ocean[x-1][y] {
			queue = append(queue, [2]int{x - 1, y})
			ocean[x-1][y] = true
		}
		if y > 0 && matrix[x][y] <= matrix[x][y-1] && !ocean[x][y-1] {
			queue = append(queue, [2]int{x, y - 1})
			ocean[x][y-1] = true
		}
		if x < len(matrix)-1 && matrix[x][y] <= matrix[x+1][y] && !ocean[x+1][y] {
			queue = append(queue, [2]int{x + 1, y})
			ocean[x+1][y] = true
		}
		if y < len(matrix[0])-1 && matrix[x][y] <= matrix[x][y+1] && !ocean[x][y+1] {
			queue = append(queue, [2]int{x, y + 1})
			ocean[x][y+1] = true
		}
	}
}

func pacificAtlantic(matrix [][]int) [][]int {
	if len(matrix) == 0 {
		return [][]int{}
	}
	pacific := make([][]bool, len(matrix))
	atlantic := make([][]bool, len(matrix))

	queuePacific := [][2]int{}
	queueAtlantic := [][2]int{}

	for i := 0; i < len(matrix); i++ {
		pacific[i] = make([]bool, len(matrix[0]))
		atlantic[i] = make([]bool, len(matrix[0]))
	}
	for i := 0; i < len(matrix); i++ {
		pacific[i][0] = true
		atlantic[i][len(matrix[0])-1] = true
		queuePacific = append(queuePacific, [2]int{i, 0})
		queueAtlantic = append(queueAtlantic, [2]int{i, len(matrix[0]) - 1})
	}
	for i := 0; i < len(matrix[0]); i++ {
		pacific[0][i] = true
		atlantic[len(matrix)-1][i] = true
		queuePacific = append(queuePacific, [2]int{0, i})
		queueAtlantic = append(queueAtlantic, [2]int{len(matrix) - 1, i})
	}

	bfs(matrix, queuePacific, pacific)
	bfs(matrix, queueAtlantic, atlantic)
	ret := [][]int{}
	for i := 0; i < len(matrix); i++ {
		for j := 0; j < len(matrix[0]); j++ {
			if pacific[i][j] && atlantic[i][j] {
				ret = append(ret, []int{i, j})
			}
		}
	}

	return ret
}
func main() {
	matrix := [][]int{
		{1, 1},
		{1, 1},
		{1, 1},
	}

	fmt.Println(pacificAtlantic(matrix))
}
