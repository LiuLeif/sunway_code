// 2018-03-01 15:41
package main

import "fmt"

func gameOfLife(board [][]int) {
	countNeighbor := func(i, j int) int {
		ret := 0
		steps := []int{-1, 0, 1}
		for _, stepX := range steps {
			for _, stepY := range steps {
				if stepX == 0 && stepY == 0 {
					continue
				}
				if !(i+stepX >= 0 && i+stepX < len(board) && j+stepY >= 0 && j+stepY < len(board[0])) {
					continue
				}
				if board[i+stepX][j+stepY] == 1 || board[i+stepX][j+stepY] == 3 {
					ret++
				}
			}
		}
		return ret
	}
	for i := 0; i < len(board); i++ {
		for j := 0; j < len(board[0]); j++ {
			neighbor := countNeighbor(i, j)
			if board[i][j] == 0 {
				if neighbor == 3 {
					board[i][j] = 2 // mark live
				}
			} else if board[i][j] == 1 {
				if neighbor < 2 || neighbor > 3 {
					board[i][j] = 3 // mark dead
				}
			}
		}
	}
	for i := 0; i < len(board); i++ {
		for j := 0; j < len(board[0]); j++ {
			if board[i][j] == 2 {
				board[i][j] = 1
			} else if board[i][j] == 3 {
				board[i][j] = 0
			}
		}
	}
}
func main() {
	// 4x4
	board := [][]int{
		[]int{0, 0, 0, 0, 0},
		[]int{0, 0, 0, 0, 0},
		[]int{0, 1, 1, 1, 0},
		[]int{0, 0, 0, 0, 0},
		[]int{0, 0, 0, 0, 0},
	}
	gameOfLife(board)
	fmt.Println(board)
}
