// 2017-12-11 16:36
package main

import "fmt"

const NUM_COLORS int = 3

func sortColors(nums []int) {
	buckets := make([]int, NUM_COLORS)
	for _, color := range nums {
		buckets[color]++
	}

	index := 0
	for color := 0; color < NUM_COLORS; color++ {
		for j := 0; j < buckets[color]; j++ {
			nums[index] = color
			index++
		}
	}
}
func main() {
	colors := []int{0, 1, 0, 2, 1}
	sortColors(colors)
	fmt.Println(colors)
}
