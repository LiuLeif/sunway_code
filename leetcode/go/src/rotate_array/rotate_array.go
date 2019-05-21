// 2017-12-11 16:36
package main

import "fmt"

func rotate(nums []int, k int) {
	k %= len(nums)
	right := nums[:len(nums)-k]
	left := nums[len(nums)-k:]
	copy(nums, append(left, right...))
}

func reverse(nums []int) {
	n := len(nums)
	for i := 0; i < n/2; i++ {
		nums[i], nums[n-i-1] = nums[n-i-1], nums[i]
	}
}
func rotateByReverse(nums []int, k int) {
	k %= len(nums)
	reverse(nums)
	reverse(nums[:k])
	reverse(nums[k:])
}
func main() {
	// 1 2 3 4 5 6 7
	// 5 6 7 1 2 3 4
	// 7 6 5 4 3 2 1
	//
	x := []int{1, 2, 3, 4, 5, 6, 7}
	rotateByReverse(x, 3)
	fmt.Println(x)
}
