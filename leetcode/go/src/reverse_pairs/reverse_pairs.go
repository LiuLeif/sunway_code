// 2018-02-01 17:08
package main

import (
	"fmt"
)

func merge(nums []int, mid int) {
	left, right := append([]int{}, nums[:mid]...), append([]int{}, nums[mid:]...)
	i := 0

	for len(left) != 0 || len(right) != 0 {
		if len(left) == 0 || (len(right) != 0 && left[0] >= right[0]) {
			nums[i] = right[0]
			right = right[1:]
		} else {
			nums[i] = left[0]
			left = left[1:]
		}
		i++
	}
}

func mergeSort(nums []int) int {
	if len(nums) <= 1 {
		return 0
	}
	mid := len(nums) / 2
	ret := mergeSort(nums[:mid]) + mergeSort(nums[mid:])
	f := 0
	// fmt.Println("merge", nums[:mid], nums[mid:])
	// 156 12
	// i   j
	for i, j := f, mid; i < mid; i++ {
		for j < len(nums) && float64(float64(nums[i])/2.0) > float64(nums[j]) {
			j++
		}
		ret += j - mid
	}

	merge(nums, mid)
	return ret
}

func reversePairs(nums []int) int {
	// 123
	// 13
	ret := mergeSort(nums)
	return ret
}
func main() {
	fmt.Println(reversePairs([]int{1, 3, 2, 3, 1}))
	fmt.Println(reversePairs([]int{2, 4, 3, 5, 1}))
	fmt.Println(reversePairs([]int{-5, -5}))
	fmt.Println(reversePairs([]int{4, 5, 6, 1, 2, 3}))
	fmt.Println(reversePairs([]int{233, 2000000001, 234, 2000000006, 235, 2000000003, 236, 2000000007, 237, 2000000002, 2000000005, 233, 233, 233, 233, 233, 2000000004}))
}
