// 2018-04-08 15:31
package main

import (
	"fmt"
	"math"
)

func largestTriangleArea(points [][]int) float64 {
	area := func(i, j, k int) float64 {
		// |1   1   1 |
		// |xi  xj  xk|
		// |yi  yj  yk|
		xi, yi, xj, yj, xk, yk := points[i][0], points[i][1],
			points[j][0], points[j][1], points[k][0], points[k][1]
		// shoelace formula
		return 0.5 * math.Abs(float64(xi*yj+xj*yk+xk*yi-xj*yi-xk*yj-xi*yk))
	}
	ret := 0.0
	for i := 0; i < len(points); i++ {
		for j := 0; j < len(points); j++ {
			for k := 0; k < len(points); k++ {
				if i == j || j == k || i == k {
					continue
				}
				ret = math.Max(ret, area(i, j, k))
			}
		}
	}
	return ret
}
func main() {
	fmt.Println(largestTriangleArea([][]int{
		{0, 0},
		{0, 1},
		{1, 0},
		{0, 2},
		{2, 0},
	}))
}
