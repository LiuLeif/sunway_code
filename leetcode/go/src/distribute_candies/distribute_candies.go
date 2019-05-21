// 2018-01-05 13:13
package main

import (
	"fmt"
	"math"
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

func distributeCandies(candies []int) int {
	m := map[int]int{}
	for i := 0; i < len(candies); i++ {
		m[candies[i]]++
	}
	return min(len(m), len(candies)/2)
}
func main() {
	fmt.Println(distributeCandies([]int{1, 1, 2, 2, 3, 3}))
	fmt.Println(distributeCandies([]int{1, 2}))
}
