// 2017-12-27 17:05
package main

import (
	"fmt"
	"sort"
)

func wiggleSort(nums []int) {
	// 1 1 1 4 5 6

	sort.Ints(nums)
	smaller := append([]int{}, nums[:(len(nums)+1)/2]...)
	for i := 0; i < len(smaller)/2; i++ {
		smaller[i], smaller[len(smaller)-i-1] = smaller[len(smaller)-i-1], smaller[i]
	}

	larger := append([]int{}, nums[(len(nums)+1)/2:]...)
	for i := 0; i < len(larger)/2; i++ {
		larger[i], larger[len(larger)-i-1] = larger[len(larger)-i-1], larger[i]
	}

	si, li := 0, 0
	for i := 0; i < len(nums); i++ {
		if i%2 == 0 {
			nums[i] = smaller[si]
			si++
		} else {
			nums[i] = larger[li]
			li++
		}
	}
}
func main() {
	// 4 5 5 6
	// 5 4 | 6 5
	nums := []int{1, 2, 1}
	wiggleSort(nums)
	fmt.Println(nums)
}
