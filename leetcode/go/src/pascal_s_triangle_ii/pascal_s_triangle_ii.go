// 2017-11-27 14:25
package main

import "fmt"

func getRow(k int) []int {
	lastColumn := make([]int, k+1)
	thisColumn := make([]int, k+1)
	lastColumn[0] = 1

	for i := 1; i < k+1; i++ {
		// i+1 value
		for j := 0; j < i+1; j++ {
			x, y := 0, 0
			if j > 0 {
				x = lastColumn[j-1]
			}
			if j < i {
				y = lastColumn[j]
			}
			thisColumn[j] = x + y
		}

		lastColumn, thisColumn = thisColumn, lastColumn
	}
	return lastColumn
}

func main() {
	fmt.Println(getRow(3))
}
