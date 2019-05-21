// 2018-01-29 12:40
package main

import "fmt"

func imageSmoother(M [][]int) [][]int {
	ret := make([][]int, len(M))
	for i := 0; i < len(M); i++ {
		ret[i] = make([]int, len(M[0]))
	}
	get := func(i, j int) (int, bool) {
		if i < 0 || j < 0 || i > len(M)-1 || j > len(M[0])-1 {
			return 0, false
		} else {
			return M[i][j], true
		}
	}
	for i := 0; i < len(M); i++ {
		for j := 0; j < len(M[0]); j++ {
			n := 0
			sum := 0
			for a := -1; a < 2; a++ {
				for b := -1; b < 2; b++ {
					if x, ok := get(i+a, j+b); ok {
						n++
						sum += x
					}
				}
			}
			ret[i][j] = sum / n
		}
	}
	return ret
}
func main() {
	fmt.Println(imageSmoother([][]int{
		{1, 1, 1},
		{1, 0, 1},
		{1, 1, 1},
	}))
}
