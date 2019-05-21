// 2018-01-02 13:51
package main

import "fmt"

func generate(numRows int) [][]int {
	ret := make([][]int, numRows)
	if numRows == 0 {
		return ret
	}

	ret[0] = []int{1}
	for i := 1; i < numRows; i++ {
		ret[i] = make([]int, i+1)
		for j := 0; j < i+1; j++ {
			left := 0
			if j != 0 {
				left = ret[i-1][j-1]
			}
			right := 0
			if j != i {
				right = ret[i-1][j]
			}
			ret[i][j] = left + right
		}
	}
	return ret
}
func main() {
	fmt.Println(generate(3))
}
