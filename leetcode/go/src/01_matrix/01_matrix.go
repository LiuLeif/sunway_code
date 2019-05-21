// 2017-12-20 14:33
package main

import (
	"fmt"
)

type Pair struct {
	i, j int
}

func updateMatrix(matrix [][]int) [][]int {
	ret := make([][]int, len(matrix))
	for i := 0; i < len(matrix); i++ {
		ret[i] = make([]int, len(matrix[0]))
		for j := 0; j < len(matrix[0]); j++ {
			ret[i][j] = -1
		}
	}
	queue := []Pair{}
	for i := 0; i < len(matrix); i++ {
		for j := 0; j < len(matrix[0]); j++ {
			if matrix[i][j] == 0 {
				ret[i][j] = 0
				queue = append(queue, Pair{i, j})
			}
		}
	}
	queue = append(queue, Pair{-1, -1})
	distance := 1
	for len(queue) != 1 {
		top := queue[0]
		queue = queue[1:]
		if top.i == -1 {
			queue = append(queue, Pair{-1, -1})
			distance += 1
			continue
		}
		if top.i > 0 && matrix[top.i-1][top.j] == 1 && ret[top.i-1][top.j] == -1 {
			queue = append(queue, Pair{top.i - 1, top.j})
			ret[top.i-1][top.j] = distance
		}
		if top.i < len(matrix)-1 && matrix[top.i+1][top.j] == 1 && ret[top.i+1][top.j] == -1 {
			queue = append(queue, Pair{top.i + 1, top.j})
			ret[top.i+1][top.j] = distance
		}
		if top.j > 0 && matrix[top.i][top.j-1] == 1 && ret[top.i][top.j-1] == -1 {
			queue = append(queue, Pair{top.i, top.j - 1})
			ret[top.i][top.j-1] = distance
		}
		if top.j < len(matrix[0])-1 && matrix[top.i][top.j+1] == 1 && ret[top.i][top.j+1] == -1 {
			queue = append(queue, Pair{top.i, top.j + 1})
			ret[top.i][top.j+1] = distance
		}
	}

	return ret
}
func main() {
	matrix := [][]int{
		{0, 0, 1, 0, 1, 1, 1, 0, 1, 1},
		{1, 1, 1, 1, 0, 1, 1, 1, 1, 1},
		{1, 1, 1, 1, 1, 0, 0, 0, 1, 1},
		{1, 0, 1, 0, 1, 1, 1, 0, 1, 1},
		{0, 0, 1, 1, 1, 0, 1, 1, 1, 1},
		{1, 0, 1, 1, 1, 1, 1, 1, 1, 1},
		{1, 1, 1, 1, 0, 1, 0, 1, 0, 1},
		{0, 1, 0, 0, 0, 1, 0, 0, 1, 1},
		{1, 1, 1, 0, 1, 1, 0, 1, 0, 1},
		{1, 0, 1, 1, 1, 0, 1, 1, 1, 0},
	}

	fmt.Println(updateMatrix(matrix))
}
