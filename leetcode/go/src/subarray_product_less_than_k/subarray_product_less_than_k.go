// 2017-11-14 18:54
package main

import "fmt"

func numSubarrayProductLessThanK(nums []int, k int) int {
	if k <= 1 {
		return 0
	}
	prod := 1
	left := 0
	ret := 0
	for right := 0; right < len(nums); right++ {
		prod *= nums[right]
		for prod >= k {
			prod /= nums[left]
			left += 1
		}
		ret += right - left + 1
		fmt.Println(right - left + 1)
	}

	return ret
}
func main() {
	fmt.Println(numSubarrayProductLessThanK([]int{10, 5, 2, 6}, 100))
}
