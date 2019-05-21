// 2018-01-18 15:21
package main

import (
	"fmt"
	"math"
)

func min(nums ...int) int {
	ret := math.MaxInt32
	for _, n := range nums {
		if ret > n {
			ret = n
		}
	}
	return ret
}

func getMaxArray(nums []int, k int) []int {
	// 1 2 3 4, k=2
	drop := len(nums) - k
	ret := []int{}
	for i := 0; i < len(nums); i++ {
		for len(ret) != 0 && drop != 0 && nums[i] > ret[len(ret)-1] {
			ret = ret[:len(ret)-1]
			drop--
		}
		ret = append(ret, nums[i])
	}
	return ret[:k]
}

func maxArray(a, b []int) []int {
	if greater(a, b) {
		return a
	} else {
		return b
	}
}

func greater(a, b []int) bool {
	for i := 0; i < min(len(a), len(b)); i++ {
		if a[i] > b[i] {
			return true
		}
		if a[i] < b[i] {
			return false
		}
	}
	return len(a) > len(b)
}

func mergeArray(a, b []int) []int {
	// 1,2,3,4 | 3 4 5
	ret := make([]int, len(a)+len(b))
	i := 0
	j := 0
	for i < len(a) || j < len(b) {
		if i == len(a) {
			ret[i+j] = b[j]
			j++
		} else if j == len(b) {
			ret[i+j] = a[i]
			i++
		} else {
			if a[i] > b[j] {
				ret[i+j] = a[i]
				i++
			} else if a[i] < b[j] {
				ret[i+j] = b[j]
				j++
			} else {
				if greater(a[i:], b[j:]) {
					ret[i+j] = a[i]
					i++
				} else {
					ret[i+j] = b[j]
					j++
				}
			}
		}
	}
	return ret
}

func maxNumber(nums1 []int, nums2 []int, k int) []int {
	ret := make([]int, k)
	for i := 0; i <= k; i++ {
		if len(nums1) >= i && len(nums2) >= k-i {
			ret = maxArray(ret, mergeArray(getMaxArray(nums1, i), getMaxArray(nums2, k-i)))
		}
	}
	return ret
}
func main() {
	fmt.Println(maxNumber([]int{6, 3, 9, 0, 5, 6}, []int{2, 2, 5, 2, 1, 4, 4, 5, 7, 8, 9, 3, 1, 6, 9, 7, 0}, 23))
}
