// 2018-03-01 15:12
package main

import (
	"fmt"
)

func findPoisonedDuration(timeSeries []int, duration int) int {
	if len(timeSeries) == 0 {
		return 0
	}
	ret := 0
	for i := 1; i < len(timeSeries); i++ {
		if timeSeries[i]-timeSeries[i-1] > duration {
			ret += duration
		} else {
			ret += timeSeries[i] - timeSeries[i-1]
		}
	}
	return ret + duration
}
func main() {
	fmt.Println(findPoisonedDuration([]int{1, 4}, 2))
	fmt.Println(findPoisonedDuration([]int{1}, 2))
}
