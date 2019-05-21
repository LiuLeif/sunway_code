// 2017-12-26 16:09
package main

import "fmt"

func numberOfArithmeticSlices(nums []int) int {
	ret := 0
	lo := 0
	for lo < len(nums)-2 {
		delta := nums[lo+1] - nums[lo]
		hi := lo + 2
		for hi < len(nums) {
			if nums[hi]-nums[hi-1] != delta {
				break
			}
			hi++
		}
		k := hi - lo
		for i := 1; i < k-1; i++ {
			ret += i
		}
		lo = hi - 1
	}
	return ret
}
func main() {
	fmt.Println(numberOfArithmeticSlices([]int{1, 2, 2, 2, 3, 4}))
}
