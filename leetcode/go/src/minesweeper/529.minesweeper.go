// 2018-10-23 18:33
package main

import "fmt"

func updateBoard(board [][]byte, click []int) [][]byte {
	if board[click[0]][click[1]] == 'M' {
		board[click[0]][click[1]] = 'X'
		return board
	}
	mine := make([][]int, len(board))
	for i := 0; i < len(board); i++ {
		mine[i] = make([]int, len(board[0]))
	}

	for i := 0; i < len(board); i++ {
		for j := 0; j < len(board[0]); j++ {
			if board[i][j] == 'M' {
				for k := -1; k < 2; k++ {
					for g := -1; g < 2; g++ {
						if i+k >= 0 && j+g >= 0 && i+k < len(board) && j+g < len(board[0]) {
							mine[i+k][j+g]++
						}
					}
				}
			}
		}
	}

	visited := make([][]bool, len(board))
	for i := 0; i < len(board); i++ {
		visited[i] = make([]bool, len(board[0]))
	}

	var dfs func(i, j int)
	dfs = func(i, j int) {
		if i < 0 || i >= len(board) || j < 0 || j >= len(board[0]) {
			return
		}
		if visited[i][j] {
			return
		}
		visited[i][j] = true
		if board[i][j] == 'E' {
			if mine[i][j] != 0 {
				board[i][j] = byte(mine[i][j] + 48)
			} else {
				board[i][j] = 'B'
				dfs(i-1, j)
				dfs(i+1, j)
				dfs(i, j-1)
				dfs(i, j+1)
				dfs(i-1, j-1)
				dfs(i-1, j+1)
				dfs(i+1, j-1)
				dfs(i+1, j+1)
			}
		}

	}
	dfs(click[0], click[1])
	// for i := 0; i < len(board); i++ {
	// 	fmt.Println(string(board[i]))
	// }
	return board
}

func main() {
	fmt.Println(updateBoard([][]byte{
		{'E', 'E', 'E', 'E', 'E', 'E', 'E', 'E'},
		{'E', 'E', 'E', 'E', 'E', 'E', 'E', 'M'},
		{'E', 'E', 'M', 'E', 'E', 'E', 'E', 'E'},
		{'M', 'E', 'E', 'E', 'E', 'E', 'E', 'E'},
		{'E', 'E', 'E', 'E', 'E', 'E', 'E', 'E'},
		{'E', 'E', 'E', 'E', 'E', 'E', 'E', 'E'},
		{'E', 'E', 'E', 'E', 'E', 'E', 'E', 'E'},
		{'E', 'E', 'M', 'M', 'E', 'E', 'E', 'E'},
	}, []int{0, 0}))
}
