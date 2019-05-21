// 2018-11-23 13:54
package main

import (
	"fmt"
	"math"
)

func movesToChessboard(board [][]int) int {
	min := func(nums ...int) int {
		ret := math.MaxInt32
		for _, n := range nums {
			if ret > n {
				ret = n
			}
		}
		return ret
	}

	N := len(board)
	for i := 0; i < N; i++ {
		for j := 0; j < N; j++ {
			if board[0][0]^board[i][j]^board[i][0]^board[0][j] != 0 {
				return -1
			}
		}
	}
	rowSum, colSum := 0, 0
	rowSwap, colSwap := 0, 0
	for i := 0; i < N; i++ {
		rowSum += board[i][0]
		colSum += board[0][i]
		if board[i][0] == i%2 {
			rowSwap++
		}
		if board[0][i] == i%2 {
			colSwap++
		}
	}
	if rowSum < N/2 || rowSum > (N+1)/2 {
		return -1
	}
	if colSum < N/2 || colSum > (N+1)/2 {
		return -1
	}
	if N%2 == 0 {
		rowSwap = min(rowSwap, N-rowSwap)
		colSwap = min(colSwap, N-colSwap)
	} else {
		if rowSwap%2 == 1 {
			rowSwap = N - rowSwap
		}
		if colSwap%2 == 1 {
			colSwap = N - colSwap
		}
	}
	return (rowSwap + colSwap) / 2
}

func main() {
	// 0 1 0 1 0 1 0
	// 1 0 1 0 1 0 1
	// 0 1 1 1 0 0 0
	//
	fmt.Println(movesToChessboard([][]int{
		{0, 1, 1, 0}, {0, 1, 1, 0}, {1, 0, 0, 1}, {1, 0, 0, 1},
	}))
}
