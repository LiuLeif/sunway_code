// 2018-12-10 17:40
package main

import (
	"fmt"
	"math/rand"
)

type Solution struct {
	rects    [][]int
	areaSum  []int
	totalSum int
}

func area(rect []int) int {
	return (rect[2] - rect[0] + 1) * (rect[3] - rect[1] + 1)
}

func Constructor(rects [][]int) Solution {
	ret := Solution{}
	ret.rects = rects
	ret.areaSum = []int{}
	for _, rect := range rects {
		ret.totalSum += area(rect)
		ret.areaSum = append(ret.areaSum, ret.totalSum)

	}
	return ret
}

func firstLarger(x int, v []int) int {
	lo, hi := 0, len(v)-1
	for lo <= hi {
		mid := (lo + hi) / 2
		if v[mid] <= x {
			lo = mid + 1
		} else {
			hi = mid - 1
		}
	}
	return lo
}

func (this *Solution) Pick() []int {
	// 4,7
	// x:0~6 = 4
	x := rand.Intn(this.totalSum)
	n := firstLarger(x, this.areaSum)
	offset := x
	if n != 0 {
		offset -= this.areaSum[n-1]
	}
	rect := this.rects[n]
	w := rect[2] - rect[0] + 1
	row := offset / w
	col := offset - row*w
	return []int{rect[0] + col, rect[1] + row}
}

func main() {
	//                                   +-----------+
	//                                   |           |
	//                                   |           |
	//                +-------------+    |           |
	//                |             |    +-----------+
	//                |             |
	//                |             |
	//                |             |
	//                |             |
	//                +-------------+
	obj := Constructor([][]int{{-2, -2, -1, -1}, {1, 0, 3, 0}})
	for i := 0; i < 10; i++ {
		fmt.Println(obj.Pick())
	}
}
