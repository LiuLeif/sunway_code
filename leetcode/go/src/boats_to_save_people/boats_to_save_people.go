// 2018-10-19 10:01
package main

import "fmt"
import "sort"

func numRescueBoats(people []int, limit int) int {
	sort.Ints(people)
	// 1 2 2 3
	lo, hi := 0, len(people)-1
	ret := 0
	for lo <= hi {
		if people[lo]+people[hi] <= limit {
			ret++
			lo++
			hi--
		} else {
			ret++
			hi--
		}
	}
	return ret
}
func main() {
	fmt.Println(numRescueBoats([]int{3, 5, 3, 4}, 5))
}
