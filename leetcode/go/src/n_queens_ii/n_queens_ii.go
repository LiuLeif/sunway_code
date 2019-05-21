// 2017-12-21 15:30
package main

import (
	"fmt"
	"math"
)

type Pair struct {
	i, t int
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

func isValid(mark map[Pair]bool, x, y int) bool {
	// - mark
	if mark[Pair{-x, 0}] {
		return false
	}
	// | mark
	if mark[Pair{y, 1}] {
		return false
	}
	// \ mark
	if mark[Pair{y - x, 2}] {
		return false
	}
	// / mark
	if mark[Pair{y + x, 3}] {
		return false
	}
	return true
}

func setMark(mark map[Pair]bool, x, y int) {
	// - mark
	mark[Pair{-x, 0}] = true
	mark[Pair{y, 1}] = true
	mark[Pair{y - x, 2}] = true
	mark[Pair{y + x, 3}] = true
}

func clearMark(mark map[Pair]bool, x, y int) {
	mark[Pair{-x, 0}] = false
	mark[Pair{y, 1}] = false
	mark[Pair{y - x, 2}] = false
	mark[Pair{y + x, 3}] = false
}

func helper(mark map[Pair]bool, n int, current int) int {
	if current == n {
		return 1
	}
	ret := 0
	for i := 0; i < n; i++ {
		if isValid(mark, current, i) {
			setMark(mark, current, i)
			ret += helper(mark, n, current+1)
			clearMark(mark, current, i)
		}
	}
	return ret
}

func totalNQueens(n int) int {
	mark := make(map[Pair]bool)
	return helper(mark, n, 0)
}

func main() {
	fmt.Println(totalNQueens(8))
}
