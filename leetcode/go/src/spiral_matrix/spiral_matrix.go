// 2017-12-14 16:21
package main

import "fmt"

const (
	right = iota
	down
	left
	up
)

func spiralOrder(matrix [][]int) []int {
	ret := []int{}
	if len(matrix) == 0 {
		return ret
	}
	direction := right
	x1, y1 := -1, -1
	x2, y2 := len(matrix), len(matrix[0])
	x, y := 0, 0
	for {
		if x <= x1 || x >= x2 || y <= y1 || y >= y2 {
			break
		}
		ret = append(ret, matrix[x][y])
		switch direction {
		case right:
			y++
			if y >= y2 {
				direction = down
				y--
				x++
				x1++
			}
		case down:
			x++
			if x >= x2 {
				direction = left
				x--
				y--
				y2--
			}
		case left:
			y--
			if y <= y1 {
				direction = up
				y++
				x--
				x2--
			}
		case up:
			x--
			if x <= x1 {
				direction = right
				x++
				y++
				y1++
			}
		}
	}
	return ret
}
func main() {
	matrix := [][]int{
		{1, 2, 3},
		{4, 5, 6},
	}
	fmt.Println(spiralOrder(matrix))
}
