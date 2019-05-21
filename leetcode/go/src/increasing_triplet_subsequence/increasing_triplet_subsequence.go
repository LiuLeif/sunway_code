// 2017-11-14 18:54
package main

import (
	"fmt"
	"math"
)

func increasingTriplet(nums []int) bool {
	min, secMin := math.MaxInt32, math.MaxInt32
	for _, v := range nums {
		if v <= min {
			min = v
		} else if v > min && v <= secMin {
			secMin = v
		} else {
			return true
		}
	}
	return false
}
func main() {
	fmt.Println(increasingTriplet([]int{5, 4, 2, 2, 3}))
}
