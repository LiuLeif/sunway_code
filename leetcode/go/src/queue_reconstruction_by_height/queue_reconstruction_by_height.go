// 2018-01-15 10:00
package main

import (
	"fmt"
	"sort"
)

func reconstructQueue(people [][]int) [][]int {
	ret := [][]int{}
	sort.Slice(people, func(i, j int) bool {
		if people[i][0] < people[j][0] {
			return false
		}
		if people[i][0] > people[j][0] {
			return true
		}
		return people[i][1] < people[j][1]
	})
	for i := 0; i < len(people); i++ {
		insert := people[i][1]
		ret = append(ret[:insert], append([][]int{{people[i][0], people[i][1]}}, ret[insert:]...)...)
	}
	return ret
}
func main() {
	fmt.Println(reconstructQueue([][]int{
		{7, 0},
		{4, 4},
		{7, 1},
		{5, 0},
		{6, 1},
		{5, 2},
	}))
}
