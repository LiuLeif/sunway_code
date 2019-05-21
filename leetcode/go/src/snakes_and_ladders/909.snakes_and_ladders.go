// 2018-10-30 14:32
package main

import "fmt"

func snakesAndLadders(board [][]int) int {
	N := len(board)
	flatBoard := []int{0}
	for i := N - 1; i >= 0; i-- {
		if (N-1-i)&1 == 0 {
			flatBoard = append(flatBoard, board[i]...)
		} else {
			// rev
			for j := N - 1; j >= 0; j-- {
				flatBoard = append(flatBoard, board[i][j])
			}
		}
	}

	follow := func(x int) int {
		if flatBoard[x] > 0 {
			return flatBoard[x]
		}
		return x
	}

	visited := make([]bool, len(flatBoard))
	queue := []int{1, -1}
	move := 1
	for len(queue) != 1 {
		top := queue[0]
		queue = queue[1:]
		if top == -1 {
			move++
			queue = append(queue, -1)
			continue
		}
		for i := 1; i < 7; i++ {
			next := follow(top + i)
			if next == N*N {
				return move
			}

			if !visited[next] {
				visited[next] = true
				queue = append(queue, next)
			}
		}
	}

	return -1
}

func main() {
	fmt.Println(snakesAndLadders([][]int{
		{-1, -1, -1, 46, 47, -1, -1, -1},
		{51, -1, -1, 63, -1, 31, 21, -1},
		{-1, -1, 26, -1, -1, 38, -1, -1},
		{-1, -1, 11, -1, 14, 23, 56, 57},
		{11, -1, -1, -1, 49, 36, -1, 48},
		{-1, -1, -1, 33, 56, -1, 57, 21},
		{-1, -1, -1, -1, -1, -1, 2, -1},
		{-1, -1, -1, 8, 3, -1, 6, 56},
	}))
}
