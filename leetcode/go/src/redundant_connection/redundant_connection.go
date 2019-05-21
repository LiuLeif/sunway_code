// 2018-01-30 10:45
package main

import "fmt"

type UFS struct {
	parent map[int]int
	height map[int]int
}

func NewUFS() *UFS {
	return &UFS{map[int]int{}, map[int]int{}}
}

func (this *UFS) Union(a, b int) {
	a = this.Find(a)
	b = this.Find(b)
	if this.height[a] > this.height[b] {
		this.parent[b] = a
	} else if this.height[a] < this.height[b] {
		this.parent[a] = b
	} else {
		this.parent[a] = b
		this.height[b]++
	}
}

func (this *UFS) Find(a int) int {
	if this.parent[a] == 0 {
		this.parent[a] = a
	}
	for this.parent[a] != a {
		a = this.parent[a]
	}
	return a
}

func findRedundantConnection(edges [][]int) []int {
	u := NewUFS()
	for _, edge := range edges {
		f, t := edge[0], edge[1]
		if u.Find(f) == u.Find(t) {
			return edge
		}
		u.Union(f, t)
	}
	return []int{}
}
func main() {
	fmt.Println(findRedundantConnection([][]int{
		{1, 2},
		{2, 3},
		{3, 4},
		{1, 4},
		{1, 5},
	}))
}
