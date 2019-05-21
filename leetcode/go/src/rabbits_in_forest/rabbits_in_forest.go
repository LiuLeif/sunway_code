// 2018-02-24 15:52
package main

import (
	"fmt"
	"math"
)

func numRabbits(answers []int) int {
	m := map[int]int{}
	for _, ans := range answers {
		m[ans+1]++
	}
	ret := 0
	for x, y := range m {
		ret += int(math.Ceil(float64(y)/float64(x)) * float64(x))
	}
	return ret
}
func main() {
	fmt.Println(numRabbits([]int{1, 1, 2}))
	fmt.Println(numRabbits([]int{10, 10, 10}))
	fmt.Println(numRabbits([]int{}))
}
