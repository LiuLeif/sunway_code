// 2018-01-09 13:53
package main

import (
	"fmt"
	"math"
	"sort"
)

func max(nums ...int) int {
	ret := math.MinInt32
	for _, n := range nums {
		if ret < n {
			ret = n
		}
	}
	return ret
}

func distance(x, y int) int {
	ret := x - y
	if ret < 0 {
		ret = -ret
	}
	return ret
}

func findRadius(houses []int, heaters []int) int {
	sort.Ints(houses)
	sort.Ints(heaters)

	heaterIndex := 0
	ret := 0
	for i := 0; i < len(houses); i++ {
		house := houses[i]
		for heaterIndex != len(heaters)-1 &&
			distance(house, heaters[heaterIndex]) >= distance(house, heaters[heaterIndex+1]) {

			heaterIndex++
		}
		ret = max(ret, distance(heaters[heaterIndex], house))
	}
	return ret
}
func main() {
	// 1 2 3 4   6
	//   ^   ^
	fmt.Println(findRadius([]int{1, 999}, []int{499, 500, 501}))
}
