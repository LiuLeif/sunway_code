// 2018-12-12 17:56
package main

import (
	"fmt"
	"math"
)

func min(nums ...int) int {
	ret := math.MaxInt32
	for _, n := range nums {
		if ret > n {
			ret = n
		}
	}
	return ret
}

func max(nums ...int) int {
	ret := math.MinInt32
	for _, n := range nums {
		if ret < n {
			ret = n
		}
	}
	return ret
}

func isRectangleCover(rectangles [][]int) bool {
	x1, y1, x2, y2 := math.MaxInt32, math.MaxInt32, math.MinInt32, math.MinInt32
	count := map[[2]int]int{}
	area := 0
	mark := func(x, y int) {
		count[[2]int{x, y}]++
		if count[[2]int{x, y}]%2 == 0 {
			delete(count, [2]int{x, y})
		}
	}
	for _, rect := range rectangles {
		x1 = min(x1, rect[0])
		y1 = min(y1, rect[1])
		x2 = max(x2, rect[2])
		y2 = max(y2, rect[3])
		mark(rect[0], rect[1])
		mark(rect[2], rect[3])
		mark(rect[0], rect[3])
		mark(rect[2], rect[1])
		area += (rect[3] - rect[1]) * (rect[2] - rect[0])
	}

	if len(count) != 4 {
		return false
	}
	if count[[2]int{x1, y1}] != 1 || count[[2]int{x2, y2}] != 1 || count[[2]int{x1, y2}] != 1 || count[[2]int{x2, y1}] != 1 {
		return false
	}
	return area == (y2-y1)*(x2-x1)
}

func main() {
	fmt.Println(isRectangleCover([][]int{
		{1, 1, 3, 3},
		{3, 1, 4, 2},
		{3, 2, 4, 4},
		{1, 3, 2, 4},
		{2, 3, 3, 4},
	}))
	fmt.Println(isRectangleCover([][]int{
		{0, 0, 1, 1}, {0, 0, 2, 1}, {1, 0, 2, 1}, {0, 2, 2, 3},
	}))
}
