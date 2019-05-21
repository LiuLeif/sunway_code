// 2017-11-14 18:54
package main

import (
	"fmt"
	"math"
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

func candy(ratings []int) int {
	forward := make([]int, len(ratings))
	backward := make([]int, len(ratings))

	for i := 0; i < len(ratings); i++ {
		if i == 0 {
			forward[i] = 1
		} else {
			if ratings[i] > ratings[i-1] {
				forward[i] = forward[i-1] + 1
			} else {
				forward[i] = 1
			}
		}
	}
	fmt.Println(forward)

	for i := len(ratings) - 1; i >= 0; i-- {
		if i == len(ratings)-1 {
			backward[i] = 1
		} else {
			if ratings[i] > ratings[i+1] {
				backward[i] = backward[i+1] + 1
			} else {
				backward[i] = 1
			}
		}
	}

	ret := 0
	for i := 0; i < len(ratings); i++ {
		ret += max(forward[i], backward[i])
	}

	return ret
}
func main() {
	// 1 2 3 3 2
	fmt.Println(candy([]int{}))
}
