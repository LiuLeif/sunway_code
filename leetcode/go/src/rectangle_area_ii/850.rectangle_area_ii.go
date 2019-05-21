// 2018-11-22 09:26
package main

import (
	"fmt"
	"math"
)

//        +------------+
//        |            |
//        |            |
// +------+-----+      |
// |      |     |   +--+-----+
// |      |     |   |  |     |
// |      +-----+---+--+     |
// |            |   |        |
// |            |   +--------+
// +------------+

var MOD = int(math.Pow(10, 9)) + 7

func max(nums ...int) int {
	ret := math.MinInt32
	for _, n := range nums {
		if ret < n {
			ret = n
		}
	}
	return ret
}

func min(nums ...int) int {
	ret := math.MaxInt32
	for _, n := range nums {
		if ret > n {
			ret = n
		}
	}
	return ret
}

func rectangleArea(rectangles [][]int) int {
	rectList := [][]int{}
	var add func(rect []int, index int)
	add = func(rect []int, index int) {
		if index == len(rectList) {
			rectList = append(rectList, rect)
			return
		}
		r := rectList[index]
		// no overlaping
		if r[0] >= rect[2] || r[2] <= rect[0] || r[1] >= rect[3] || r[3] <= rect[1] {
			add(rect, index+1)
			return
		}
		if rect[0] < r[0] {
			add([]int{rect[0], rect[1], r[0], rect[3]}, index+1)
		}
		if rect[2] > r[2] {
			add([]int{r[2], rect[1], rect[2], rect[3]}, index+1)
		}
		if rect[1] < r[1] {
			add([]int{max(rect[0], r[0]), rect[1], min(rect[2], r[2]), r[1]}, index+1)
		}
		if rect[3] > r[3] {
			add([]int{max(rect[0], r[0]), r[3], min(rect[2], r[2]), rect[3]}, index+1)
		}
	}
	for _, rec := range rectangles {
		add(rec, 0)
	}

	area := func(rect []int) int {
		return (rect[2] - rect[0]) * (rect[3] - rect[1]) % MOD
	}
	ret := 0
	for _, rec := range rectList {
		ret += area(rec)
		ret %= MOD
	}
	return ret
}

func main() {
	fmt.Println(rectangleArea([][]int{
		{0, 0, 2, 2}, {1, 0, 2, 3}, {1, 0, 3, 1},
	}))
	fmt.Println(rectangleArea([][]int{
		{25, 20, 70, 27}, {68, 80, 79, 100}, {37, 41, 66, 76},
	}))
}
