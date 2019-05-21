// 2017-12-05 10:50
package main

import "fmt"

func majorityElement(nums []int) int {
	candidate := 0
	count := 0
	for _, v := range nums {
		if count == 0 {
			candidate = v
			count = 1
		} else {
			if v == candidate {
				count++
			} else {
				count--
			}
		}
	}

	return candidate
}
func main() {
	fmt.Println(majorityElement([]int{1, 1, 2, 2, 2}))
}
