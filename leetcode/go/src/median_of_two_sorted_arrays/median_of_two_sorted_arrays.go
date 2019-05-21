// 2017-11-14 18:54
package main

import "fmt"

func doFind(nums1 []int, nums2 []int, n int) int {
	// 123 1 2
	fmt.Println(nums1, nums2, n)
	l1 := len(nums1)
	l2 := len(nums2)
	if l1 == 0 {
		return nums2[n-1]
	}
	if l2 == 0 {
		return nums1[n-1]
	}

	m1 := l1 / 2
	m2 := l2 / 2

	x := nums1[m1]
	y := nums2[m2]

	if m1+m2 < n-1 {
		if x > y {
			return doFind(nums1, nums2[m2+1:], n-m2-1)
		} else {
			return doFind(nums1[m1+1:], nums2, n-m1-1)
		}
	} else {
		if x > y {
			return doFind(nums1[:m1], nums2, n)
		} else {
			return doFind(nums1, nums2[:m2], n)
		}
	}
}

func findMedianSortedArrays(nums1 []int, nums2 []int) float64 {
	l := len(nums1) + len(nums2)
	if l%2 == 0 {
		return float64(doFind(nums1, nums2, l/2)+doFind(nums1, nums2, l/2+1)) / 2.0
	}
	return float64(doFind(nums1, nums2, l/2+1))
}

func main() {
	fmt.Println(findMedianSortedArrays([]int{1, 2}, []int{1, 1}))
}
