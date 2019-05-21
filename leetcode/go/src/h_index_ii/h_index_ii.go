// 2017-12-11 16:36
package main

import "fmt"

func hIndex(citations []int) int {
	if len(citations) == 0 || citations[0] >= len(citations) {
		return len(citations)
	}
	lo, hi := 1, len(citations)-1
	for lo <= hi {
		mid := lo + (hi-lo)/2
		index := len(citations) - mid
		if citations[index] >= mid {
			lo = mid + 1
		} else {
			hi = mid - 1
		}
	}
	return hi
}
func main() {
	fmt.Println(hIndex([]int{1, 2, 2, 2, 3, 3, 4, 5, 5, 5}))
}
