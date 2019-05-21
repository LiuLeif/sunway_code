// 2018-01-03 14:30
package main

import (
	"fmt"
	"sort"
	"strconv"
)

type Pair struct {
	x, y int
}

func findRelativeRanks(nums []int) []string {
	pairs := []Pair{}
	for i := 0; i < len(nums); i++ {
		pairs = append(pairs, Pair{nums[i], i})
	}
	sort.Slice(pairs, func(i, j int) bool {
		return pairs[i].x > pairs[j].x
	})
	fmt.Println(pairs)
	ret := make([]string, len(nums))
	for i := 0; i < len(pairs); i++ {
		switch i {
		case 0:
			ret[pairs[i].y] = "Gold Medal"
		case 1:
			ret[pairs[i].y] = "Silver Medal"
		case 2:
			ret[pairs[i].y] = "Bronze Medal"
		default:
			ret[pairs[i].y] = strconv.Itoa(i + 1)
		}
	}

	return ret
}
func main() {
	fmt.Println(findRelativeRanks([]int{3, 5}))
}
