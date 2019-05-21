// 2018-12-14 11:48
package main

import (
	"fmt"
)

type Tuple struct {
	x, y int
	mask int
	step int
}

type Queue []Tuple

func (this *Queue) Poll() Tuple {
	tmp := (*this)[0]
	*this = (*this)[1:]
	return tmp
}

func (this *Queue) Offer(x Tuple) {
	*this = append(*this, x)
}

func (this *Queue) Empty() bool {
	return len(*this) == 0
}

func (this *Queue) Peek() Tuple {
	return (*this)[0]
}

func shortestPathAllKeys(grid []string) int {
	M, N := len(grid), len(grid[0])
	X, Y := 0, 0
	totalMask := 0
	for i := 0; i < M; i++ {
		for j := 0; j < N; j++ {
			if grid[i][j] == '@' {
				X, Y = i, j
			}
			if grid[i][j] >= 'a' && grid[i][j] <= 'z' {
				totalMask |= 1 << (grid[i][j] - 'a')
			}
		}
	}

	queue := Queue{Tuple{X, Y, 0, 0}}
	visited := map[[3]int]bool{}
	visited[[3]int{X, Y, 0}] = true

	for !queue.Empty() {
		top := queue.Poll()
		for _, d := range [][2]int{{0, 1}, {0, -1}, {1, 0}, {-1, 0}} {
			xx, yy := top.x+d[0], top.y+d[1]
			if xx < 0 || xx >= M || yy < 0 || yy >= N {
				continue
			}
			if grid[xx][yy] == '#' {
				continue
			}
			mask := top.mask
			if grid[xx][yy] >= 'a' && grid[xx][yy] <= 'z' {
				mask |= 1 << (grid[xx][yy] - 'a')
			}
			if mask == totalMask {
				return top.step + 1
			}
			if grid[xx][yy] >= 'A' && grid[xx][yy] <= 'Z' {
				if mask&(1<<(grid[xx][yy]-'A')) == 0 {
					// no key
					continue
				}
			}
			if !visited[[3]int{xx, yy, mask}] {
				visited[[3]int{xx, yy, mask}] = true
				queue.Offer(Tuple{xx, yy, mask, top.step + 1})
			}
		}
	}
	return -1
}

func main() {
	fmt.Println(shortestPathAllKeys([]string{
		"@...a",
		".###A",
		"b.BCc",
	}))
	// fmt.Println(shortestPathAllKeys([]string{
	// 	"@.a.#",
	// 	"###.#",
	// 	"b.A.B",
	// }))
}
