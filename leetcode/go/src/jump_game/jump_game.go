// 2017-11-20 16:14
package main

import "fmt"

// LTE
func canJumpRecursive(nums []int) bool {
	if len(nums) == 1 {
		return true
	}
	for i := 1; i <= nums[0]; i++ {
		if canJumpRecursive(nums[i:]) {
			return true
		}
	}

	return false
}

// LTE
func canJumpHash(nums []int) bool {
	m := make(map[int]bool)
	m[len(nums)-1] = true
	for i := len(nums) - 2; i >= 0; i-- {
		for j := 1; j <= nums[i]; j++ {
			if m[i+j] {
				m[i] = true
				break
			}
		}
	}

	return m[0]
}

func canJump(nums []int) bool {
	lastPos := len(nums) - 1
	for i := lastPos - 1; i >= 0; i-- {
		if i+nums[i] >= lastPos {
			lastPos = i
		}
	}

	return lastPos == 0
}

func main() {
	fmt.Println(canJump([]int{2, 2, 2, 0, 4}))
	fmt.Println(canJump([]int{3, 2, 1, 1, 4}))
}
