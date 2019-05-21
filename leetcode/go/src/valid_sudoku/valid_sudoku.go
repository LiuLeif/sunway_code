// 2018-01-26 12:36
package main

import "fmt"

func isValidSudoku(board [][]byte) bool {
	mH := make([]map[byte]bool, 9)
	mV := make([]map[byte]bool, 9)
	mX := make([]map[byte]bool, 9)
	for i := 0; i < 9; i++ {
		mH[i] = map[byte]bool{}
		mV[i] = map[byte]bool{}
		mX[i] = map[byte]bool{}
	}

	for i := 0; i < 9; i++ {
		for j := 0; j < 9; j++ {
			c := board[i][j]
			if c != '.' {
				// H
				if mH[i][c] {
					return false
				}
				mH[i][c] = true
				// V
				if mV[j][c] {
					return false
				}
				mV[j][c] = true
				// X
				xIndex := (i/3)*3 + j/3
				if mX[xIndex][c] {
					return false
				}
				mX[xIndex][c] = true
			}
		}
	}

	return true
}
func main() {
	fmt.Println(isValidSudoku([][]byte{
		[]byte(".87654321"),
		[]byte("2........"),
		[]byte("3........"),
		[]byte("4........"),
		[]byte("5........"),
		[]byte("6........"),
		[]byte("7........"),
		[]byte("8........"),
		[]byte("9........"),
	}))
}
