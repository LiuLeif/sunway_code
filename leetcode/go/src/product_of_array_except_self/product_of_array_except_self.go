// 2017-12-06 16:13
package main

import "fmt"

func productExceptSelf(nums []int) []int {
	ret := make([]int, len(nums))
	n := len(nums)
	ret[0] = 1

	for i := 1; i < n; i++ {
		ret[i] = ret[i-1] * nums[i-1]
	}

	right := 1
	for i := n - 1; i >= 0; i-- {
		ret[i] = ret[i] * right
		right *= nums[i]
	}

	return ret
}
func main() {
	fmt.Println(productExceptSelf([]int{1, 2, 3, 4}))
}
