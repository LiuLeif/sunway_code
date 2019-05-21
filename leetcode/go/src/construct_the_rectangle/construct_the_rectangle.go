// 2018-01-22 15:27
package main

import (
	"fmt"
	"math"
)

func constructRectangle(area int) []int {
	ret := []int{}
	start := int(math.Ceil(math.Sqrt(float64(area))))
	for i := start; i <= area; i++ {
		if area%i == 0 {
			return []int{i, area / i}
		}
	}
	return ret
}

func main() {
	// 6 2 3
	// 8 2 4
	// 9 3 3
	// 5
	fmt.Println(constructRectangle(12))
}
