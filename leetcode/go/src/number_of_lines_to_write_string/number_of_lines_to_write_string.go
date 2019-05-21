// 2018-04-13 10:59
package main

import "fmt"

func numberOfLines(widths []int, S string) []int {
	line := 0
	currentWidth := 0
	for _, c := range S {
		width := widths[c-'a']
		if currentWidth+width > 100 {
			line++
			currentWidth = width
		} else {
			currentWidth += width
		}
	}
	ret := [2]int{line + 1, currentWidth}
	return ret[:]
}
func main() {
	fmt.Println(numberOfLines([]int{4, 10, 10, 10, 10, 10, 10, 10, 10, 10, 10, 10, 10, 10, 10, 10, 10, 10, 10, 10, 10, 10, 10, 10, 10, 10}, "bbbcccdddaaa"))
}
