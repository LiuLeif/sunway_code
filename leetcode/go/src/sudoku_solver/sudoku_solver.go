// 2017-11-16 16:53
package main

import "fmt"

func dfs(board [][]byte, rowSet, columnSet, areaSet []map[byte]bool, row, column int) bool {
	if row == 9 {
		return true
	}
	if board[row][column] != '.' {
		column++
		if column == 9 {
			row++
			column = 0
		}
		return dfs(board, rowSet, columnSet, areaSet, row, column)
	}
	area := area(row, column)
	for k := '1'; k <= '9'; k++ {
		b := byte(k)
		if rowSet[row][b] || columnSet[column][b] || areaSet[area][b] {
			continue
		}
		rowSet[row][b] = true
		columnSet[column][b] = true
		areaSet[area][b] = true
		board[row][column] = b

		nextColumn := column
		nextRow := row
		nextColumn++
		if nextColumn == 9 {
			nextRow++
			nextColumn = 0
		}

		if dfs(board, rowSet, columnSet, areaSet, nextRow, nextColumn) {
			return true
		}
		rowSet[row][b] = false
		columnSet[column][b] = false
		areaSet[area][b] = false
		board[row][column] = '.'
	}
	return false
}

func area(i, j int) int {
	x := i / 3
	y := j / 3
	return x*3 + y
}

func solveSudoku(board [][]byte) {
	rowSet := make([]map[byte]bool, 9)
	columnSet := make([]map[byte]bool, 9)
	areaSet := make([]map[byte]bool, 9)

	for i := 0; i < 9; i++ {
		rowSet[i] = map[byte]bool{}
		columnSet[i] = map[byte]bool{}
		areaSet[i] = map[byte]bool{}
	}

	for i := 0; i < 9; i++ {
		for j := 0; j < 9; j++ {
			if board[i][j] != '.' {
				rowSet[i][board[i][j]] = true
				columnSet[j][board[i][j]] = true
				areaSet[area(i, j)][board[i][j]] = true
			}
		}
	}

	dfs(board, rowSet, columnSet, areaSet, 0, 0)

}
func main() {
	board := []string{
		"..9748...",
		"7........",
		".2.1.9...",
		"..7...24.",
		".64.1.59.",
		".98...3..",
		"...8.3.2.",
		"........6",
		"...2759..",
	}
	tmp := make([][]byte, len(board))
	for i, v := range board {
		tmp[i] = []byte(v)
	}

	solveSudoku(tmp)
	for i := 0; i < len(tmp); i++ {
		fmt.Println(string(tmp[i]))
	}

}
