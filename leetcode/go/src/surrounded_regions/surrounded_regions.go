// 2017-12-12 16:01
package main

import "fmt"

func dfs(board [][]byte, i, j int) {
	board[i][j] = 'Y'
	if i > 0 && board[i-1][j] == 'O' {
		dfs(board, i-1, j)
	}
	if i < len(board)-1 && board[i+1][j] == 'O' {
		dfs(board, i+1, j)
	}
	if j > 0 && board[i][j-1] == 'O' {
		dfs(board, i, j-1)
	}
	if j < len(board[0])-1 && board[i][j+1] == 'O' {
		dfs(board, i, j+1)
	}
}

func solve(board [][]byte) {
	for i := 0; i < len(board); i++ {
		for j := 0; j < len(board[0]); j++ {
			if i != 0 && i != len(board)-1 && j != 0 && j != len(board[0])-1 {
				continue
			}
			if board[i][j] == 'O' {
				dfs(board, i, j)
			}
		}
	}
	for i := 0; i < len(board); i++ {
		for j := 0; j < len(board[0]); j++ {
			if board[i][j] != 'Y' {
				board[i][j] = 'X'
			} else {
				board[i][j] = 'O'
			}
		}
	}

}
func main() {
	board := [][]byte{
		{'O', 'O'},
		{'O', 'O'},
	}
	solve(board)
	fmt.Println(board)
}
