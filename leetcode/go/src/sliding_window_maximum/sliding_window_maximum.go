// 2017-11-14 18:54
package main

import "fmt"

func maxSlidingWindow(nums []int, k int) []int {
	if k == 0 {
		return []int{}
	}
	queue := make([]int, k)
	ret := []int{}

	for i := 0; i < k; i++ {
		for len(queue) != 0 && nums[i] >= nums[queue[len(queue)-1]] {
			queue = queue[:len(queue)-1]
		}
		queue = append(queue, i)
	}

	for i := k; i < len(nums); i++ {
		ret = append(ret, nums[queue[0]])
		for len(queue) != 0 && queue[0] <= i-k {
			queue = queue[1:]
		}
		for len(queue) != 0 && nums[i] >= nums[queue[len(queue)-1]] {
			queue = queue[:len(queue)-1]
		}
		queue = append(queue, i)
	}
	ret = append(ret, nums[queue[0]])
	return ret
}
func main() {
	fmt.Println(maxSlidingWindow([]int{}, 0))
}
