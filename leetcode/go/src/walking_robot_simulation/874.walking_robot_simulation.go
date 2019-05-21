// 2018-11-06 16:48
package main

import (
	"fmt"
	"math"
)

type Pair struct {
	a, b int
}

func robotSim(commands []int, obstacles [][]int) int {
	max := func(nums ...int) int {
		ret := math.MinInt32
		for _, n := range nums {
			if ret < n {
				ret = n
			}
		}
		return ret
	}

	obstaclesSet := map[Pair]bool{}
	for i := 0; i < len(obstacles); i++ {
		obstaclesSet[Pair{obstacles[i][0], obstacles[i][1]}] = true
	}
	directions := []Pair{{0, 1}, {-1, 0}, {0, -1}, {1, 0}}
	d := 0
	x, y := 0, 0
	ret := 0
	for _, command := range commands {
		if command == -2 {
			d = (d + 1) % 4
		} else if command == -1 {
			d = (d + 3) % 4
		} else {
			for command > 0 {
				if obstaclesSet[Pair{x + directions[d].a, y + directions[d].b}] {
					break
				}
				x += directions[d].a
				y += directions[d].b
				command--
				ret = max(ret, x*x+y*y)
			}
		}
	}
	return ret
}

func main() {
	fmt.Println(robotSim(
		[]int{
			4, -1, 4, -2, 4,
		},
		[][]int{
			{2, 4},
		}))
}
