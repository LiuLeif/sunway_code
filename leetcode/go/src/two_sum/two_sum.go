// 2017-11-14 18:54
package main

import "fmt"

func twoSum(nums []int, target int) []int {
	m := make(map[int]int)
	for i, v := range nums {
		m[v] = i
	}
	for i := 0; i < len(nums); i++ {
		j, ok := m[target-nums[i]]
		if ok && i != j {
			return []int{i, j}
		}
	}
	return []int{0, 0}
}

func main() {
	x := []int{3, 2, 4}
	target := 6
	result := twoSum(x, target)
	fmt.Println(result)
}
