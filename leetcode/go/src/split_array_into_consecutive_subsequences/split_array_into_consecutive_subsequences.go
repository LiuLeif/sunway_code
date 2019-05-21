// 2017-11-17 11:20
package main

import "fmt"

func isPossible(nums []int) bool {
	m := map[int]int{}
	group := map[int]int{}

	for _, v := range nums {
		m[v]++
	}

	for _, v := range nums {
		if m[v] == 0 {
			continue
		}
		m[v]--
		if group[v-1] != 0 {
			group[v-1]--
			group[v]++
		} else {
			if m[v+1] != 0 && m[v+2] != 0 {
				m[v+1]--
				m[v+2]--
				group[v+2]++
			} else {
				return false
			}
		}
	}

	return true
}

func main() {
	fmt.Println(isPossible([]int{}))
}
