// 2018-01-17 15:15
package main

import (
	"fmt"
	"math"
)

type Point struct {
	X int
	Y int
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

type Slope struct {
	i, j int
}

func gcd(x, y int) int {
	if x < y {
		x, y = y, x
	}
	for y != 0 {
		x, y = y, x%y
	}
	return x
}

func getSlope(p1, p2 Point) Slope {
	if p1.X == p2.X {
		return Slope{1, 0}
	}
	if p1.Y == p2.Y {
		return Slope{0, 1}
	}

	deltaX := p1.X - p2.X
	deltaY := p1.Y - p2.Y

	negative := false
	if deltaX*deltaY < 0 {
		negative = true
	}

	if deltaX < 0 {
		deltaX = -deltaX
	}
	if deltaY < 0 {
		deltaY = -deltaY
	}

	d := gcd(deltaX, deltaY)
	deltaX /= d
	deltaY /= d

	if negative {
		deltaX = -deltaX
	}
	return Slope{deltaX, deltaY}

}

func maxPoints(points []Point) int {
	if len(points) == 0 {
		return 0
	}
	ret := 0
	for i := 0; i < len(points); i++ {
		m := map[Slope]int{}
		dup := 1
		maxSlope := 0
		for j := 0; j < len(points); j++ {
			if i == j {
				continue
			}
			if points[i] == points[j] {
				dup++
				continue
			}
			slope := getSlope(points[i], points[j])
			m[slope]++
			maxSlope = max(maxSlope, m[slope])
		}
		ret = max(ret, maxSlope+dup)
	}

	return ret
}
func main() {
	fmt.Println(maxPoints([]Point{
		{0, 0},
		{0, 0},
		{0, 0},
		{1, 1},
		{2, 3},
		{2, 3},
	}))
}
