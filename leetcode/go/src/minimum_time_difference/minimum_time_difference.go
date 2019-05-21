// 2018-02-24 10:44
package main

import (
	"fmt"
	"math"
	"sort"
	"strconv"
	"strings"
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

func timeToTick(time string) int {
	tmp := strings.Split(time, ":")
	hour, _ := strconv.Atoi(tmp[0])
	min, _ := strconv.Atoi(tmp[1])
	return hour*60 + min
}

func abs(x int) int {
	if x > 0 {
		return x
	}
	return -x
}

func findMinDifference(timePoints []string) int {
	ticks := make([]int, len(timePoints))
	for i := 0; i < len(timePoints); i++ {
		ticks[i] = timeToTick(timePoints[i])
	}

	sort.Ints(ticks)
	ticks = append(ticks, 24*60+ticks[0])

	ret := math.MaxInt32
	for i := 1; i < len(ticks); i++ {
		ret = min(ret, abs(ticks[i]-ticks[i-1]))
	}

	return ret
}
func main() {
	fmt.Println(findMinDifference([]string{"23:59", "00:01"}))
}
