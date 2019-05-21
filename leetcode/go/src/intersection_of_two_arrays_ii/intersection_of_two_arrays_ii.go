// 2017-12-28 14:52
package main

import (
	"fmt"
	"sort"
)

func intersect(nums1 []int, nums2 []int) []int {
	ret := []int{}
	sort.Ints(nums1)
	sort.Ints(nums2)
	a, b := 0, 0
	for a < len(nums1) && b < len(nums2) {
		if nums1[a] == nums2[b] {
			ret = append(ret, nums1[a])
			a++
			b++
		} else if nums1[a] > nums2[b] {
			b++
		} else {
			a++
		}
	}

	return ret
}
func main() {
	fmt.Println(intersect([]int{1, 2, 2, 1}, []int{1, 1}))
}
