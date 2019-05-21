// 2018-01-10 10:05
package main

import (
	"fmt"
	"sort"
)

func intersection(nums1 []int, nums2 []int) []int {
	ret := []int{}
	sort.Ints(nums1)
	sort.Ints(nums2)
	i, j := 0, 0
	for i < len(nums1) && j < len(nums2) {
		if nums1[i] == nums2[j] {
			if len(ret) == 0 || ret[len(ret)-1] != nums1[i] {
				ret = append(ret, nums1[i])
			}
			i++
			j++
		} else if nums1[i] > nums2[j] {
			j++
		} else {
			i++
		}
	}
	return ret
}
func main() {
	fmt.Println(intersection([]int{1, 2, 2, 1}, []int{1, 2, 2}))
}
