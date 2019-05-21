// 2018-10-16 09:42
package main

import "fmt"

func projectionArea(grid [][]int) int {
	max := func(v ...int) int {
		ret := v[0]
		for i := 0; i < len(v); i++ {
			if v[i] > ret {
				ret = v[i]
			}
		}
		return ret
	}

	sum := func(v ...int) int {
		ret := 0
		for i := 0; i < len(v); i++ {
			ret += v[i]
		}
		return ret
	}

	rowMax := make([]int, len(grid))
	colMax := make([]int, len(grid[0]))
	total := 0
	for i := 0; i < len(grid); i++ {
		for j := 0; j < len(grid[0]); j++ {
			if grid[i][j] != 0 {
				total++
			}
			colMax[j] = max(colMax[j], grid[i][j])
			rowMax[i] = max(rowMax[i], grid[i][j])
		}
	}

	total += sum(rowMax...) + sum(colMax...)

	return total
}
func main() {
	fmt.Println(projectionArea([][]int{{1, 0}, {0, 2}}))
}
