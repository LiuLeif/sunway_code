// 2018-10-29 22:59
package main

import "fmt"

func find132pattern(nums []int) bool {
	firstSmaller := make([]int, len(nums))
	for i := 0; i < len(firstSmaller); i++ {
		firstSmaller[i] = -1
	}
	for i := 0; i < len(nums); i++ {
		for j := i + 1; j < len(nums); j++ {
			if nums[i] < nums[j] && firstSmaller[j] == -1 {
				firstSmaller[j] = i
			}
			if nums[i] > nums[j] {
				if i > firstSmaller[j] && firstSmaller[j] != -1 {
					return true
				}
			}
		}
	}
	return false
}

func main() {
	fmt.Println(find132pattern([]int{3, 1, 4, 3}))
}
