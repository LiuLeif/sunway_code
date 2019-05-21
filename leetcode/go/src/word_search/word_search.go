// 2017-11-14 18:54
package main

import "fmt"

func dfs(word string, board [][]byte, visited [][]bool, i, j int) bool {
	if board[i][j] != word[0] {
		return false
	}
	if len(word) == 1 {
		return true
	}
	row := len(board)
	column := len(board[0])
	visited[i][j] = true
	if i > 0 && visited[i-1][j] == false {
		if dfs(word[1:], board, visited, i-1, j) {
			return true
		}
	}
	if i < row-1 && visited[i+1][j] == false {
		if dfs(word[1:], board, visited, i+1, j) {
			return true
		}
	}
	if j > 0 && visited[i][j-1] == false {
		if dfs(word[1:], board, visited, i, j-1) {
			return true
		}
	}
	if j < column-1 && visited[i][j+1] == false {
		if dfs(word[1:], board, visited, i, j+1) {
			return true
		}
	}
	visited[i][j] = false
	return false
}

func exist(board [][]byte, word string) bool {
	row := len(board)
	column := len(board[0])
	visisted := make([][]bool, row)
	for i := 0; i < row; i++ {
		visisted[i] = make([]bool, column)
	}
	for i := 0; i < row; i++ {
		for j := 0; j < column; j++ {
			if dfs(word, board, visisted, i, j) {
				return true
			}
		}
	}

	return false
}
func main() {
	board := make([][]byte, 3)
	board[0] = []byte{'A', 'B', 'C', 'E'}
	board[1] = []byte{'S', 'F', 'C', 'S'}
	board[2] = []byte{'A', 'D', 'E', 'E'}
	fmt.Println(exist(board, "ABCESCFSADEE"))
}
