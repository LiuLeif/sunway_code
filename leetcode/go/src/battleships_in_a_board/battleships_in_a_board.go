// 2018-01-18 11:21
package main

import "fmt"

func countBattleships(board [][]byte) int {
	total, boundary := 0, 0
	for i := 0; i < len(board); i++ {
		for j := 0; j < len(board[0]); j++ {
			if board[i][j] == '.' {
				continue
			}
			total++
			if i != 0 && board[i-1][j] == 'X' {
				boundary++
			}
			if i != len(board)-1 && board[i+1][j] == 'X' {
				boundary++
			}
			if j != 0 && board[i][j-1] == 'X' {
				boundary++
			}
			if j != len(board[0])-1 && board[i][j+1] == 'X' {
				boundary++
			}
		}
	}

	return total - boundary/2
}
func main() {
	fmt.Println(countBattleships([][]byte{
		[]byte("X..X"),
		[]byte("...X"),
		[]byte("...X"),
		[]byte("...X"),
	}))
}
