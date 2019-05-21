// 2018-10-29 09:05
package main

import "fmt"

func numSubarraysWithSum(A []int, S int) int {
	// 0 2 4 5
	counter := []int{}
	current := 0
	for i := 0; i < len(A); i++ {
		if A[i] == 1 {
			counter = append(counter, current)
			current = 0
		} else {
			current++
		}
	}
	counter = append(counter, current)

	ret := 0
	if S == 0 {
		for _, n := range counter {
			ret += (1 + n) * n / 2
		}
		return ret
	}

	for i := 0; i < len(counter)-S; i++ {
		ret += (counter[i] + 1) * (counter[S+i] + 1)
	}
	return ret
}

func main() {
	fmt.Println(numSubarraysWithSum([]int{0, 0, 0, 0, 0}, 0))
}
