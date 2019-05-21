// 2018-11-23 09:59
package main

import (
	"fmt"
	"sort"
)

func intersectionSizeTwo(intervals [][]int) int {

	sort.Slice(intervals, func(i, j int) bool {
		if intervals[i][1] < intervals[j][1] {
			return true
		}
		if intervals[i][1] > intervals[j][1] {
			return false
		}
		return intervals[i][0] > intervals[j][0]
	})

	firstInter := intervals[0]
	p0, p1 := firstInter[len(firstInter)-1], firstInter[len(firstInter)-1]-1
	ret := 2
	for _, inter := range intervals[1:] {
		last := inter[len(inter)-1]
		if !(p0 >= inter[0] && p0 <= inter[len(inter)-1]) {
			p0 = last
			last--
			ret++
		}

		if !(p1 >= inter[0] && p1 <= inter[len(inter)-1]) {
			p1 = last
			ret++
		}
	}
	return ret
}

func main() {
	fmt.Println(intersectionSizeTwo([][]int{
		{1, 3}, {1, 4}, {2, 5}, {3, 5},
	}))
	fmt.Println(intersectionSizeTwo([][]int{
		{1, 2}, {2, 3}, {2, 4}, {4, 5},
	}))
	fmt.Println(intersectionSizeTwo([][]int{
		{2, 10}, {3, 7}, {3, 15}, {4, 11}, {6, 12}, {6, 16}, {7, 8}, {7, 11}, {7, 15}, {11, 12},
	}))
	fmt.Println(intersectionSizeTwo([][]int{
		{16, 18}, {11, 18}, {15, 23}, {1, 16}, {10, 16}, {6, 19}, {18, 20}, {7, 19}, {10, 11}, {11, 23}, {6, 7}, {23, 25}, {1, 3}, {7, 12}, {1, 13}, {23, 25}, {10, 22}, {23, 25}, {0, 19}, {0, 13}, {7, 12}, {14, 19}, {8, 17}, {7, 23}, {4, 24},
	}))

}
