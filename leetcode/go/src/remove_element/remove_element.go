// 2017-12-29 14:47
package main

import "fmt"

func removeElement(nums []int, val int) int {
	lastPos := 0
	for i := 0; i < len(nums); i++ {
		if nums[i] != val {
			nums[lastPos], nums[i] = nums[i], nums[lastPos]
			lastPos++
		}
	}
	return lastPos
}
func main() {
	fmt.Println(removeElement([]int{3, 2, 2, 3}, 3))
}
