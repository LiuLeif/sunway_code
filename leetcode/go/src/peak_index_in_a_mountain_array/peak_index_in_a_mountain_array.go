// 2018-10-13 23:07
package main

import "fmt"

func peakIndexInMountainArray(A []int) int {
	left, right := 0, len(A)-1
	for left < right {
		mid := (left + right) / 2
		if A[mid] > A[mid+1] {
			right = mid
		} else {
			left = mid + 1
		}
	}
	return left
}
func main() {
	fmt.Println(peakIndexInMountainArray([]int{0, 1, 2, 0}))
}
