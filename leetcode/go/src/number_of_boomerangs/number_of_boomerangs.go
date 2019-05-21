// 2018-01-22 13:37
package main

import "fmt"

func distance(p1, p2 []int) int {
	dx := p1[0] - p2[0]
	dy := p1[1] - p2[1]
	return dx*dx + dy*dy
}

func numberOfBoomerangs(points [][]int) int {
	ret := 0
	for i := 0; i < len(points); i++ {
		m := map[int]int{}
		for j := 0; j < len(points); j++ {
			if i == j {
				continue
			}
			d := distance(points[i], points[j])
			m[d]++
		}
		for _, v := range m {
			ret += v * (v - 1)
		}
	}
	return ret
}
func main() {
	//    e
	// a  b  c
	//    d
	fmt.Println(numberOfBoomerangs([][]int{
		{0, 0},
		{1, 0},
		{2, 0},
	}))
}
