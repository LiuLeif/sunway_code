// 2017-12-13 13:52
package main

import "fmt"

func judgeCircle(moves string) bool {
	x, y := 0, 0
	for i := 0; i < len(moves); i++ {
		switch moves[i] {
		case 'L':
			x--
		case 'U':
			y++
		case 'R':
			x++
		case 'D':
			y--
		}
	}
	return x == 0 && y == 0
}
func main() {
	fmt.Println(judgeCircle("LUR"))
}
