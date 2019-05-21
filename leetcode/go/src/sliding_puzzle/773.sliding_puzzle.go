// 2018-11-07 09:47
package main

import "fmt"

type Tuple struct {
	array [2][3]int
	x, y  int
}

func slidingPuzzle(board [][]int) int {
	boardArray := [2][3]int{}
	x, y := 0, 0
	for i := 0; i < 2; i++ {
		for j := 0; j < 3; j++ {
			boardArray[i][j] = board[i][j]
			if board[i][j] == 0 {
				x, y = i, j
			}
		}
	}

	targetBoardArray := [2][3]int{
		{1, 2, 3},
		{4, 5, 0},
	}

	queue := []Tuple{{boardArray, x, y}, {[2][3]int{}, -1, -1}}
	visited := map[Tuple]bool{queue[0]: true}

	depth := 0
	for len(queue) != 1 {
		top := queue[0]
		queue = queue[1:]
		if top.x == -1 {
			queue = append(queue, Tuple{[2][3]int{}, -1, -1})
			depth++
			continue
		}
		if top.array == targetBoardArray {
			return depth
		}
		for _, k := range []int{-1, 0, 1} {
			for _, v := range []int{-1, 0, 1} {
				if (k == 0 && v == 0) || k*v != 0 {
					continue
				}
				x, y := top.x+k, top.y+v
				if x < 0 || x > 1 || y < 0 || y > 2 {
					continue
				}
				newArray := top.array
				newArray[top.x][top.y], newArray[x][y] = newArray[x][y], newArray[top.x][top.y]
				tmp := Tuple{newArray, x, y}
				if !visited[tmp] {
					visited[tmp] = true
					queue = append(queue, tmp)
				}
			}
		}
	}

	return -1
}

func main() {
	fmt.Println(slidingPuzzle([][]int{
		{1, 2, 3},
		{5, 4, 0},
	}))
}
