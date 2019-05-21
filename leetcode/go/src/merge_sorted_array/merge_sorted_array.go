// 2017-11-14 18:54
package main

import (
	"fmt"
	"math"
)

func merge(nums1 []int, m int, nums2 []int, n int) {
	copy(nums1[n:], nums1)
	i := n // to m+n
	j := 0 // to n
	k := 0
	for k < m+n {
		x := math.MaxInt32
		if i < m+n {
			x = nums1[i]
		}
		y := math.MaxInt32
		if j < n {
			y = nums2[j]
		}
		if x <= y {
			nums1[k] = x
			i++
		} else {
			nums1[k] = y
			j++
		}
		k += 1
	}
	fmt.Println(nums1)
}
func main() {
	nums1 := []int{1, 1, 0, 0, 0, 0}
	nums2 := []int{}
	merge(nums1, 2, nums2, 0)
}
