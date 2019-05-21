// 2018-10-16 15:43
package main

import "fmt"

func numMagicSquaresInside(grid [][]int) int {
	ok := func(x, y int) bool {
		bits := 0
		diag := 0
		for i := 0; i < 3; i++ {
			row := 0
			col := 0
			for j := 0; j < 3; j++ {
				// dup
				if grid[i+x][j+y] < 1 || grid[i+x][j+y] > 9 {
					return false
				}
				bit := 1 << uint(grid[i+x][j+y])
				if bits&bit != 0 {
					return false
				}
				bit |= bit
				// row
				if i != 0 {
					row += grid[i+x][j+y] - grid[i-1+x][j+y]
				}
				// col
				if i != 0 {
					col += grid[j+x][i+y] - grid[j+x][i-1+y]
				}
				// diag
				if i == j {
					diag += grid[i+x][j+y]
				}
				if i+j == 2 {
					diag -= grid[i+x][j+y]
				}
			}
			if row != 0 || col != 0 {
				return false
			}
		}
		if diag != 0 {
			return false
		}
		return true
	}
	ret := 0
	for i := 0; i < len(grid)-2; i++ {
		for j := 0; j < len(grid[0])-2; j++ {
			if ok(i, j) {
				fmt.Println(i, j)
				ret++
			}
		}
	}
	return ret
}
func main() {
	fmt.Println(numMagicSquaresInside([][]int{
		{7, 6, 2, 2, 4},
		{4, 4, 9, 2, 10},
		{9, 7, 8, 3, 10},
		{8, 1, 9, 7, 5},
		{7, 10, 4, 11, 6},
	}))
}
