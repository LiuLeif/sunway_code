// 2018-10-12 17:27
package main

import "fmt"

func isRectangleOverlap(rec1 []int, rec2 []int) bool {
	min := func(a, b int) int {
		if a < b {
			return a
		}
		return b
	}

	max := func(a, b int) int {
		if a > b {
			return a
		}
		return b
	}

	width := min(rec2[2], rec1[2]) - max(rec2[0], rec1[0])
	height := min(rec2[3], rec1[3]) - max(rec2[1], rec1[1])

	if width > 0 && height > 0 {
		return true
	}
	return false
}
func main() {
	fmt.Println(isRectangleOverlap([]int{0, 0, 2, 2}, []int{1, 1, 3, 3}))
	fmt.Println(isRectangleOverlap([]int{0, 0, 1, 1}, []int{1, 0, 2, 1}))
}
