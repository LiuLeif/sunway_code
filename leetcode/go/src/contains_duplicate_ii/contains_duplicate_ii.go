// 2017-11-14 18:54
package main

import "fmt"

func containsNearbyDuplicate(nums []int, k int) bool {
	m := make(map[int]int)
	for i := 0; i < len(nums); i++ {
		p, ok := m[nums[i]]
		if ok && i-p <= k {
			return true
		}
		m[nums[i]] = i
	}

	return false
}
func main() {
	fmt.Println(containsNearbyDuplicate([]int{1, 2, 1}, 2))
}
