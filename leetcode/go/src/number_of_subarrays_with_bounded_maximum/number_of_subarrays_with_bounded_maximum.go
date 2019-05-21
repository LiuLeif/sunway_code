// 2018-03-22 14:49
package main

import "fmt"

func numSubarrayBoundedMax(A []int, L int, R int) int {
	ret := 0
	left, right := -1, -1
	for i := 0; i < len(A); i++ {
		if A[i] > R {
			left = i
		}
		if A[i] >= L {
			right = i
		}
		ret += right - left
	}
	return ret
}
func main() {
	fmt.Println(numSubarrayBoundedMax([]int{2, 1}, 2, 4))
}
