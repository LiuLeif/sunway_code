// 2018-01-26 10:50
package main

import (
	"fmt"
	"math"
	"sort"
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

func max(nums ...int) int {
	ret := math.MinInt32
	for _, n := range nums {
		if ret < n {
			ret = n
		}
	}
	return ret
}

type Node struct {
	lo, hi      int
	val         int
	left, right *Node
}

func buildTree(f, t int) *Node {
	root := &Node{f, t, 0, nil, nil}
	if f == t {
		return root
	}
	mid := (f + t) / 2
	root.left = buildTree(f, mid)
	root.right = buildTree(mid+1, t)
	return root
}

func update(root *Node, f, t int) {
	if root.lo == root.hi {
		root.val++
		return
	}
	mid := (root.lo + root.hi) / 2
	if t <= mid {
		update(root.left, f, t)
	} else if f > mid {
		update(root.right, f, t)
	} else {
		update(root.left, f, mid)
		update(root.right, mid+1, t)
	}
	root.val = max(root.left.val, root.right.val)
}

func query(root *Node, f, t int) int {
	if root.lo == f && root.hi == t {
		return root.val
	}
	mid := (root.lo + root.hi) / 2
	if t <= mid {
		return query(root.left, f, t)
	} else if f > mid {
		return query(root.right, f, t)
	} else {
		return max(query(root.left, f, mid), query(root.right, mid+1, t))
	}
}

func discrete(intervals [][]int) (int, map[int]int) {
	counter := map[int]bool{}
	for _, interval := range intervals {
		counter[interval[0]] = true
		counter[interval[1]] = true
	}

	points := []int{}
	for t, _ := range counter {
		points = append(points, t)
	}
	sort.Ints(points)
	m := map[int]int{}
	for i, v := range points {
		m[v] = i
	}

	return len(points), m
}

func maxIntervalsOverlap(intervals [][]int) int {
	n, mapping := discrete(intervals)

	root := buildTree(0, n-1)
	for _, interval := range intervals {
		update(root, mapping[interval[0]], mapping[interval[1]])
	}
	return query(root, 0, n-1)
}

const (
	ENTER = iota
	LEAVE
)

type Event struct {
	x     int
	event int
}

func maxIntervalsOverlap2(intervals [][]int) int {
	events := []Event{}
	for _, interval := range intervals {
		events = append(events, Event{interval[0], ENTER})
		events = append(events, Event{interval[1] + 1, LEAVE})
	}
	sort.Slice(events, func(i, j int) bool {
		if events[i].x < events[j].x {
			return true
		}
		if events[i].x > events[j].x {
			return false
		}
		if events[i].event == LEAVE {
			return true
		}
		return false
	})

	ret, curr := 0, 0
	for _, event := range events {
		if event.event == ENTER {
			curr++
		} else {
			curr--
		}
		ret = max(ret, curr)
	}

	return ret
}

func main() {
	fmt.Println(maxIntervalsOverlap([][]int{
		{1, 2},
		{1, 3},
		{3, 5},
		{3, 8},
	}))
	fmt.Println(maxIntervalsOverlap2([][]int{
		{1, 2},
		{1, 3},
		{3, 5},
		{3, 8},
	}))
}
