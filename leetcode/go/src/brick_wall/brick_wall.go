// 2017-12-20 14:33
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

func leastBricks(wall [][]int) int {
	m := make(map[int]int)
	for i := 0; i < len(wall); i++ {
		sum := 0
		for j := 0; j < len(wall[i])-1; j++ {
			sum += wall[i][j]
			m[sum] -= 1
		}
	}

	ret := 0
	for _, v := range m {
		ret = min(ret, v)
	}
	ret += len(wall)
	return ret
}

func main() {
	wall := [][]int{
		{1, 2, 3},
		{2, 1, 3},
	}
	fmt.Println(leastBricks(wall))
}
