// 2017-12-04 17:39
package main

import "fmt"

func nextIndex(i int, nums []int) int {
	ret := (i + nums[i]) % (len(nums))
	if ret < 0 {
		ret += len(nums)
	}
	return ret
}

func circularArrayLoop(nums []int) bool {
	for i := 0; i < len(nums); i++ {
		if nums[i] == 0 {
			continue
		}
		slow, fast := i, nextIndex(i, nums)
		startValue := nums[i]
		for slow != fast {
			slow = nextIndex(slow, nums)
			fast = nextIndex(fast, nums)
			if startValue*nums[slow] < 0 || startValue*nums[fast] < 0 {
				break
			}
			fast = nextIndex(fast, nums)
			if startValue*nums[fast] < 0 {
				break
			}
		}
		if slow == fast {
			if nextIndex(slow, nums) != slow {
				return true
			}
		}
		// mark all from i as 0
		start := i
		for startValue*nums[start] > 0 {
			nums[start] = 0
			start = nextIndex(start, nums)
		}
	}

	return false
}

func main() {
	fmt.Println(circularArrayLoop([]int{-2, 1, -1, -2, -2}))
	fmt.Println(circularArrayLoop([]int{2, -1, 1, 2, 2}))
	fmt.Println(circularArrayLoop([]int{3, 1, 2}))
	fmt.Println(circularArrayLoop([]int{1, 2, 3}))
}
