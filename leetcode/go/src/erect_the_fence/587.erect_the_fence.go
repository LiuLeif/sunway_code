// 2018-12-07 18:26
package main

import (
	"fmt"
	"sort"
)

// graham scan
type Point struct {
	X int
	Y int
}

//   x
// x o x o x
//       x
//     x
//
// ----
//
// o o x
//     x
//
//     x

func toVec(v1, v2 Point) Point {
	return Point{v2.X - v1.X, v2.Y - v1.Y}
}

func det(v1, v2 Point) int {
	return v1.X*v2.Y - v2.X*v1.Y
}

func outerTrees(points []Point) []Point {
	part1 := []Point{}
	N := len(points)
	sort.Slice(points, func(i, j int) bool {
		if points[i].Y != points[j].Y {
			return points[i].Y < points[j].Y
		}
		return points[i].X < points[j].X
	})
	for i := 0; i < N; i++ {
		for len(part1) > 1 && det(toVec(part1[len(part1)-1], part1[len(part1)-2]), toVec(points[i], part1[len(part1)-1])) < 0 {
			part1 = part1[:len(part1)-1]
		}
		part1 = append(part1, points[i])
	}

	part2 := []Point{}
	for i := N - 1; i >= 0; i-- {
		for len(part2) > 1 && det(toVec(part2[len(part2)-1], part2[len(part2)-2]), toVec(points[i], part2[len(part2)-1])) < 0 {
			part2 = part2[:len(part2)-1]
		}
		part2 = append(part2, points[i])
	}
	m := map[Point]bool{}
	for _, p := range part1 {
		m[p] = true
	}
	for _, p := range part2 {
		m[p] = true
	}
	ret := []Point{}
	for k := range m {
		ret = append(ret, k)
	}
	return ret
}

func main() {
	fmt.Println(outerTrees([]Point{{1, 1}, {2, 2}, {2, 0}, {2, 4}, {3, 3}, {4, 2}}))
	fmt.Println(outerTrees([]Point{{1, 2}, {2, 2}, {4, 2}}))
	fmt.Println(outerTrees([]Point{{1, 2}}))
}
