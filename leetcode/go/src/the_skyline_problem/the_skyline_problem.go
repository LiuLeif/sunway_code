// 2018-01-25 09:47
package main

import (
	"container/heap"
	"fmt"
	"math"
	"sort"
)

// 1. heap approach
type Heap []int

func (h Heap) Len() int {
	return len(h)
}

func (h Heap) Less(i, j int) bool {
	return h[i] > h[j]
}

func (h Heap) Swap(i, j int) {
	h[i], h[j] = h[j], h[i]
}

func (ph *Heap) Pop() interface{} {
	t := (*ph)[len(*ph)-1]
	*ph = (*ph)[:len(*ph)-1]
	return t
}

func (ph *Heap) Push(x interface{}) {
	*ph = append(*ph, x.(int))
}

type PriorityQueue struct {
	innerHeap *Heap
	count     map[int]int
}

func newPriorityQueue() *PriorityQueue {
	ret := &PriorityQueue{}
	ret.innerHeap = &Heap{}
	heap.Init(ret.innerHeap)
	ret.count = map[int]int{}
	return ret
}

func (this *PriorityQueue) Peek() int {
	for len(*(this.innerHeap)) != 0 {
		ret := (*(this.innerHeap))[0]
		if this.count[ret] != 0 {
			return ret
		}
		heap.Pop(this.innerHeap)
	}
	return 0
}

func (this *PriorityQueue) Remove(x int) {
	this.count[x]--
}

func (this *PriorityQueue) Push(x int) {
	this.count[x]++
	heap.Push(this.innerHeap, x)
}

const (
	ENTER = iota
	LEAVE
)

type Height struct {
	x, y    int
	variant int
}

func getSkyline(buildings [][]int) [][]int {
	ret := [][]int{}
	heights := []Height{}
	for _, building := range buildings {
		heights = append(heights, Height{building[0], building[2], ENTER})
		heights = append(heights, Height{building[1], building[2], LEAVE})
	}

	sort.Slice(heights, func(i, j int) bool {
		if heights[i].x < heights[j].x {
			return true
		}
		if heights[i].x > heights[j].x {
			return false
		}

		iVal, jVal := heights[i].y, heights[j].y
		if heights[i].variant == ENTER {
			iVal = -iVal
		}
		if heights[j].variant == ENTER {
			jVal = -jVal
		}
		return iVal < jVal
	})

	pq := newPriorityQueue()

	prevHeight := 0
	for _, height := range heights {
		if height.variant == ENTER {
			pq.Push(height.y)
		} else {
			pq.Remove(height.y)
		}

		top := pq.Peek()
		if prevHeight != top {
			ret = append(ret, []int{height.x, top})
			prevHeight = top
		}
	}

	return ret
}

// 2. interval tree approach

type Node struct {
	lo, hi      int
	val         int
	left, right *Node
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

func update(root *Node, f, t, val int) {
	if root.lo == root.hi {
		root.val = max(root.val, val)
		return
	}
	mid := (root.lo + root.hi) / 2
	if t <= mid {
		update(root.left, f, t, val)
	} else if f > mid {
		update(root.right, f, t, val)
	} else {
		update(root.left, f, mid, val)
		update(root.right, mid+1, t, val)
	}
}

func query(root *Node, index int) int {
	if root.lo == root.hi && root.lo == index {
		return root.val
	}

	mid := (root.lo + root.hi) / 2
	if index <= mid {
		return query(root.left, index)
	} else {
		return query(root.right, index)
	}
}

func buildTree(f, t int) *Node {
	// 0~n-1
	root := &Node{f, t, 0, nil, nil}
	if f == t {
		return root
	}
	mid := (f + t) / 2
	root.left = buildTree(f, mid)
	root.right = buildTree(mid+1, t)
	return root
}

func discrete(buildings [][]int) (map[int]int, map[int]int, int) {
	mark := map[int]bool{}
	for _, building := range buildings {
		x := building[0]
		y := building[1]
		mark[x] = true
		mark[y] = true
	}
	points := []int{}
	for k, _ := range mark {
		points = append(points, k)
	}
	sort.Ints(points)
	k := 0
	rm, m := map[int]int{}, map[int]int{}
	for _, p := range points {
		m[p] = k
		rm[k] = p
		k++
	}
	return m, rm, len(points)
}

func getSkyline2(buildings [][]int) [][]int {
	ret := [][]int{}

	m, rm, count := discrete(buildings)
	root := buildTree(0, count-1)
	for _, building := range buildings {
		update(root, m[building[0]], m[building[1]]-1, building[2])
	}

	prev := 0
	for i := 0; i < count; i++ {
		curr := query(root, i)
		if curr != prev {
			ret = append(ret, []int{rm[i], curr})
			prev = curr
		}
	}
	return ret
}
func main() {
	fmt.Println(getSkyline([][]int{
		{2, 9, 10},
		{3, 7, 15},
		{5, 12, 12},
		{15, 20, 10},
		{19, 24, 8},
	}))
	fmt.Println(getSkyline2([][]int{
		{2, 9, 10},
		{3, 7, 15},
		{5, 12, 12},
		{15, 20, 10},
		{19, 24, 8},
	}))

	fmt.Println(getSkyline([][]int{
		{0, 2, 3},
		{2, 5, 3},
	}))
	fmt.Println(getSkyline2([][]int{
		{0, 2, 3},
		{2, 5, 3},
	}))
}
