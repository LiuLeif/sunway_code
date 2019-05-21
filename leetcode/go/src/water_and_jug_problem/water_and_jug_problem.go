// 2017-12-25 17:35
package main

import "fmt"

type Pair struct {
	x, y int
}

func canMeasureWater(x int, y int, z int) bool {
	if z > x+y {
		return false
	}
	queue := []Pair{Pair{0, 0}}
	visited := map[Pair]bool{}
	for len(queue) != 0 {
		top := queue[0]
		queue = queue[1:]
		if top.x+top.y == z {
			return true
		}
		// drop x
		p := Pair{0, top.y}
		if !visited[p] {
			queue = append(queue, p)
			visited[p] = true
		}
		// drop x to y
		if y-top.y >= top.x {
			p.x = 0
			p.y = top.y + top.x
		} else {
			p.x = top.x + top.y - y
			p.y = y
		}
		if !visited[p] {
			queue = append(queue, p)
			visited[p] = true
		}
		// drop y
		p.x = top.x
		p.y = 0
		if !visited[p] {
			queue = append(queue, p)
			visited[p] = true
		}
		// drop y to x
		if x-top.x >= top.y {
			p.x = top.x + top.y
			p.y = 0
		} else {
			p.x = x
			p.y = top.y + top.x - x
		}
		if !visited[p] {
			queue = append(queue, p)
			visited[p] = true
		}
		// fill x
		p.x = x
		p.y = top.y
		if !visited[p] {
			queue = append(queue, p)
			visited[p] = true
		}
		// fill y
		p.x = top.x
		p.y = y
		if !visited[p] {
			queue = append(queue, p)
			visited[p] = true
		}
	}
	return false
}
func main() {
	fmt.Println(canMeasureWater(1, 2, 3))
}
