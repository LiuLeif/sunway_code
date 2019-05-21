// 2017-12-28 14:52
package main

import (
	"fmt"
	"sort"
)

func findMinArrowShots(points [][]int) int {
	if len(points) == 0 {
		return 0
	}
	sort.Slice(points, func(i, j int) bool {
		return points[i][1] < points[j][1]
	})
	end := points[0][1]
	count := 1
	for i := 1; i < len(points); i++ {
		if points[i][0] <= end {
			continue
		}
		count++
		end = points[i][1]
	}

	return count
}
func main() {
	fmt.Println(findMinArrowShots(
		[][]int{
			{10, 16},
			{2, 8},
			{1, 6},
			{7, 12},
		}))
}
