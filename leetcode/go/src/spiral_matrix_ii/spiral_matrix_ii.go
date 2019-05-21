// 2017-12-14 16:21
package main

import "fmt"

const (
	right = iota
	down
	left
	up
)

func generateMatrix(n int) [][]int {
	ret := make([][]int, n)
	for i := 0; i < n; i++ {
		ret[i] = make([]int, n)
	}
	k := 1
	direction := right
	x1, y1 := -1, -1
	x2, y2 := n, n
	x, y := 0, 0
	for {
		if x <= x1 || x >= x2 || y <= y1 || y >= y2 {
			break
		}
		ret[x][y] = k
		k++
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
	fmt.Println(generateMatrix(3))
}
