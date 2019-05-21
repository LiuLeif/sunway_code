// 2018-10-12 16:07
package main

import "fmt"

func maxDistToClosest(seats []int) int {
	seats = append(seats, 1)
	ret := 0
	prev := -1
	max := func(a, b int) int {
		if a > b {
			return a
		}
		return b
	}

	for i := 0; i < len(seats); i++ {
		if seats[i] == 0 {
			continue
		}
		if prev == -1 {
			ret = max(ret, i)
		} else if i == len(seats)-1 {
			ret = max(ret, (i - prev - 1))
		} else {
			ret = max(ret, (i-prev)/2)
		}
		prev = i
	}
	return ret
}
func main() {
	fmt.Println(maxDistToClosest([]int{1, 0}))
}
