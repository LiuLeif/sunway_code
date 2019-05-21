// 2018-01-04 17:04
package main

import (
	"fmt"
	"sort"
)

func distance(p1, p2 []int) int {
	x := p1[0] - p2[0]
	y := p1[1] - p2[1]
	return x*x + y*y
}

func validSquare(p1 []int, p2 []int, p3 []int, p4 []int) bool {
	points := [][]int{p1, p2, p3, p4}
	sort.Slice(points, func(i, j int) bool {
		if points[i][0] < points[j][0] {
			return true
		} else if points[i][0] > points[j][0] {
			return false
		} else {
			return points[i][1] < points[j][1]
		}
	})

	return distance(points[0], points[3]) == distance(points[1], points[2]) &&
		distance(points[0], points[2]) == distance(points[1], points[3]) &&
		distance(points[0], points[1]) == distance(points[2], points[3]) &&
		distance(points[0], points[1]) == distance(points[0], points[2]) &&
		distance(points[0], points[1]) != 0

}
func main() {
	fmt.Println(validSquare([]int{0, 0}, []int{1, 1}, []int{1, 0}, []int{0, 1}))
	fmt.Println(validSquare([]int{0, 0}, []int{5, 0}, []int{5, 4}, []int{0, 4}))
}
