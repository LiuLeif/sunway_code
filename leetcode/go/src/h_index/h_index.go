// 2017-12-11 16:36
package main

import "fmt"

func hIndex(citations []int) int {
	n := len(citations)
	buckets := make([]int, n+1)
	for _, v := range citations {
		if v > n {
			v = n
		}
		for j := 1; j < v+1; j++ {
			buckets[j]++
		}

	}
	for i := n; i > 0; i-- {
		if buckets[i] >= i {
			return i
		}
	}

	return 0
}
func main() {
	fmt.Println(hIndex([]int{1, 2, 3}))
}
