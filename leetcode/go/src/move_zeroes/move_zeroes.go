// 2017-12-29 14:47
package main

import "fmt"

func moveZeroes(nums []int) {
	lastPos := 0
	for curr := 0; curr < len(nums); curr++ {
		if nums[curr] != 0 {
			nums[curr], nums[lastPos] = nums[lastPos], nums[curr]
			lastPos++
		}
	}
}
func main() {
	nums := []int{2, 1, 0, 3, 12}
	moveZeroes(nums)
	fmt.Println(nums)
}
